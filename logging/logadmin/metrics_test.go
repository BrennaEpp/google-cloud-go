// Copyright 2016 Google LLC
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

package logadmin

import (
	"context"
	"log"
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	"cloud.google.com/go/internal/uid"
	ltest "cloud.google.com/go/logging/internal/testing"
	"google.golang.org/api/iterator"
)

var metricIDs = uid.NewSpace("GO-CLIENT-TEST-METRIC", nil)

// Initializes the tests before they run.
func initMetrics(ctx context.Context) {
	// Clean up from aborted tests.
	it := client.Metrics(ctx)
loop:
	for {
		m, err := it.Next()
		switch err {
		case nil:
			if metricIDs.Older(m.ID, 24*time.Hour) {
				client.DeleteMetric(ctx, m.ID)
			}
		case iterator.Done:
			break loop
		default:
			log.Printf("cleanupMetrics: %v", err)
			return
		}
	}
}

func TestCreateDeleteMetric(t *testing.T) {
	ctx := context.Background()
	metric := &Metric{
		ID:          metricIDs.New(),
		Description: "DESC",
		Filter:      "FILTER",
	}

	if err := client.CreateMetric(ctx, metric); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteMetric(ctx, metric.ID)

	var got *Metric
	ltest.Retry(t, func(r *testutil.R) error {
		var err error
		got, err = client.Metric(ctx, metric.ID)
		return err
	})
	if want := metric; !testutil.Equal(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	ltest.Retry(t, func(r *testutil.R) error {
		return client.DeleteMetric(ctx, metric.ID)
	})

	// client.Metric should give an error. Test if this is the case, but retry on
	// retryable errors.
	ltest.RetryAndExpectError(t, func(r *testutil.R) error {
		_, err := client.Metric(ctx, metric.ID)
		return err
	})
}

func TestUpdateMetric(t *testing.T) {
	ctx := context.Background()
	metric := &Metric{
		ID:          metricIDs.New(),
		Description: "DESC",
		Filter:      "FILTER",
	}

	// Updating a non-existent metric creates a new one.
	ltest.Retry(t, func(r *testutil.R) error {
		return client.UpdateMetric(ctx, metric)
	})
	defer client.DeleteMetric(ctx, metric.ID)

	var got *Metric
	ltest.Retry(t, func(r *testutil.R) error {
		var err error
		got, err = client.Metric(ctx, metric.ID)
		return err
	})
	if want := metric; !testutil.Equal(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	// Updating an existing metric changes it.
	metric.Description = "CHANGED"
	ltest.Retry(t, func(r *testutil.R) error {
		return client.UpdateMetric(ctx, metric)
	})
	ltest.Retry(t, func(r *testutil.R) error {
		var err error
		got, err = client.Metric(ctx, metric.ID)
		return err
	})
	if want := metric; !testutil.Equal(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestListMetrics(t *testing.T) {
	ctx := context.Background()

	var metrics []*Metric
	want := map[string]*Metric{}
	for i := 0; i < 10; i++ {
		m := &Metric{
			ID:          metricIDs.New(),
			Description: "DESC",
			Filter:      "FILTER",
		}
		metrics = append(metrics, m)
		want[m.ID] = m
	}
	for _, m := range metrics {
		ltest.Retry(t, func(r *testutil.R) error {
			return client.CreateMetric(ctx, m)
		})
		defer client.DeleteMetric(ctx, m.ID)
	}

	got := map[string]*Metric{}
	it := client.Metrics(ctx)
	for m, done := ltest.RetryIteratorNext(t, it); !done; m, done = ltest.RetryIteratorNext(t, it) {
		// If tests run simultaneously, we may have more metrics than we
		// created. So only check for our own.
		if _, ok := want[m.ID]; ok {
			got[m.ID] = m
		}
	}
	if !testutil.Equal(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
