// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: comunitiy.proto

package community

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComunityServiceClient is the client API for ComunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComunityServiceClient interface {
	CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error)
	GetCommunity(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityResponse, error)
	UpdateCommunity(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityResponse, error)
	DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error)
	ListCommunities(ctx context.Context, in *ListCommunitiesRequest, opts ...grpc.CallOption) (*ListCommunitiesResponse, error)
	JoinCommunity(ctx context.Context, in *JoinCommunityRequest, opts ...grpc.CallOption) (*JoinCommunityResponse, error)
	LeaveCommunity(ctx context.Context, in *LeaveCommunityRequest, opts ...grpc.CallOption) (*LeaveCommunityResponse, error)
	CreateCommunityEvent(ctx context.Context, in *CreateCommunityEventRequest, opts ...grpc.CallOption) (*CreateCommunityEventResponse, error)
	ListCommunityEvents(ctx context.Context, in *ListCommunityEventsRequest, opts ...grpc.CallOption) (*ListCommunityEventsResponse, error)
	CreateCommunityForumPost(ctx context.Context, in *CreateCommunityForumPostRequest, opts ...grpc.CallOption) (*CreateCommunityForumPostRespnse, error)
	ListCommunityForumPosts(ctx context.Context, in *ListCommunityForumPostsRequest, opts ...grpc.CallOption) (*ListCommunityForumPostsResponse, error)
	AddForumPostComment(ctx context.Context, in *AddForumPostCommentRequest, opts ...grpc.CallOption) (*AddForumPostCommentResponse, error)
	ListForumPostComments(ctx context.Context, in *ListForumPostCommentsRequest, opts ...grpc.CallOption) (*ListForumPostCommentsResponse, error)
}

type comunityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComunityServiceClient(cc grpc.ClientConnInterface) ComunityServiceClient {
	return &comunityServiceClient{cc}
}

