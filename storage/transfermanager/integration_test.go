// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transfermanager

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	"cloud.google.com/go/internal/uid"
	"cloud.google.com/go/storage"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	testPrefix      = "go-integration-test-tm"
	grpcTestPrefix  = "golang-grpc-test-tm"
	bucketExpiryAge = 24 * time.Hour
	minObjectSize   = 1024
	maxObjectSize   = 1024 * 1024
)

var (
	uidSpace = uid.NewSpace("", nil)
	//  These buckets are shared amongst download tests. They are created,
	// populated with objects and cleaned up in TestMain.
	httpTestBucket = downloadTestBucket{}
	grpcTestBucket = downloadTestBucket{}
)

func TestMain(m *testing.M) {
	flag.Parse()

	if err := httpTestBucket.Create(testPrefix); err != nil {
		log.Fatalf("test bucket creation failed: %v", err)
	}

	if err := grpcTestBucket.Create(grpcTestPrefix); err != nil {
		log.Fatalf("test bucket creation failed: %v", err)
	}

	m.Run()

	if err := httpTestBucket.Cleanup(); err != nil {
		log.Printf("test bucket cleanup failed: %v", err)
	}
	if err := grpcTestBucket.Cleanup(); err != nil {
		log.Printf("grpc test bucket cleanup failed: %v", err)
	}
	if err := deleteExpiredBuckets(testPrefix); err != nil {
		log.Printf("expired http bucket cleanup failed: %v", err)
	}
	if err := deleteExpiredBuckets(grpcTestPrefix); err != nil {
		log.Printf("expired grpc bucket cleanup failed: %v", err)
	}
}

// Downloads the entire contents of the bucket synchronously.
func TestIntegration_DownloadDirectory(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		localDir := t.TempDir()
		numCallbacks := 0
		callbackMu := sync.Mutex{}

		d, err := NewDownloader(c, WithWorkers(8), WithPartSize(maxObjectSize/2))
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		if err := d.DownloadDirectory(ctx, &DownloadDirectoryInput{
			Bucket:         tb.bucket,
			LocalDirectory: localDir,
			StripPrefix:    "dir/",
			OnObjectDownload: func(got *DownloadOutput) {
				callbackMu.Lock()
				numCallbacks++
				callbackMu.Unlock()

				if got.Err != nil {
					t.Errorf("result.Err: %v", got.Err)
				}

				if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
					t.Errorf("expected object size %d, got %d", want, got)
				}

				objectWithoutPrefix, _ := strings.CutPrefix(got.Object, "dir/")
				path := filepath.Join(localDir, objectWithoutPrefix)

				f, err := os.Open(path)
				if err != nil {
					t.Errorf("os.Open(%q): %v", path, err)
				}
				defer f.Close()

				b := bytes.NewBuffer(make([]byte, 0, got.Attrs.Size))
				if _, err := io.Copy(b, f); err != nil {
					t.Errorf("io.Copy: %v", err)
				}

				if wantCRC, gotCRC := tb.contentHashes[got.Object], crc32c(b.Bytes()); gotCRC != wantCRC {
					t.Errorf("object(%q) at filepath(%q): content crc32c does not match; got: %v, expected: %v", got.Object, path, gotCRC, wantCRC)
				}
				got.Object = "modifying this shouldn't be a problem"
			},
		}); err != nil {
			t.Errorf("d.DownloadDirectory: %v", err)
		}

		results, err := d.WaitAndClose()
		if err != nil {
			t.Fatalf("d.WaitAndClose: %v", err)
		}

		if len(results) != len(tb.objects) {
			t.Errorf("expected to receive %d results, got %d results", len(tb.objects), len(results))
		}
		if numCallbacks != len(tb.objects) {
			t.Errorf("expected to receive %d callbacks, got %d", len(tb.objects), numCallbacks)
		}

		for _, got := range results {
			if got.Err != nil {
				t.Errorf("result.Err: %v", got.Err)
				continue
			}

			if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
				t.Errorf("expected object size %d, got %d", want, got)
			}

			objectWithoutPrefix, _ := strings.CutPrefix(got.Object, "dir/")
			path := filepath.Join(localDir, objectWithoutPrefix)
			f, err := os.Open(path)
			if err != nil {
				t.Errorf("os.Open(%q): %v", path, err)
			}
			defer f.Close()

			b := bytes.NewBuffer(make([]byte, 0, got.Attrs.Size))
			if _, err := io.Copy(b, f); err != nil {
				t.Errorf("io.Copy: %v", err)
			}

			if wantCRC, gotCRC := tb.contentHashes[got.Object], crc32c(b.Bytes()); gotCRC != wantCRC {
				t.Errorf("object(%q) at filepath(%q): content crc32c does not match; got: %v, expected: %v", got.Object, path, gotCRC, wantCRC)
			}
		}
	})
}

