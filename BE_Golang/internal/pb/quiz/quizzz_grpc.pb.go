// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: quizzz.proto

package quiz

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

const (
	QuizzesService_FindAllQuiz_FullMethodName = "/proto.QuizzesService/FindAllQuiz"
)

// QuizzesServiceClient is the client API for QuizzesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuizzesServiceClient interface {
	FindAllQuiz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Quizzes, error)
}

type quizzesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuizzesServiceClient(cc grpc.ClientConnInterface) QuizzesServiceClient {
	return &quizzesServiceClient{cc}
}

func (c *quizzesServiceClient) FindAllQuiz(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Quizzes, error) {
	out := new(Quizzes)
	err := c.cc.Invoke(ctx, QuizzesService_FindAllQuiz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizzesServiceServer is the server API for QuizzesService service.
// All implementations must embed UnimplementedQuizzesServiceServer
// for forward compatibility
type QuizzesServiceServer interface {
	FindAllQuiz(context.Context, *Empty) (*Quizzes, error)
	mustEmbedUnimplementedQuizzesServiceServer()
}

// UnimplementedQuizzesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuizzesServiceServer struct {
}

func (UnimplementedQuizzesServiceServer) FindAllQuiz(context.Context, *Empty) (*Quizzes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllQuiz not implemented")
}
func (UnimplementedQuizzesServiceServer) mustEmbedUnimplementedQuizzesServiceServer() {}

// UnsafeQuizzesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuizzesServiceServer will
// result in compilation errors.
type UnsafeQuizzesServiceServer interface {
	mustEmbedUnimplementedQuizzesServiceServer()
}

func RegisterQuizzesServiceServer(s grpc.ServiceRegistrar, srv QuizzesServiceServer) {
	s.RegisterService(&QuizzesService_ServiceDesc, srv)
}

func _QuizzesService_FindAllQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizzesServiceServer).FindAllQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuizzesService_FindAllQuiz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizzesServiceServer).FindAllQuiz(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// QuizzesService_ServiceDesc is the grpc.ServiceDesc for QuizzesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuizzesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuizzesService",
	HandlerType: (*QuizzesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAllQuiz",
			Handler:    _QuizzesService_FindAllQuiz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quizzz.proto",
}