func (c *comunityServiceClient) CreateCommunity(ctx context.Context, in *CreateCommunityRequest, opts ...grpc.CallOption) (*CreateCommunityResponse, error) {
	out := new(CreateCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/CreateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) GetCommunity(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityResponse, error) {
	out := new(GetCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/GetCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) UpdateCommunity(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityResponse, error) {
	out := new(UpdateCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/UpdateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) DeleteCommunity(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error) {
	out := new(DeleteCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/DeleteCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) ListCommunities(ctx context.Context, in *ListCommunitiesRequest, opts ...grpc.CallOption) (*ListCommunitiesResponse, error) {
	out := new(ListCommunitiesResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/ListCommunities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) JoinCommunity(ctx context.Context, in *JoinCommunityRequest, opts ...grpc.CallOption) (*JoinCommunityResponse, error) {
	out := new(JoinCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/JoinCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) LeaveCommunity(ctx context.Context, in *LeaveCommunityRequest, opts ...grpc.CallOption) (*LeaveCommunityResponse, error) {
	out := new(LeaveCommunityResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/LeaveCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) CreateCommunityEvent(ctx context.Context, in *CreateCommunityEventRequest, opts ...grpc.CallOption) (*CreateCommunityEventResponse, error) {
	out := new(CreateCommunityEventResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/CreateCommunityEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) ListCommunityEvents(ctx context.Context, in *ListCommunityEventsRequest, opts ...grpc.CallOption) (*ListCommunityEventsResponse, error) {
	out := new(ListCommunityEventsResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/ListCommunityEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) CreateCommunityForumPost(ctx context.Context, in *CreateCommunityForumPostRequest, opts ...grpc.CallOption) (*CreateCommunityForumPostRespnse, error) {
	out := new(CreateCommunityForumPostRespnse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/CreateCommunityForumPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) ListCommunityForumPosts(ctx context.Context, in *ListCommunityForumPostsRequest, opts ...grpc.CallOption) (*ListCommunityForumPostsResponse, error) {
	out := new(ListCommunityForumPostsResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/ListCommunityForumPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) AddForumPostComment(ctx context.Context, in *AddForumPostCommentRequest, opts ...grpc.CallOption) (*AddForumPostCommentResponse, error) {
	out := new(AddForumPostCommentResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/AddForumPostComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comunityServiceClient) ListForumPostComments(ctx context.Context, in *ListForumPostCommentsRequest, opts ...grpc.CallOption) (*ListForumPostCommentsResponse, error) {
	out := new(ListForumPostCommentsResponse)
	err := c.cc.Invoke(ctx, "/comunity.ComunityService/ListForumPostComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComunityServiceServer is the server API for ComunityService service.
// All implementations must embed UnimplementedComunityServiceServer
// for forward compatibility
type ComunityServiceServer interface {
	CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error)
	GetCommunity(context.Context, *GetCommunityRequest) (*GetCommunityResponse, error)
	UpdateCommunity(context.Context, *UpdateCommunityRequest) (*UpdateCommunityResponse, error)
	DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error)
	ListCommunities(context.Context, *ListCommunitiesRequest) (*ListCommunitiesResponse, error)
	JoinCommunity(context.Context, *JoinCommunityRequest) (*JoinCommunityResponse, error)
	LeaveCommunity(context.Context, *LeaveCommunityRequest) (*LeaveCommunityResponse, error)
	CreateCommunityEvent(context.Context, *CreateCommunityEventRequest) (*CreateCommunityEventResponse, error)
	ListCommunityEvents(context.Context, *ListCommunityEventsRequest) (*ListCommunityEventsResponse, error)
	CreateCommunityForumPost(context.Context, *CreateCommunityForumPostRequest) (*CreateCommunityForumPostRespnse, error)
	ListCommunityForumPosts(context.Context, *ListCommunityForumPostsRequest) (*ListCommunityForumPostsResponse, error)
	AddForumPostComment(context.Context, *AddForumPostCommentRequest) (*AddForumPostCommentResponse, error)
	ListForumPostComments(context.Context, *ListForumPostCommentsRequest) (*ListForumPostCommentsResponse, error)
	mustEmbedUnimplementedComunityServiceServer()
}

// UnimplementedComunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComunityServiceServer struct {
}

func (UnimplementedComunityServiceServer) CreateCommunity(context.Context, *CreateCommunityRequest) (*CreateCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedComunityServiceServer) GetCommunity(context.Context, *GetCommunityRequest) (*GetCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunity not implemented")
}
func (UnimplementedComunityServiceServer) UpdateCommunity(context.Context, *UpdateCommunityRequest) (*UpdateCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunity not implemented")
}
func (UnimplementedComunityServiceServer) DeleteCommunity(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunity not implemented")
}
func (UnimplementedComunityServiceServer) ListCommunities(context.Context, *ListCommunitiesRequest) (*ListCommunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunities not implemented")
}
func (UnimplementedComunityServiceServer) JoinCommunity(context.Context, *JoinCommunityRequest) (*JoinCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunity not implemented")
}
func (UnimplementedComunityServiceServer) LeaveCommunity(context.Context, *LeaveCommunityRequest) (*LeaveCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCommunity not implemented")
}
func (UnimplementedComunityServiceServer) CreateCommunityEvent(context.Context, *CreateCommunityEventRequest) (*CreateCommunityEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunityEvent not implemented")
}
func (UnimplementedComunityServiceServer) ListCommunityEvents(context.Context, *ListCommunityEventsRequest) (*ListCommunityEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunityEvents not implemented")
}
func (UnimplementedComunityServiceServer) CreateCommunityForumPost(context.Context, *CreateCommunityForumPostRequest) (*CreateCommunityForumPostRespnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunityForumPost not implemented")
}
func (UnimplementedComunityServiceServer) ListCommunityForumPosts(context.Context, *ListCommunityForumPostsRequest) (*ListCommunityForumPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunityForumPosts not implemented")
}
func (UnimplementedComunityServiceServer) AddForumPostComment(context.Context, *AddForumPostCommentRequest) (*AddForumPostCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddForumPostComment not implemented")
}
func (UnimplementedComunityServiceServer) ListForumPostComments(context.Context, *ListForumPostCommentsRequest) (*ListForumPostCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListForumPostComments not implemented")
}
func (UnimplementedComunityServiceServer) mustEmbedUnimplementedComunityServiceServer() {}

// UnsafeComunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComunityServiceServer will
// result in compilation errors.
type UnsafeComunityServiceServer interface {
	mustEmbedUnimplementedComunityServiceServer()
}

func RegisterComunityServiceServer(s grpc.ServiceRegistrar, srv ComunityServiceServer) {
	s.RegisterService(&ComunityService_ServiceDesc, srv)
}

func _ComunityService_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/CreateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).CreateCommunity(ctx, req.(*CreateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_GetCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).GetCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/GetCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).GetCommunity(ctx, req.(*GetCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_UpdateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).UpdateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/UpdateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).UpdateCommunity(ctx, req.(*UpdateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_DeleteCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).DeleteCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/DeleteCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).DeleteCommunity(ctx, req.(*DeleteCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_ListCommunities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).ListCommunities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/ListCommunities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).ListCommunities(ctx, req.(*ListCommunitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_JoinCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).JoinCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/JoinCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).JoinCommunity(ctx, req.(*JoinCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_LeaveCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).LeaveCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/LeaveCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).LeaveCommunity(ctx, req.(*LeaveCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_CreateCommunityEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).CreateCommunityEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/CreateCommunityEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).CreateCommunityEvent(ctx, req.(*CreateCommunityEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_ListCommunityEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunityEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).ListCommunityEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/ListCommunityEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).ListCommunityEvents(ctx, req.(*ListCommunityEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_CreateCommunityForumPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityForumPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).CreateCommunityForumPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/CreateCommunityForumPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).CreateCommunityForumPost(ctx, req.(*CreateCommunityForumPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_ListCommunityForumPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunityForumPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).ListCommunityForumPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/ListCommunityForumPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).ListCommunityForumPosts(ctx, req.(*ListCommunityForumPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_AddForumPostComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddForumPostCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).AddForumPostComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/AddForumPostComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).AddForumPostComment(ctx, req.(*AddForumPostCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComunityService_ListForumPostComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListForumPostCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComunityServiceServer).ListForumPostComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comunity.ComunityService/ListForumPostComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComunityServiceServer).ListForumPostComments(ctx, req.(*ListForumPostCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComunityService_ServiceDesc is the grpc.ServiceDesc for ComunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comunity.ComunityService",
	HandlerType: (*ComunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunity",
			Handler:    _ComunityService_CreateCommunity_Handler,
		},
		{
			MethodName: "GetCommunity",
			Handler:    _ComunityService_GetCommunity_Handler,
		},
		{
			MethodName: "UpdateCommunity",
			Handler:    _ComunityService_UpdateCommunity_Handler,
		},
		{
			MethodName: "DeleteCommunity",
			Handler:    _ComunityService_DeleteCommunity_Handler,
		},
		{
			MethodName: "ListCommunities",
			Handler:    _ComunityService_ListCommunities_Handler,
		},
		{
			MethodName: "JoinCommunity",
			Handler:    _ComunityService_JoinCommunity_Handler,
		},
		{
			MethodName: "LeaveCommunity",
			Handler:    _ComunityService_LeaveCommunity_Handler,
		},
		{
			MethodName: "CreateCommunityEvent",
			Handler:    _ComunityService_CreateCommunityEvent_Handler,
		},
		{
			MethodName: "ListCommunityEvents",
			Handler:    _ComunityService_ListCommunityEvents_Handler,
		},
		{
			MethodName: "CreateCommunityForumPost",
			Handler:    _ComunityService_CreateCommunityForumPost_Handler,
		},
		{
			MethodName: "ListCommunityForumPosts",
			Handler:    _ComunityService_ListCommunityForumPosts_Handler,
		},
		{
			MethodName: "AddForumPostComment",
			Handler:    _ComunityService_AddForumPostComment_Handler,
		},
		{
			MethodName: "ListForumPostComments",
			Handler:    _ComunityService_ListForumPostComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comunitiy.proto",
}