// Downloads a subset of objects using callbacks.
func TestIntegration_DownloadDirectoryAsync(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		localDir := t.TempDir()

		d, err := NewDownloader(c, WithWorkers(2), WithPartSize(maxObjectSize/2), WithCallbacks())
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		// Test download options - only download directory "dir".
		// In lex order we have:
		// "dir/nested/again/obj1", -- excluded by StartOffset
		// "dir/nested/objA", -- included
		// "dir/file" -- excluded by MatchGlob
		// "dir/objA", -- included
		// "dir/objB", -- included
		// "dir/objC", -- excluded by EndOffset
		includedObjs := []string{"dir/nested/objA", "dir/objA", "dir/objB"}

		objectDownloaded := make(chan bool)
		done := make(chan bool)

		trackObjectsDownloaded := func(numObjects *int) {
			for {
				select {
				case <-done:
					return
				case <-objectDownloaded:
					*numObjects++
				}
			}
		}

		objectsDownloaded := 0
		go trackObjectsDownloaded(&objectsDownloaded)

		if err := d.DownloadDirectory(ctx, &DownloadDirectoryInput{
			Bucket:         tb.bucket,
			LocalDirectory: localDir,
			Prefix:         "dir",
			StartOffset:    "dir/nested/o",
			EndOffset:      "dir/objC",
			MatchGlob:      "**obj**",
			OnObjectDownload: func(got *DownloadOutput) {
				objectDownloaded <- true

				if got.Err != nil {
					t.Errorf("result.Err: %v", got.Err)
				}

				if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
					t.Errorf("expected object size %d, got %d", want, got)
				}

				path := filepath.Join(localDir, got.Object)
				f, err := os.Open(path)
				if err != nil {
					t.Errorf("os.Open(%q): %v", path, err)
				}
				defer f.Close()

				b := bytes.NewBuffer(make([]byte, 0, got.Attrs.Size))
				if _, err := io.Copy(b, f); err != nil {
					t.Errorf("io.Copy: %v", err)
				}

				if wantCRC, gotCRC := tb.contentHashes[got.Object], crc32c(b.Bytes()); gotCRC != wantCRC {
					t.Errorf("object(%q) at filepath(%q): content crc32c does not match; got: %v, expected: %v", got.Object, path, gotCRC, wantCRC)
				}
				got.Object = "modifying this shouldn't be a problem"
			},
			Callback: func(outs []DownloadOutput) {
				if len(outs) != len(includedObjs) {
					t.Errorf("expected to receive %d results, got %d results", len(includedObjs), len(outs))
				}

				for _, got := range outs {
					if got.Err != nil {
						t.Errorf("result.Err: %v", got.Err)
						continue
					}

					if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
						t.Errorf("expected object size %d, got %d", want, got)
					}

					path := filepath.Join(localDir, got.Object)
					f, err := os.Open(path)
					if err != nil {
						t.Errorf("os.Open(%q): %v", path, err)
					}
					defer f.Close()

					b := bytes.NewBuffer(make([]byte, 0, got.Attrs.Size))
					if _, err := io.Copy(b, f); err != nil {
						t.Errorf("io.Copy: %v", err)
					}

					if wantCRC, gotCRC := tb.contentHashes[got.Object], crc32c(b.Bytes()); gotCRC != wantCRC {
						t.Errorf("object(%q) at filepath(%q): content crc32c does not match; got: %v, expected: %v", got.Object, path, gotCRC, wantCRC)
					}
				}
			},
		}); err != nil {
			t.Errorf("d.DownloadDirectory: %v", err)
		}

		if _, err := d.WaitAndClose(); err != nil {
			t.Errorf("d.WaitAndClose: %v", err)
		}
		done <- true

		if want, got := len(includedObjs), objectsDownloaded; want != got {
			t.Errorf("expected to receive %d callbacks, got %d", want, got)
		}

		entries, err := os.ReadDir(filepath.Join(localDir, "dir"))
		if err != nil {
			t.Fatalf("os.ReadDir: %v", err)
		}

		if len(entries) != len(includedObjs) {
			t.Errorf("expected %d entries in dir, got %d", len(includedObjs), len(entries))
		}

		for _, entry := range entries {
			if entry.IsDir() && entry.Name() != "nested" {
				t.Errorf("unexpected subdirectory %q in dir", entry.Name())
			}
			if !entry.IsDir() && entry.Name() != "objA" && entry.Name() != "objB" {
				t.Errorf("unexpected file %q in dir", entry.Name())
			}
		}

		// Test SkipIfExists() by attempting to download the entire directory.
		// The files previously downloaded (ie. includedObjects) should be skipped.
		d, err = NewDownloader(c, WithWorkers(2), SkipIfExists())
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		objectsDownloaded = 0
		go trackObjectsDownloaded(&objectsDownloaded)

		if err := d.DownloadDirectory(ctx, &DownloadDirectoryInput{
			Bucket:         tb.bucket,
			LocalDirectory: localDir,
			OnObjectDownload: func(got *DownloadOutput) {
				objectDownloaded <- true

				if got.Err != nil {
					t.Errorf("result.Err: %v", got.Err)
				}

				for _, obj := range includedObjs {
					if strings.EqualFold(got.Object, obj) {
						t.Errorf("should have skipped download of object %s", got.Object)
					}
				}

				if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
					t.Errorf("expected object size %d, got %d", want, got)
				}

				path := filepath.Join(localDir, got.Object)
				f, err := os.Open(path)
				if err != nil {
					t.Errorf("os.Open(%q): %v", path, err)
				}
				defer f.Close()

				b := bytes.NewBuffer(make([]byte, 0, got.Attrs.Size))
				if _, err := io.Copy(b, f); err != nil {
					t.Errorf("io.Copy: %v", err)
				}

				if wantCRC, gotCRC := tb.contentHashes[got.Object], crc32c(b.Bytes()); gotCRC != wantCRC {
					t.Errorf("object(%q) at filepath(%q): content crc32c does not match; got: %v, expected: %v", got.Object, path, gotCRC, wantCRC)
				}
			},
		}); err != nil {
			t.Errorf("d.DownloadDirectory: %v", err)
		}

		if _, err := d.WaitAndClose(); err != nil {
			t.Errorf("d.WaitAndClose: %v", err)
		}
		done <- true

		if want, got := len(tb.objects)-len(includedObjs), objectsDownloaded; want != got {
			t.Errorf("expected to receive %d callbacks, got %d", want, got)
		}
	})
}

