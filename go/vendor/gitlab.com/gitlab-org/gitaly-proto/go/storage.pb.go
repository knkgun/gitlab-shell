// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package gitaly

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeleteAllRepositoriesRequest struct {
	StorageName string `protobuf:"bytes,1,opt,name=storage_name,json=storageName" json:"storage_name,omitempty"`
}

func (m *DeleteAllRepositoriesRequest) Reset()                    { *m = DeleteAllRepositoriesRequest{} }
func (m *DeleteAllRepositoriesRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteAllRepositoriesRequest) ProtoMessage()               {}
func (*DeleteAllRepositoriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *DeleteAllRepositoriesRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

type DeleteAllRepositoriesResponse struct {
}

func (m *DeleteAllRepositoriesResponse) Reset()                    { *m = DeleteAllRepositoriesResponse{} }
func (m *DeleteAllRepositoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteAllRepositoriesResponse) ProtoMessage()               {}
func (*DeleteAllRepositoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func init() {
	proto.RegisterType((*DeleteAllRepositoriesRequest)(nil), "gitaly.DeleteAllRepositoriesRequest")
	proto.RegisterType((*DeleteAllRepositoriesResponse)(nil), "gitaly.DeleteAllRepositoriesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StorageService service

type StorageServiceClient interface {
	DeleteAllRepositories(ctx context.Context, in *DeleteAllRepositoriesRequest, opts ...grpc.CallOption) (*DeleteAllRepositoriesResponse, error)
}

type storageServiceClient struct {
	cc *grpc.ClientConn
}

func NewStorageServiceClient(cc *grpc.ClientConn) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) DeleteAllRepositories(ctx context.Context, in *DeleteAllRepositoriesRequest, opts ...grpc.CallOption) (*DeleteAllRepositoriesResponse, error) {
	out := new(DeleteAllRepositoriesResponse)
	err := grpc.Invoke(ctx, "/gitaly.StorageService/DeleteAllRepositories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StorageService service

type StorageServiceServer interface {
	DeleteAllRepositories(context.Context, *DeleteAllRepositoriesRequest) (*DeleteAllRepositoriesResponse, error)
}

func RegisterStorageServiceServer(s *grpc.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_DeleteAllRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DeleteAllRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.StorageService/DeleteAllRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DeleteAllRepositories(ctx, req.(*DeleteAllRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteAllRepositories",
			Handler:    _StorageService_DeleteAllRepositories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0x49, 0xcc,
	0xa9, 0x54, 0x72, 0xe4, 0x92, 0x71, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x75, 0xcc, 0xc9, 0x09, 0x4a,
	0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xca, 0x4c, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x52, 0xe4, 0xe2, 0x81, 0x6a, 0x8c, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0xe2, 0x86, 0x8a, 0xf9, 0x25, 0xe6, 0xa6, 0x2a, 0xc9, 0x73, 0xc9, 0xe2, 0x30,
	0xa2, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xa8, 0x82, 0x8b, 0x2f, 0x18, 0xa2, 0x3e, 0x38, 0xb5,
	0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x8d, 0x4b, 0x14, 0xab, 0x16, 0x21, 0x15, 0x3d, 0x88, 0xbb,
	0xf4, 0xf0, 0x39, 0x4a, 0x4a, 0x95, 0x80, 0x2a, 0x88, 0xbd, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xcf,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x21, 0xd8, 0x88, 0xfd, 0x00, 0x00, 0x00,
}