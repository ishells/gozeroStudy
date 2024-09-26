// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: recommend.proto

package recommend

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Recommend_Ping_FullMethodName = "/recommend.Recommend/Ping"
)

// RecommendClient is the client API for Recommend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommendClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type recommendClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendClient(cc grpc.ClientConnInterface) RecommendClient {
	return &recommendClient{cc}
}

func (c *recommendClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, Recommend_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendServer is the server API for Recommend service.
// All implementations must embed UnimplementedRecommendServer
// for forward compatibility.
type RecommendServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedRecommendServer()
}

// UnimplementedRecommendServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecommendServer struct{}

func (UnimplementedRecommendServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedRecommendServer) mustEmbedUnimplementedRecommendServer() {}
func (UnimplementedRecommendServer) testEmbeddedByValue()                   {}

// UnsafeRecommendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendServer will
// result in compilation errors.
type UnsafeRecommendServer interface {
	mustEmbedUnimplementedRecommendServer()
}

func RegisterRecommendServer(s grpc.ServiceRegistrar, srv RecommendServer) {
	// If the following call pancis, it indicates UnimplementedRecommendServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Recommend_ServiceDesc, srv)
}

func _Recommend_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Recommend_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Recommend_ServiceDesc is the grpc.ServiceDesc for Recommend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Recommend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.Recommend",
	HandlerType: (*RecommendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Recommend_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recommend.proto",
}