// TestIntegration_DownloadDirectoryError tests that an error is returned if the
// file already exists and that any created file structure is deleted as well.
func TestIntegration_DownloadDirectoryError(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		localDir := t.TempDir()
		callbacks := make(chan bool)

		d, err := NewDownloader(c, WithWorkers(2))
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}
		defer d.WaitAndClose()

		// First download a nested file.
		obj := "dir/objC"
		if err := d.DownloadDirectory(ctx, &DownloadDirectoryInput{
			Bucket:         tb.bucket,
			LocalDirectory: localDir,
			Prefix:         obj,
			OnObjectDownload: func(got *DownloadOutput) {
				callbacks <- true

				if got.Err != nil {
					t.Errorf("result.Err: %v", got.Err)
				}

				if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
					t.Errorf("expected object size %d, got %d", want, got)
				}
			},
		}); err != nil {
			t.Fatalf("d.DownloadDirectory: %v", err)
		}

		// Then add another file and another directory to the temp dir.
		localNestedDir := "localonly"
		localFile := "localonly-file"

		f, err := os.Create(filepath.Join(localDir, localFile))
		if err != nil {
			t.Errorf("os.Create: %v", err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("f.Close: %v", err)
		}

		if err := os.Mkdir(filepath.Join(localDir, localNestedDir), fs.ModeDir|fs.ModePerm); err != nil {
			t.Errorf("os.Mkdir: %v", err)
		}

		// Now attempt to download the directory; it should fail.
		<-callbacks
		err = d.DownloadDirectory(ctx, &DownloadDirectoryInput{
			Bucket:         tb.bucket,
			LocalDirectory: localDir,
		})
		if !errors.Is(err, os.ErrExist) {
			t.Errorf("d.DownloadDirectory should have failed with error %q; got %v", os.ErrExist, err)
		}

		// Check the local directory, it should have the first file, the second and another directory only.
		expected := []string{obj, localNestedDir, localFile}

		entries, err := os.ReadDir(localDir)
		if err != nil {
			t.Errorf("os.ReadDir: %v", err)
		}
		if got, want := len(entries), len(expected); got != want {
			t.Errorf("localDir does not have the expected amount of entries, got %d, want %d", got, want)
		}

		for _, exp := range expected {
			_, err := os.Stat(filepath.Join(localDir, exp))
			if err != nil {
				t.Errorf("os.Stat: %v", err)
			}
		}
	})
}

