// Copyright 2025 Google LLC
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

package pubsub

import (
	"context"
	"testing"
	"time"

	vkit "cloud.google.com/go/pubsub/v2/apiv1"
	pb "cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"cloud.google.com/go/pubsub/v2/pstest"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestClient_ApplyClientConfig(t *testing.T) {
	ctx := context.Background()
	srv := pstest.NewServer()
	// Add a retry for an obscure error.
	tco := &vkit.TopicAdminCallOptions{
		Publish: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DataLoss,
				}, gax.Backoff{
					Initial:    200 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.25,
				})
			}),
		},
	}
	c, err := NewClientWithConfig(ctx, "P", &ClientConfig{
		TopicAdminCallOptions: tco,
	},
		option.WithEndpoint(srv.Addr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		t.Fatal(err)
	}

	srv.SetAutoPublishResponse(false)
	// Create a fake publish response with the obscure error we are retrying.
	srv.AddPublishResponse(&pb.PublishResponse{
		MessageIds: []string{},
	}, status.Errorf(codes.DataLoss, "obscure error"))

	srv.AddPublishResponse(&pb.PublishResponse{
		MessageIds: []string{"1"},
	}, nil)

	topic, err := c.TopicAdminClient.CreateTopic(ctx, &pb.Topic{Name: "projects/P/topics/t"})
	if err != nil {
		t.Fatal(err)
	}
	publisher := c.Publisher(topic.Name)
	res := publisher.Publish(ctx, &Message{
		Data: []byte("test"),
	})
	if id, err := res.Get(ctx); err != nil {
		t.Fatalf("got error from res.Get(): %v", err)
	} else {
		if id != "1" {
			t.Fatalf("got wrong message id from server, got %s, want 1", id)
		}
	}
}

func TestClient_EmptyProjectID(t *testing.T) {
	ctx := context.Background()
	_, err := NewClient(ctx, "")
	if err != ErrEmptyProjectID {
		t.Fatalf("passing empty project ID got %v, want%v", err, ErrEmptyProjectID)
	}
}
