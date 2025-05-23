// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.7
// source: google/shopping/merchant/reviews/v1beta/merchantreviews.proto

package reviewspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MerchantReviewsService_GetMerchantReview_FullMethodName    = "/google.shopping.merchant.reviews.v1beta.MerchantReviewsService/GetMerchantReview"
	MerchantReviewsService_ListMerchantReviews_FullMethodName  = "/google.shopping.merchant.reviews.v1beta.MerchantReviewsService/ListMerchantReviews"
	MerchantReviewsService_InsertMerchantReview_FullMethodName = "/google.shopping.merchant.reviews.v1beta.MerchantReviewsService/InsertMerchantReview"
	MerchantReviewsService_DeleteMerchantReview_FullMethodName = "/google.shopping.merchant.reviews.v1beta.MerchantReviewsService/DeleteMerchantReview"
)

// MerchantReviewsServiceClient is the client API for MerchantReviewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantReviewsServiceClient interface {
	// Gets a merchant review.
	GetMerchantReview(ctx context.Context, in *GetMerchantReviewRequest, opts ...grpc.CallOption) (*MerchantReview, error)
	// Lists merchant reviews.
	ListMerchantReviews(ctx context.Context, in *ListMerchantReviewsRequest, opts ...grpc.CallOption) (*ListMerchantReviewsResponse, error)
	// Inserts a review for your Merchant Center account. If the review
	// already exists, then the review is replaced with the new instance.
	InsertMerchantReview(ctx context.Context, in *InsertMerchantReviewRequest, opts ...grpc.CallOption) (*MerchantReview, error)
	// Deletes merchant review.
	DeleteMerchantReview(ctx context.Context, in *DeleteMerchantReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type merchantReviewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantReviewsServiceClient(cc grpc.ClientConnInterface) MerchantReviewsServiceClient {
	return &merchantReviewsServiceClient{cc}
}

func (c *merchantReviewsServiceClient) GetMerchantReview(ctx context.Context, in *GetMerchantReviewRequest, opts ...grpc.CallOption) (*MerchantReview, error) {
	out := new(MerchantReview)
	err := c.cc.Invoke(ctx, MerchantReviewsService_GetMerchantReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantReviewsServiceClient) ListMerchantReviews(ctx context.Context, in *ListMerchantReviewsRequest, opts ...grpc.CallOption) (*ListMerchantReviewsResponse, error) {
	out := new(ListMerchantReviewsResponse)
	err := c.cc.Invoke(ctx, MerchantReviewsService_ListMerchantReviews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantReviewsServiceClient) InsertMerchantReview(ctx context.Context, in *InsertMerchantReviewRequest, opts ...grpc.CallOption) (*MerchantReview, error) {
	out := new(MerchantReview)
	err := c.cc.Invoke(ctx, MerchantReviewsService_InsertMerchantReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantReviewsServiceClient) DeleteMerchantReview(ctx context.Context, in *DeleteMerchantReviewRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MerchantReviewsService_DeleteMerchantReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantReviewsServiceServer is the server API for MerchantReviewsService service.
// All implementations should embed UnimplementedMerchantReviewsServiceServer
// for forward compatibility
type MerchantReviewsServiceServer interface {
	// Gets a merchant review.
	GetMerchantReview(context.Context, *GetMerchantReviewRequest) (*MerchantReview, error)
	// Lists merchant reviews.
	ListMerchantReviews(context.Context, *ListMerchantReviewsRequest) (*ListMerchantReviewsResponse, error)
	// Inserts a review for your Merchant Center account. If the review
	// already exists, then the review is replaced with the new instance.
	InsertMerchantReview(context.Context, *InsertMerchantReviewRequest) (*MerchantReview, error)
	// Deletes merchant review.
	DeleteMerchantReview(context.Context, *DeleteMerchantReviewRequest) (*emptypb.Empty, error)
}

// UnimplementedMerchantReviewsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMerchantReviewsServiceServer struct {
}

func (UnimplementedMerchantReviewsServiceServer) GetMerchantReview(context.Context, *GetMerchantReviewRequest) (*MerchantReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMerchantReview not implemented")
}
func (UnimplementedMerchantReviewsServiceServer) ListMerchantReviews(context.Context, *ListMerchantReviewsRequest) (*ListMerchantReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMerchantReviews not implemented")
}
func (UnimplementedMerchantReviewsServiceServer) InsertMerchantReview(context.Context, *InsertMerchantReviewRequest) (*MerchantReview, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMerchantReview not implemented")
}
func (UnimplementedMerchantReviewsServiceServer) DeleteMerchantReview(context.Context, *DeleteMerchantReviewRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchantReview not implemented")
}

// UnsafeMerchantReviewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantReviewsServiceServer will
// result in compilation errors.
type UnsafeMerchantReviewsServiceServer interface {
	mustEmbedUnimplementedMerchantReviewsServiceServer()
}

func RegisterMerchantReviewsServiceServer(s grpc.ServiceRegistrar, srv MerchantReviewsServiceServer) {
	s.RegisterService(&MerchantReviewsService_ServiceDesc, srv)
}

func _MerchantReviewsService_GetMerchantReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMerchantReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantReviewsServiceServer).GetMerchantReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantReviewsService_GetMerchantReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantReviewsServiceServer).GetMerchantReview(ctx, req.(*GetMerchantReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantReviewsService_ListMerchantReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMerchantReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantReviewsServiceServer).ListMerchantReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantReviewsService_ListMerchantReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantReviewsServiceServer).ListMerchantReviews(ctx, req.(*ListMerchantReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantReviewsService_InsertMerchantReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertMerchantReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantReviewsServiceServer).InsertMerchantReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantReviewsService_InsertMerchantReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantReviewsServiceServer).InsertMerchantReview(ctx, req.(*InsertMerchantReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantReviewsService_DeleteMerchantReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMerchantReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantReviewsServiceServer).DeleteMerchantReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MerchantReviewsService_DeleteMerchantReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantReviewsServiceServer).DeleteMerchantReview(ctx, req.(*DeleteMerchantReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantReviewsService_ServiceDesc is the grpc.ServiceDesc for MerchantReviewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantReviewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.reviews.v1beta.MerchantReviewsService",
	HandlerType: (*MerchantReviewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMerchantReview",
			Handler:    _MerchantReviewsService_GetMerchantReview_Handler,
		},
		{
			MethodName: "ListMerchantReviews",
			Handler:    _MerchantReviewsService_ListMerchantReviews_Handler,
		},
		{
			MethodName: "InsertMerchantReview",
			Handler:    _MerchantReviewsService_InsertMerchantReview_Handler,
		},
		{
			MethodName: "DeleteMerchantReview",
			Handler:    _MerchantReviewsService_DeleteMerchantReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/reviews/v1beta/merchantreviews.proto",
}