func TestIntegration_DownloaderSynchronous(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		objects := tb.objects

		// Start a downloader. Give it a smaller amount of workers than objects,
		// to make sure we aren't blocking anywhere.
		d, err := NewDownloader(c, WithWorkers(2), WithPartSize(maxObjectSize/2))
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}
		// Download several objects.
		writers := make([]*DownloadBuffer, len(objects))
		objToWriter := make(map[string]int) // so we can map the resulting content back to the correct object
		for i, obj := range objects {
			writers[i] = NewDownloadBuffer(make([]byte, tb.objectSizes[obj]))
			objToWriter[obj] = i
			if err := d.DownloadObject(ctx, &DownloadObjectInput{
				Bucket:      tb.bucket,
				Object:      obj,
				Destination: writers[i],
			}); err != nil {
				t.Errorf("d.DownloadObject: %v", err)
			}
		}

		results, err := d.WaitAndClose()
		if err != nil {
			t.Fatalf("d.WaitAndClose: %v", err)
		}

		// Check the results.
		for _, got := range results {
			writerIdx := objToWriter[got.Object]

			if got.Err != nil {
				t.Errorf("result.Err: %v", got.Err)
				continue
			}

			if want, got := tb.contentHashes[got.Object], crc32c(writers[writerIdx].Bytes()); got != want {
				t.Fatalf("content crc32c does not match; got: %v, expected: %v", got, want)
			}

			if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
				t.Errorf("expected object size %d, got %d", want, got)
			}

			if got, want := got.Attrs.StartOffset, int64(0); want != got {
				t.Errorf("expected start offset %d, got %d", want, got)
			}
		}

		if len(results) != len(objects) {
			t.Errorf("expected to receive %d results, got %d results", len(objects), len(results))
		}
	})
}

// Tests that a single error does not affect the rest of the downloads.
func TestIntegration_DownloaderErrorSync(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		// Make a copy of the objects slice.
		objects := make([]string, len(tb.objects))
		copy(objects, tb.objects)

		// Add another object to attempt to download; since it hasn't been written,
		// this one will fail. Append to the start so that it will (likely) be
		// attempted in the first 2 downloads.
		nonexistentObject := "not-written"
		objects = append([]string{nonexistentObject}, objects...)

		// Start a downloader.
		d, err := NewDownloader(c, WithWorkers(2), WithPartSize(maxObjectSize/2))
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		// Download objects.
		writers := make([]*DownloadBuffer, len(objects))
		objToWriter := make(map[string]int) // so we can map the resulting content back to the correct object
		for i, obj := range objects {
			writers[i] = NewDownloadBuffer(make([]byte, tb.objectSizes[obj]))
			objToWriter[obj] = i
			if err := d.DownloadObject(ctx, &DownloadObjectInput{
				Bucket:      tb.bucket,
				Object:      obj,
				Destination: writers[i],
			}); err != nil {
				t.Errorf("d.DownloadObject: %v", err)
			}
		}

		// WaitAndClose should return an error since one of our downloads should have failed.
		results, err := d.WaitAndClose()
		if err == nil {
			t.Error("d.WaitAndClose should return an error, instead got nil")
		}

		// Check the results.
		for _, got := range results {
			writerIdx := objToWriter[got.Object]

			// Check that the nonexistent object returned an error.
			if got.Object == nonexistentObject {
				if !errors.Is(got.Err, storage.ErrObjectNotExist) {
					t.Errorf("Object(%q) should not exist, err found to be %v", got.Object, got.Err)
				}
				continue
			}

			// All other objects should complete correctly.
			if got.Err != nil {
				t.Errorf("result.Err: %v", got.Err)
				continue
			}

			if want, got := tb.contentHashes[got.Object], crc32c(writers[writerIdx].Bytes()); got != want {
				t.Fatalf("content crc32c does not match; got: %v, expected: %v", got, want)
			}

			if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
				t.Errorf("expected object size %d, got %d", want, got)
			}
		}

		if len(results) != len(objects) {
			t.Errorf("expected to receive %d results, got %d results", len(objects), len(results))
		}
	})
}

