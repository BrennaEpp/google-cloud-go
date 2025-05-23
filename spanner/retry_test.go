/*
Copyright 2017 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spanner

import (
	"context"
	"testing"
	"time"

	"github.com/googleapis/gax-go/v2"
	edpb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

func TestRetryInfo(t *testing.T) {
	s := status.New(codes.Aborted, "")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(time.Second),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	gotDelay, ok := ExtractRetryDelay(toSpannerErrorWithCommitInfo(s.Err(), true))
	if !ok || !testEqual(time.Second, gotDelay) {
		t.Errorf("<ok, retryDelay> = <%t, %v>, want <true, %v>", ok, gotDelay, time.Second)
	}
}

func TestRetryInfoResourceExhausted(t *testing.T) {
	s := status.New(codes.ResourceExhausted, "")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(time.Second),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	gotDelay, ok := ExtractRetryDelay(toSpannerErrorWithCommitInfo(s.Err(), true))
	if !ok || !testEqual(time.Second, gotDelay) {
		t.Errorf("<ok, retryDelay> = <%t, %v>, want <true, %v>", ok, gotDelay, time.Second)
	}
}

func TestRetryInfoInWrappedError(t *testing.T) {
	s := status.New(codes.Aborted, "")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(time.Second),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	gotDelay, ok := ExtractRetryDelay(
		&wrappedTestError{wrapped: toSpannerErrorWithCommitInfo(s.Err(), true), msg: "Error that is wrapping a Spanner error"},
	)
	if !ok || !testEqual(time.Second, gotDelay) {
		t.Errorf("<ok, retryDelay> = <%t, %v>, want <true, %v>", ok, gotDelay, time.Second)
	}
}

func TestRetryInfoInWrappedErrorResourceExhausted(t *testing.T) {
	s := status.New(codes.ResourceExhausted, "")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(time.Second),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	gotDelay, ok := ExtractRetryDelay(
		&wrappedTestError{wrapped: toSpannerErrorWithCommitInfo(s.Err(), true), msg: "Error that is wrapping a Spanner error"},
	)
	if !ok || !testEqual(time.Second, gotDelay) {
		t.Errorf("<ok, retryDelay> = <%t, %v>, want <true, %v>", ok, gotDelay, time.Second)
	}
}

func TestRetryInfoTransactionOutcomeUnknownError(t *testing.T) {
	err := toSpannerErrorWithCommitInfo(context.DeadlineExceeded, true)
	if gotDelay, ok := ExtractRetryDelay(err); ok {
		t.Errorf("Got unexpected delay\nGot: %v\nWant: %v", gotDelay, 0)
	}
	want := &TransactionOutcomeUnknownError{status.FromContextError(context.DeadlineExceeded).Err()}
	if !testEqual(err.(*Error).err.Error(), want.Error()) {
		t.Errorf("Missing expected TransactionOutcomeUnknownError wrapped error")
	}
}

func TestRetryerRespectsServerDelay(t *testing.T) {
	t.Parallel()
	serverDelay := 50 * time.Millisecond
	s := status.New(codes.Aborted, "transaction was aborted")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(serverDelay),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	retryer := onCodes(gax.Backoff{}, codes.Aborted)
	err = toSpannerErrorWithCommitInfo(s.Err(), true)
	maxSeenDelay, shouldRetry := retryer.Retry(err)
	if !shouldRetry {
		t.Fatalf("expected shouldRetry to be true")
	}
	if maxSeenDelay != serverDelay {
		t.Fatalf("Retry delay mismatch:\ngot: %v\nwant: %v", maxSeenDelay, serverDelay)
	}
}

func TestRetryerRespectsServerDelayResourceExhausted(t *testing.T) {
	t.Parallel()
	serverDelay := 50 * time.Millisecond
	s := status.New(codes.ResourceExhausted, "transaction was aborted")
	s, err := s.WithDetails(&edpb.RetryInfo{
		RetryDelay: durationpb.New(serverDelay),
	})
	if err != nil {
		t.Fatalf("Error setting retry details: %v", err)
	}
	retryer := onCodes(gax.Backoff{}, codes.ResourceExhausted)
	err = toSpannerErrorWithCommitInfo(s.Err(), true)
	maxSeenDelay, shouldRetry := retryer.Retry(err)
	if !shouldRetry {
		t.Fatalf("expected shouldRetry to be true")
	}
	if maxSeenDelay != serverDelay {
		t.Fatalf("Retry delay mismatch:\ngot: %v\nwant: %v", maxSeenDelay, serverDelay)
	}
}

func TestRunWithRetryOnInternalAuthError(t *testing.T) {
	ctx := context.Background()
	targetErr := status.Error(codes.Internal, "Authentication backend internal server error. Please retry")

	callCount := 0
	// Mock function that fails with the target error first, then succeeds on the second attempt after retry.
	mockFunc := func(ctx context.Context) error {
		callCount++
		if callCount == 1 {
			return targetErr
		}
		return nil
	}

	// Use a very short backoff for testing purposes to speed it up.
	originalBackoff := DefaultRetryBackoff
	DefaultRetryBackoff = gax.Backoff{
		Initial:    1 * time.Nanosecond,
		Max:        10 * time.Nanosecond,
		Multiplier: 1.0,
	}
	defer func() { DefaultRetryBackoff = originalBackoff }() // Restore original backoff

	err := runWithRetryOnAbortedOrFailedInlineBeginOrSessionNotFound(ctx, mockFunc)

	if err != nil {
		t.Errorf("Expected runWithRetry to succeed after retry, but got error: %v", err)
	}

	if callCount != 2 {
		t.Errorf("Expected mockFunc to be called 2 times (1 initial + 1 retry), but got %d", callCount)
	}
}
