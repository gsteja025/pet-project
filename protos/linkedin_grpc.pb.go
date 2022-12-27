// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: linkedin.proto

package petproject

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

// LinkedinDatabaseCrudClient is the client API for LinkedinDatabaseCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkedinDatabaseCrudClient interface {
	Createpost(ctx context.Context, in *NewPost, opts ...grpc.CallOption) (*Post, error)
	GetConnectedUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*Users, error)
	GetPostComments(ctx context.Context, in *Post, opts ...grpc.CallOption) (LinkedinDatabaseCrud_GetPostCommentsClient, error)
	GetPostLikes(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Users, error)
	ConnectWithOtherUser(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*Emptyresponse, error)
	LikeOtherPost(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Emptyresponse, error)
	SearchUser(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Users, error)
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
}

type linkedinDatabaseCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkedinDatabaseCrudClient(cc grpc.ClientConnInterface) LinkedinDatabaseCrudClient {
	return &linkedinDatabaseCrudClient{cc}
}

func (c *linkedinDatabaseCrudClient) Createpost(ctx context.Context, in *NewPost, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/Createpost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) GetConnectedUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/GetConnectedUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) GetPostComments(ctx context.Context, in *Post, opts ...grpc.CallOption) (LinkedinDatabaseCrud_GetPostCommentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LinkedinDatabaseCrud_ServiceDesc.Streams[0], "/petproject.LinkedinDatabaseCrud/GetPostComments", opts...)
	if err != nil {
		return nil, err
	}
	x := &linkedinDatabaseCrudGetPostCommentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LinkedinDatabaseCrud_GetPostCommentsClient interface {
	Recv() (*Comment, error)
	grpc.ClientStream
}

type linkedinDatabaseCrudGetPostCommentsClient struct {
	grpc.ClientStream
}

func (x *linkedinDatabaseCrudGetPostCommentsClient) Recv() (*Comment, error) {
	m := new(Comment)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *linkedinDatabaseCrudClient) GetPostLikes(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/GetPostLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) ConnectWithOtherUser(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*Emptyresponse, error) {
	out := new(Emptyresponse)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/ConnectWithOtherUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) LikeOtherPost(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Emptyresponse, error) {
	out := new(Emptyresponse)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/LikeOtherPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) SearchUser(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedinDatabaseCrudClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/petproject.LinkedinDatabaseCrud/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkedinDatabaseCrudServer is the server API for LinkedinDatabaseCrud service.
// All implementations must embed UnimplementedLinkedinDatabaseCrudServer
// for forward compatibility
type LinkedinDatabaseCrudServer interface {
	Createpost(context.Context, *NewPost) (*Post, error)
	GetConnectedUsers(context.Context, *User) (*Users, error)
	GetPostComments(*Post, LinkedinDatabaseCrud_GetPostCommentsServer) error
	GetPostLikes(context.Context, *Post) (*Users, error)
	ConnectWithOtherUser(context.Context, *ConnectionRequest) (*Emptyresponse, error)
	LikeOtherPost(context.Context, *Request) (*Emptyresponse, error)
	SearchUser(context.Context, *SearchRequest) (*Users, error)
	CreateComment(context.Context, *Comment) (*Comment, error)
	mustEmbedUnimplementedLinkedinDatabaseCrudServer()
}

// UnimplementedLinkedinDatabaseCrudServer must be embedded to have forward compatible implementations.
type UnimplementedLinkedinDatabaseCrudServer struct {
}

func (UnimplementedLinkedinDatabaseCrudServer) Createpost(context.Context, *NewPost) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Createpost not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) GetConnectedUsers(context.Context, *User) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectedUsers not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) GetPostComments(*Post, LinkedinDatabaseCrud_GetPostCommentsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPostComments not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) GetPostLikes(context.Context, *Post) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostLikes not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) ConnectWithOtherUser(context.Context, *ConnectionRequest) (*Emptyresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectWithOtherUser not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) LikeOtherPost(context.Context, *Request) (*Emptyresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeOtherPost not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) SearchUser(context.Context, *SearchRequest) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) CreateComment(context.Context, *Comment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedLinkedinDatabaseCrudServer) mustEmbedUnimplementedLinkedinDatabaseCrudServer() {}

// UnsafeLinkedinDatabaseCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkedinDatabaseCrudServer will
// result in compilation errors.
type UnsafeLinkedinDatabaseCrudServer interface {
	mustEmbedUnimplementedLinkedinDatabaseCrudServer()
}

func RegisterLinkedinDatabaseCrudServer(s grpc.ServiceRegistrar, srv LinkedinDatabaseCrudServer) {
	s.RegisterService(&LinkedinDatabaseCrud_ServiceDesc, srv)
}

func _LinkedinDatabaseCrud_Createpost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).Createpost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/Createpost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).Createpost(ctx, req.(*NewPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_GetConnectedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).GetConnectedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/GetConnectedUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).GetConnectedUsers(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_GetPostComments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Post)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LinkedinDatabaseCrudServer).GetPostComments(m, &linkedinDatabaseCrudGetPostCommentsServer{stream})
}

type LinkedinDatabaseCrud_GetPostCommentsServer interface {
	Send(*Comment) error
	grpc.ServerStream
}

type linkedinDatabaseCrudGetPostCommentsServer struct {
	grpc.ServerStream
}

func (x *linkedinDatabaseCrudGetPostCommentsServer) Send(m *Comment) error {
	return x.ServerStream.SendMsg(m)
}

func _LinkedinDatabaseCrud_GetPostLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).GetPostLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/GetPostLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).GetPostLikes(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_ConnectWithOtherUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).ConnectWithOtherUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/ConnectWithOtherUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).ConnectWithOtherUser(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_LikeOtherPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).LikeOtherPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/LikeOtherPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).LikeOtherPost(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).SearchUser(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkedinDatabaseCrud_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedinDatabaseCrudServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petproject.LinkedinDatabaseCrud/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedinDatabaseCrudServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkedinDatabaseCrud_ServiceDesc is the grpc.ServiceDesc for LinkedinDatabaseCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkedinDatabaseCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petproject.LinkedinDatabaseCrud",
	HandlerType: (*LinkedinDatabaseCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Createpost",
			Handler:    _LinkedinDatabaseCrud_Createpost_Handler,
		},
		{
			MethodName: "GetConnectedUsers",
			Handler:    _LinkedinDatabaseCrud_GetConnectedUsers_Handler,
		},
		{
			MethodName: "GetPostLikes",
			Handler:    _LinkedinDatabaseCrud_GetPostLikes_Handler,
		},
		{
			MethodName: "ConnectWithOtherUser",
			Handler:    _LinkedinDatabaseCrud_ConnectWithOtherUser_Handler,
		},
		{
			MethodName: "LikeOtherPost",
			Handler:    _LinkedinDatabaseCrud_LikeOtherPost_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _LinkedinDatabaseCrud_SearchUser_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _LinkedinDatabaseCrud_CreateComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPostComments",
			Handler:       _LinkedinDatabaseCrud_GetPostComments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "linkedin.proto",
}