func TestIntegration_DownloaderAsynchronous(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		objects := tb.objects

		d, err := NewDownloader(c, WithWorkers(2), WithCallbacks(), WithPartSize(maxObjectSize/2))
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		numCallbacks := 0
		callbackMu := sync.Mutex{}

		// Download objects.
		writers := make([]*DownloadBuffer, len(objects))
		for i, obj := range objects {
			i := i
			writers[i] = NewDownloadBuffer(make([]byte, tb.objectSizes[obj]))
			if err := d.DownloadObject(ctx, &DownloadObjectInput{
				Bucket:      tb.bucket,
				Object:      obj,
				Destination: writers[i],
				Callback: func(got *DownloadOutput) {
					callbackMu.Lock()
					numCallbacks++
					callbackMu.Unlock()

					if got.Err != nil {
						t.Errorf("result.Err: %v", got.Err)
					}

					if want, got := tb.contentHashes[got.Object], crc32c(writers[i].Bytes()); got != want {
						t.Errorf("content crc32c does not match; got: %v, expected: %v", got, want)
					}

					if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
						t.Errorf("expected object size %d, got %d", want, got)
					}
				},
			}); err != nil {
				t.Errorf("d.DownloadObject: %v", err)
			}
		}

		if _, err := d.WaitAndClose(); err != nil {
			t.Fatalf("d.WaitAndClose: %v", err)
		}

		if numCallbacks != len(objects) {
			t.Errorf("expected to receive %d results, got %d callbacks", len(objects), numCallbacks)
		}
	})
}

func TestIntegration_DownloaderErrorAsync(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		d, err := NewDownloader(c, WithWorkers(2), WithCallbacks())
		if err != nil {
			t.Fatalf("NewDownloader: %v", err)
		}

		// Keep track of the number of callbacks. Since the callbacks may happen
		// in parallel, we sync access to this variable.
		numCallbacks := 0
		callbackMu := sync.Mutex{}

		// Download an object with incorrect generation.
		if err := d.DownloadObject(ctx, &DownloadObjectInput{
			Bucket:      tb.bucket,
			Object:      tb.objects[0],
			Destination: &DownloadBuffer{},
			Conditions: &storage.Conditions{
				GenerationMatch: -10,
			},
			Callback: func(got *DownloadOutput) {
				callbackMu.Lock()
				numCallbacks++
				callbackMu.Unlock()

				// This will match both the expected http and grpc errors.
				wantErr := errors.Join(&googleapi.Error{Code: 412}, status.Error(codes.FailedPrecondition, ""))

				if !errorIs(got.Err, wantErr) {
					t.Errorf("mismatching errors: got %v, want %v", got.Err, wantErr)
				}
			},
		}); err != nil {
			t.Errorf("d.DownloadObject: %v", err)
		}

		// Download existing objects.
		writers := make([]*DownloadBuffer, len(tb.objects))
		for i, obj := range tb.objects {
			i := i
			writers[i] = NewDownloadBuffer(make([]byte, tb.objectSizes[obj]))
			if err := d.DownloadObject(ctx, &DownloadObjectInput{
				Bucket:      tb.bucket,
				Object:      obj,
				Destination: writers[i],
				Callback: func(got *DownloadOutput) {
					callbackMu.Lock()
					numCallbacks++
					callbackMu.Unlock()

					if got.Err != nil {
						t.Errorf("result.Err: %v", got.Err)
					}

					if want, got := tb.contentHashes[got.Object], crc32c(writers[i].Bytes()); got != want {
						t.Errorf("content crc32c does not match; got: %v, expected: %v", got, want)
					}

					if got, want := got.Attrs.Size, tb.objectSizes[got.Object]; want != got {
						t.Errorf("expected object size %d, got %d", want, got)
					}
				},
			}); err != nil {
				t.Errorf("d.DownloadObject: %v", err)
			}
		}

		// Download a nonexistent object.
		if err := d.DownloadObject(ctx, &DownloadObjectInput{
			Bucket:      tb.bucket,
			Object:      "not-written",
			Destination: &DownloadBuffer{},
			Callback: func(got *DownloadOutput) {
				callbackMu.Lock()
				numCallbacks++
				callbackMu.Unlock()

				// Check that the nonexistent object returned an error.
				if !errors.Is(got.Err, storage.ErrObjectNotExist) {
					t.Errorf("Object(%q) should not exist, err found to be %v", got.Object, got.Err)
				}
			},
		}); err != nil {
			t.Errorf("d.DownloadObject: %v", err)
		}

		// WaitAndClose should return an error since 2 of our downloads should have failed.
		_, err = d.WaitAndClose()
		if err == nil {
			t.Error("d.WaitAndClose should return an error, instead got nil")
		}

		// Check that both errors were returned.
		wantErrs := []error{errors.Join(&googleapi.Error{Code: 412}, status.Error(codes.FailedPrecondition, "")),
			storage.ErrObjectNotExist}

		for _, want := range wantErrs {
			if !errorIs(err, want) {
				t.Errorf("got error does not wrap expected error %q, got:\n%v", want, err)
			}
		}

		// We expect num objects callbacks + 2 for the errored calls.
		if want, got := len(tb.objects)+2, numCallbacks; want != got {
			t.Errorf("expected to receive %d results, got %d callbacks", want, got)
		}
	})
}

