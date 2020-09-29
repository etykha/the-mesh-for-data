// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy_manager_service.proto

package connectors

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("policy_manager_service.proto", fileDescriptor_76b6ed75392807da)
}

var fileDescriptor_76b6ed75392807da = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xbf, 0x0a, 0xc2, 0x40,
	0x0c, 0x87, 0x9d, 0x1c, 0x4e, 0x5c, 0x4a, 0xa7, 0x62, 0x1d, 0x7c, 0x80, 0x1b, 0xf4, 0x09, 0xfc,
	0x03, 0x4e, 0x82, 0xe0, 0xe0, 0x58, 0xce, 0x18, 0x34, 0xd0, 0x26, 0xe7, 0x25, 0x8a, 0xbe, 0xbd,
	0xd0, 0x0a, 0x8a, 0x9d, 0xbf, 0x2f, 0xf9, 0x7e, 0x6e, 0x12, 0xa5, 0x26, 0x78, 0x55, 0x4d, 0xe0,
	0x70, 0xc1, 0x54, 0x29, 0xa6, 0x07, 0x01, 0xfa, 0x98, 0xc4, 0x24, 0x73, 0x20, 0xcc, 0x08, 0x26,
	0x49, 0x8b, 0x7f, 0x33, 0xe1, 0xed, 0x8e, 0x6a, 0x9d, 0x59, 0x94, 0x3d, 0xaa, 0x51, 0x58, 0x3f,
	0x8f, 0xe6, 0xe2, 0xf2, 0x7d, 0x2b, 0xec, 0x3a, 0x7e, 0xe8, 0x32, 0xd9, 0xd1, 0xe5, 0x5b, 0xb4,
	0x16, 0x11, 0xea, 0x06, 0x81, 0x94, 0x84, 0x35, 0x9b, 0xfa, 0x6f, 0xd9, 0x2f, 0x63, 0xac, 0x09,
	0x82, 0x91, 0xf0, 0x5a, 0xd8, 0xf0, 0x69, 0x45, 0xf9, 0xcb, 0x7b, 0xe7, 0xb3, 0xc1, 0x6a, 0xec,
	0x46, 0x20, 0x8d, 0x3f, 0x07, 0x6b, 0x50, 0xaf, 0xa7, 0x61, 0x3b, 0x63, 0xf1, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x20, 0xe2, 0x54, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PolicyManagerServiceClient is the client API for PolicyManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolicyManagerServiceClient interface {
	GetPoliciesDecisions(ctx context.Context, in *ApplicationContext, opts ...grpc.CallOption) (*PoliciesDecisions, error)
}

type policyManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyManagerServiceClient(cc grpc.ClientConnInterface) PolicyManagerServiceClient {
	return &policyManagerServiceClient{cc}
}

func (c *policyManagerServiceClient) GetPoliciesDecisions(ctx context.Context, in *ApplicationContext, opts ...grpc.CallOption) (*PoliciesDecisions, error) {
	out := new(PoliciesDecisions)
	err := c.cc.Invoke(ctx, "/connectors.PolicyManagerService/GetPoliciesDecisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyManagerServiceServer is the server API for PolicyManagerService service.
type PolicyManagerServiceServer interface {
	GetPoliciesDecisions(context.Context, *ApplicationContext) (*PoliciesDecisions, error)
}

// UnimplementedPolicyManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPolicyManagerServiceServer struct {
}

func (*UnimplementedPolicyManagerServiceServer) GetPoliciesDecisions(ctx context.Context, req *ApplicationContext) (*PoliciesDecisions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoliciesDecisions not implemented")
}

func RegisterPolicyManagerServiceServer(s *grpc.Server, srv PolicyManagerServiceServer) {
	s.RegisterService(&_PolicyManagerService_serviceDesc, srv)
}

func _PolicyManagerService_GetPoliciesDecisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyManagerServiceServer).GetPoliciesDecisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.PolicyManagerService/GetPoliciesDecisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyManagerServiceServer).GetPoliciesDecisions(ctx, req.(*ApplicationContext))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.PolicyManagerService",
	HandlerType: (*PolicyManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPoliciesDecisions",
			Handler:    _PolicyManagerService_GetPoliciesDecisions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy_manager_service.proto",
}