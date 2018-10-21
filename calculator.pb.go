// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator.proto

package calculator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the first and second arguments for the operation.
type SimpleOperationRequest struct {
	Arg1                 float64  `protobuf:"fixed64,1,opt,name=arg1,proto3" json:"arg1,omitempty"`
	Arg2                 float64  `protobuf:"fixed64,2,opt,name=arg2,proto3" json:"arg2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleOperationRequest) Reset()         { *m = SimpleOperationRequest{} }
func (m *SimpleOperationRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleOperationRequest) ProtoMessage()    {}
func (*SimpleOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{0}
}

func (m *SimpleOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleOperationRequest.Unmarshal(m, b)
}
func (m *SimpleOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleOperationRequest.Marshal(b, m, deterministic)
}
func (m *SimpleOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleOperationRequest.Merge(m, src)
}
func (m *SimpleOperationRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleOperationRequest.Size(m)
}
func (m *SimpleOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleOperationRequest proto.InternalMessageInfo

func (m *SimpleOperationRequest) GetArg1() float64 {
	if m != nil {
		return m.Arg1
	}
	return 0
}

func (m *SimpleOperationRequest) GetArg2() float64 {
	if m != nil {
		return m.Arg2
	}
	return 0
}

// The response message containing the result
type SimpleOperationReply struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleOperationReply) Reset()         { *m = SimpleOperationReply{} }
func (m *SimpleOperationReply) String() string { return proto.CompactTextString(m) }
func (*SimpleOperationReply) ProtoMessage()    {}
func (*SimpleOperationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c686ea360062a8cf, []int{1}
}

func (m *SimpleOperationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleOperationReply.Unmarshal(m, b)
}
func (m *SimpleOperationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleOperationReply.Marshal(b, m, deterministic)
}
func (m *SimpleOperationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleOperationReply.Merge(m, src)
}
func (m *SimpleOperationReply) XXX_Size() int {
	return xxx_messageInfo_SimpleOperationReply.Size(m)
}
func (m *SimpleOperationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleOperationReply.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleOperationReply proto.InternalMessageInfo

func (m *SimpleOperationReply) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SimpleOperationRequest)(nil), "calculator.SimpleOperationRequest")
	proto.RegisterType((*SimpleOperationReply)(nil), "calculator.SimpleOperationReply")
}

func init() { proto.RegisterFile("calculator.proto", fileDescriptor_c686ea360062a8cf) }

var fileDescriptor_c686ea360062a8cf = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0xcc, 0x49,
	0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x39, 0x70, 0x89, 0x05, 0x67, 0xe6, 0x16, 0xe4, 0xa4, 0xfa, 0x17, 0xa4, 0x16, 0x25, 0x96,
	0x64, 0xe6, 0xe7, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x24, 0x16,
	0xa5, 0x1b, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x06, 0x81, 0xd9, 0x50, 0x31, 0x23, 0x09, 0x26,
	0xb8, 0x98, 0x91, 0x92, 0x1e, 0x97, 0x08, 0x86, 0x09, 0x05, 0x39, 0x95, 0x42, 0x62, 0x5c, 0x6c,
	0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x50, 0x13, 0xa0, 0x3c, 0xa3, 0xab, 0x4c, 0x5c, 0x5c, 0xce,
	0x70, 0x07, 0x08, 0xf9, 0x72, 0x31, 0x3b, 0xa6, 0xa4, 0x08, 0x29, 0xe9, 0x21, 0x39, 0x13, 0xbb,
	0x8b, 0xa4, 0x14, 0xf0, 0xaa, 0x29, 0xc8, 0xa9, 0x54, 0x62, 0x10, 0x0a, 0xe2, 0xe2, 0x08, 0x2e,
	0x4d, 0x2a, 0x29, 0x4a, 0x4c, 0x2e, 0xa1, 0xa6, 0x99, 0xbe, 0xa5, 0x39, 0x25, 0x99, 0x20, 0x5f,
	0x51, 0xcb, 0xcc, 0x00, 0x2e, 0x36, 0x97, 0xcc, 0xb2, 0xcc, 0x94, 0x54, 0x6a, 0x99, 0x98, 0xc4,
	0x06, 0x8e, 0x5c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xfe, 0x66, 0xe8, 0xf0, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error)
	Subtract(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error)
	Multiply(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error)
	Divide(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error) {
	out := new(SimpleOperationReply)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Subtract(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error) {
	out := new(SimpleOperationReply)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Multiply(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error) {
	out := new(SimpleOperationReply)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) Divide(ctx context.Context, in *SimpleOperationRequest, opts ...grpc.CallOption) (*SimpleOperationReply, error) {
	out := new(SimpleOperationReply)
	err := c.cc.Invoke(ctx, "/calculator.Calculator/Divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Add(context.Context, *SimpleOperationRequest) (*SimpleOperationReply, error)
	Subtract(context.Context, *SimpleOperationRequest) (*SimpleOperationReply, error)
	Multiply(context.Context, *SimpleOperationRequest) (*SimpleOperationReply, error)
	Divide(context.Context, *SimpleOperationRequest) (*SimpleOperationReply, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Add(ctx, req.(*SimpleOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Subtract(ctx, req.(*SimpleOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Multiply(ctx, req.(*SimpleOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.Calculator/Divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Divide(ctx, req.(*SimpleOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Calculator_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _Calculator_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _Calculator_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _Calculator_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