func TestIntegration_DownloaderTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests skipped in short mode")
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	t.Cleanup(func() { client.Close() })

	// Start a downloader.
	d, err := NewDownloader(client, WithPerOpTimeout(time.Nanosecond))
	if err != nil {
		t.Fatalf("NewDownloader: %v", err)
	}

	// Download an object.
	if err := d.DownloadObject(ctx, &DownloadObjectInput{
		Bucket:      httpTestBucket.bucket,
		Object:      httpTestBucket.objects[0],
		Destination: &DownloadBuffer{},
	}); err != nil {
		t.Errorf("d.DownloadObject: %v", err)
	}

	// WaitAndClose should return an error since the timeout is too short.
	results, err := d.WaitAndClose()
	if err == nil {
		t.Error("d.WaitAndClose should return an error, instead got nil")
	}

	// Check the result.
	got := results[0]

	// Check that the nonexistent object returned an error.
	if got.Err != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded error, got: %v", got.Err)
	}
}

func TestIntegration_DownloadShard(t *testing.T) {
	multiTransportTest(context.Background(), t, func(t *testing.T, ctx context.Context, c *storage.Client, tb downloadTestBucket) {
		objectName := tb.objects[0]
		objectSize := tb.objectSizes[objectName]

		// Get expected Attrs.
		o := c.Bucket(tb.bucket).Object(objectName)
		r, err := o.NewReader(ctx)
		if err != nil {
			t.Fatalf("o.NewReader: %v", err)
		}

		incorrectGen := r.Attrs.Generation - 1

		cancelledCtx, cancel := context.WithCancel(ctx)
		cancel()

		for _, test := range []struct {
			desc    string
			timeout time.Duration
			in      *DownloadObjectInput
			want    *DownloadOutput
		}{
			{
				desc: "basic input",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
				},
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Attrs:  &r.Attrs,
				},
			},
			{
				desc: "range",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
					Range: &DownloadRange{
						Offset: objectSize - 5,
						Length: -1,
					},
				},
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Range: &DownloadRange{
						Offset: objectSize - 5,
						Length: -1,
					},
					Attrs: &storage.ReaderObjectAttrs{
						Size:            objectSize,
						StartOffset:     objectSize - 5,
						ContentType:     r.Attrs.ContentType,
						ContentEncoding: r.Attrs.ContentEncoding,
						CacheControl:    r.Attrs.CacheControl,
						LastModified:    r.Attrs.LastModified,
						Generation:      r.Attrs.Generation,
						Metageneration:  r.Attrs.Metageneration,
						CRC32C:          r.Attrs.CRC32C,
					},
				},
			},
			{
				desc: "incorrect generation",
				in: &DownloadObjectInput{
					Bucket:     tb.bucket,
					Object:     objectName,
					Generation: &incorrectGen,
				},
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Err:    storage.ErrObjectNotExist,
				},
			},
			{
				desc: "conditions: generationmatch",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
					Conditions: &storage.Conditions{
						GenerationMatch: r.Attrs.Generation,
					},
				},
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Attrs:  &r.Attrs,
				},
			},
			{
				desc: "conditions do not hold",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
					Conditions: &storage.Conditions{
						GenerationMatch: incorrectGen,
					},
				},
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Err:    errors.Join(&googleapi.Error{Code: 412}, status.Error(codes.FailedPrecondition, "")),
				},
			},
			{
				desc: "timeout",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
				},
				timeout: time.Nanosecond,
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Err:    context.DeadlineExceeded,
				},
			},
			{
				desc: "cancelled ctx",
				in: &DownloadObjectInput{
					Bucket: tb.bucket,
					Object: objectName,
					ctx:    cancelledCtx,
				},
				timeout: time.Nanosecond,
				want: &DownloadOutput{
					Bucket: tb.bucket,
					Object: objectName,
					Err:    context.Canceled,
				},
			},
		} {
			t.Run(test.desc, func(t *testing.T) {
				w := &DownloadBuffer{}
				test.in.Destination = w

				if test.in.ctx == nil {
					test.in.ctx = ctx
				}

				got := test.in.downloadShard(c, test.timeout, 1024)

				if got.Bucket != test.want.Bucket || got.Object != test.want.Object {
					t.Errorf("wanted bucket %q, object %q, got: %q, %q", test.want.Bucket, test.want.Object, got.Bucket, got.Object)
				}

				if diff := cmp.Diff(got.Range, test.want.Range); diff != "" {
					t.Errorf("DownloadOutput.Range: got(-) vs. want(+): %v", diff)
				}

				if diff := cmp.Diff(got.Attrs, test.want.Attrs); diff != "" {
					t.Errorf("DownloadOutput.Attrs: got(-) vs. want(+): %v", diff)
				}

				if !errorIs(got.Err, test.want.Err) {
					t.Errorf("mismatching errors: got %v, want %v", got.Err, test.want.Err)
				}
			})
		}
	})
}

