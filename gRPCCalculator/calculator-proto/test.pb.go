// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package calculator

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

type Numbers struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (m *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(m, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Numbers) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=Created,proto3" json:"Created,omitempty"`
	Result               float64  `protobuf:"fixed64,2,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "calculator.Numbers")
	proto.RegisterType((*Response)(nil), "calculator.Response")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c,
	0xc9, 0x2f, 0x52, 0x52, 0xe5, 0x62, 0xf7, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0x2a, 0x16, 0xe2, 0xe1,
	0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98,
	0x20, 0xbc, 0x24, 0x25, 0x1b, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x09, 0x2e, 0x76, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0x14, 0xb0, 0x6a, 0x8e, 0x20, 0x18, 0x57,
	0x48, 0x8c, 0x8b, 0x2d, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x04, 0xac, 0x91, 0x31, 0x08, 0xca, 0x33,
	0xfa, 0xc0, 0xc8, 0x25, 0xe8, 0x0c, 0xb7, 0x33, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8,
	0x88, 0x8b, 0xd9, 0x31, 0x25, 0x45, 0x48, 0x58, 0x0f, 0xe1, 0x1c, 0x3d, 0xa8, 0x5b, 0xa4, 0x44,
	0x90, 0x05, 0x61, 0x36, 0x2b, 0x31, 0x08, 0x99, 0x73, 0x71, 0x04, 0x97, 0x26, 0x95, 0x14, 0x25,
	0x26, 0x97, 0x90, 0xac, 0xd1, 0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x20, 0xa7, 0x92, 0x34, 0x8d, 0xa6,
	0x5c, 0x6c, 0x2e, 0x99, 0x65, 0x99, 0x29, 0xa9, 0x24, 0x69, 0x4b, 0x62, 0x03, 0x07, 0xb5, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x19, 0x52, 0x8e, 0xd9, 0x78, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error)
	Subtract(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error)
	Divide(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Add(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Subtract(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Multiply(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Divide(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Add(context.Context, *Numbers) (*Response, error)
	Subtract(context.Context, *Numbers) (*Response, error)
	Multiply(context.Context, *Numbers) (*Response, error)
	Divide(context.Context, *Numbers) (*Response, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Add(ctx context.Context, req *Numbers) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCalculatorServiceServer) Subtract(ctx context.Context, req *Numbers) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (*UnimplementedCalculatorServiceServer) Multiply(ctx context.Context, req *Numbers) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (*UnimplementedCalculatorServiceServer) Divide(ctx context.Context, req *Numbers) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Add(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Subtract(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Multiply(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Divide(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalculatorService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _CalculatorService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CalculatorService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _CalculatorService_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}