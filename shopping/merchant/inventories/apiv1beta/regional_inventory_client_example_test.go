// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package inventories_test

import (
	"context"

	inventories "cloud.google.com/go/shopping/merchant/inventories/apiv1beta"
	inventoriespb "cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb"
	"google.golang.org/api/iterator"
)

func ExampleNewRegionalInventoryClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := inventories.NewRegionalInventoryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRegionalInventoryRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := inventories.NewRegionalInventoryRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRegionalInventoryClient_DeleteRegionalInventory() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := inventories.NewRegionalInventoryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &inventoriespb.DeleteRegionalInventoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb#DeleteRegionalInventoryRequest.
	}
	err = c.DeleteRegionalInventory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRegionalInventoryClient_InsertRegionalInventory() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := inventories.NewRegionalInventoryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &inventoriespb.InsertRegionalInventoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb#InsertRegionalInventoryRequest.
	}
	resp, err := c.InsertRegionalInventory(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRegionalInventoryClient_ListRegionalInventories() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := inventories.NewRegionalInventoryClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &inventoriespb.ListRegionalInventoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/shopping/merchant/inventories/apiv1beta/inventoriespb#ListRegionalInventoriesRequest.
	}
	it := c.ListRegionalInventories(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*inventoriespb.ListRegionalInventoriesResponse)
	}
}