// errorIs is equivalent to errors.Is, except that it additionally will return
// true if err and targetErr are googleapi.Errors with identical error codes,
// or if both errors have the same gRPC status code.
func errorIs(err error, targetErr error) bool {
	if errors.Is(err, targetErr) {
		return true
	}

	// Check http
	var e, targetE *googleapi.Error
	if errors.As(err, &e) && errors.As(targetErr, &targetE) {
		return e.Code == targetE.Code
	}

	// Check grpc errors
	if status.Code(err) != codes.Unknown {
		return status.Code(err) == status.Code(targetErr)
	}

	return false
}

// generateRandomFileInGCS uploads a file with random contents to GCS and returns
// the crc32c hash of the contents.
func generateFileInGCS(ctx context.Context, o *storage.ObjectHandle, size int64) (uint32, error) {
	w := o.Retryer(storage.WithPolicy(storage.RetryAlways)).NewWriter(ctx)

	crc32cHash := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	mw := io.MultiWriter(w, crc32cHash)

	if _, err := io.CopyN(mw, crand.Reader, size); err != nil {
		w.Close()
		return 0, err
	}
	return crc32cHash.Sum32(), w.Close()
}

// randomInt64 returns a value in the closed interval [min, max].
// That is, the endpoints are possible return values.
func randomInt64(min, max int64) int64 {
	if min > max {
		log.Fatalf("min cannot be larger than max; min: %d max: %d", min, max)
	}
	return rand.Int63n(max-min+1) + min
}

func deleteExpiredBuckets(prefix string) error {
	if testing.Short() {
		return nil
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("NewClient: %v", err)
	}

	projectID := testutil.ProjID()
	it := client.Buckets(ctx, projectID)
	it.Prefix = prefix
	for {
		bktAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if time.Since(bktAttrs.Created) > bucketExpiryAge {
			log.Printf("deleting bucket %q, which is more than %s old", bktAttrs.Name, bucketExpiryAge)
			if err := killBucket(ctx, client, bktAttrs.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

// killBucket deletes a bucket and all its objects.
func killBucket(ctx context.Context, client *storage.Client, bucketName string) error {
	bkt := client.Bucket(bucketName)
	// Bucket must be empty to delete.
	it := bkt.Objects(ctx, nil)
	for {
		objAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := bkt.Object(objAttrs.Name).Delete(ctx); err != nil {
			return fmt.Errorf("deleting %q: %v", bucketName+"/"+objAttrs.Name, err)
		}
	}
	// GCS is eventually consistent, so this delete may fail because the
	// replica still sees an object in the bucket. We log the error and expect
	// a later test run to delete the bucket.
	if err := bkt.Delete(ctx); err != nil {
		log.Printf("deleting %q: %v", bucketName, err)
	}
	return nil
}

// downloadTestBucket provides a bucket that can be reused for tests that only
// download from the bucket.
type downloadTestBucket struct {
	bucket        string
	objects       []string
	contentHashes map[string]uint32
	objectSizes   map[string]int64
}

// Create initializes the downloadTestBucket, creating a bucket and populating
// objects in it. All objects are of the same size but with different contents
// and can be mapped to their respective crc32c hash in contentHashes.
func (tb *downloadTestBucket) Create(prefix string) error {
	if testing.Short() {
		return nil
	}
	ctx := context.Background()

	tb.bucket = prefix + uidSpace.New()
	// Note: some tests depend on these object names.
	// Verify that tests pass when adding or removing objects.
	tb.objects = []string{
		"!#$&'()*+,:;=,?@,[] and spaces",
		"./obj",
		"obj1",
		"obj2",
		"dir/file",
		"dir/objA",
		"dir/objB",
		"dir/objC",
		"dir/nested/objA",
		"dir/nested/again/obj1",
		"anotherDir/objC",
	}
	tb.contentHashes = make(map[string]uint32)
	tb.objectSizes = make(map[string]int64)

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	b := client.Bucket(tb.bucket)
	if err := b.Create(ctx, testutil.ProjID(), nil); err != nil {
		return fmt.Errorf("bucket(%q).Create: %v", tb.bucket, err)
	}

	// Write objects.
	for _, obj := range tb.objects {
		size := randomInt64(minObjectSize, maxObjectSize)
		crc, err := generateFileInGCS(ctx, b.Object(obj), size)
		if err != nil {
			return fmt.Errorf("generateFileInGCS: %v", err)
		}
		tb.contentHashes[obj] = crc
		tb.objectSizes[obj] = size
	}
	return nil
}

// Cleanup deletes the objects and bucket created in Create.
func (tb *downloadTestBucket) Cleanup() error {
	if testing.Short() {
		return nil
	}
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	b := client.Bucket(tb.bucket)

	for _, obj := range tb.objects {
		if err := b.Object(obj).Delete(ctx); err != nil {
			return fmt.Errorf("object.Delete: %v", err)
		}
	}

	return b.Delete(ctx)
}

// multiTransportTest initializes fresh clients for each transport, then runs
// given testing function using each transport-specific client, supplying the
// test function with the sub-test instance, the context it was given, a test
// bucket and the client to use.
func multiTransportTest(ctx context.Context, t *testing.T, test func(*testing.T, context.Context, *storage.Client, downloadTestBucket)) {
	if testing.Short() {
		t.Skip("Integration tests skipped in short mode")
	}

	clients, err := initTransportClients(ctx)
	if err != nil {
		t.Fatalf("init clients: %v", err)
	}

	for transport, client := range clients {
		t.Run(transport, func(t *testing.T) {
			t.Cleanup(func() {
				client.Close()
			})

			testBucket := httpTestBucket

			if transport == "grpc" {
				testBucket = grpcTestBucket
			}

			test(t, ctx, client, testBucket)
		})
	}
}

func initTransportClients(ctx context.Context) (map[string]*storage.Client, error) {
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	gc, err := storage.NewGRPCClient(ctx)
	if err != nil {
		return nil, err
	}

	return map[string]*storage.Client{
		"http": c,
		"grpc": gc,
	}, nil
}

func crc32c(b []byte) uint32 {
	return crc32.Checksum(b, crc32.MakeTable(crc32.Castagnoli))
}
