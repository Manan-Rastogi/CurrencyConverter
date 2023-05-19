// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/currencyExchange.proto

package currencyProto

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

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConverterClient interface {
	CurrencyConverter(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*Welcome, error)
	CurrencyConverterServerStreaming(ctx context.Context, in *CurrencyRequestList, opts ...grpc.CallOption) (Converter_CurrencyConverterServerStreamingClient, error)
	CurrencyConverterClientStreaming(ctx context.Context, opts ...grpc.CallOption) (Converter_CurrencyConverterClientStreamingClient, error)
	CurrencyConverterBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Converter_CurrencyConverterBidirectionalStreamingClient, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) CurrencyConverter(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*Welcome, error) {
	out := new(Welcome)
	err := c.cc.Invoke(ctx, "/currencyProto.Converter/CurrencyConverter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) CurrencyConverterServerStreaming(ctx context.Context, in *CurrencyRequestList, opts ...grpc.CallOption) (Converter_CurrencyConverterServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Converter_ServiceDesc.Streams[0], "/currencyProto.Converter/CurrencyConverterServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &converterCurrencyConverterServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Converter_CurrencyConverterServerStreamingClient interface {
	Recv() (*CurrencyResponse, error)
	grpc.ClientStream
}

type converterCurrencyConverterServerStreamingClient struct {
	grpc.ClientStream
}

func (x *converterCurrencyConverterServerStreamingClient) Recv() (*CurrencyResponse, error) {
	m := new(CurrencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *converterClient) CurrencyConverterClientStreaming(ctx context.Context, opts ...grpc.CallOption) (Converter_CurrencyConverterClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Converter_ServiceDesc.Streams[1], "/currencyProto.Converter/CurrencyConverterClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &converterCurrencyConverterClientStreamingClient{stream}
	return x, nil
}

type Converter_CurrencyConverterClientStreamingClient interface {
	Send(*CurrencyRequest) error
	CloseAndRecv() (*CurrencyResponseList, error)
	grpc.ClientStream
}

type converterCurrencyConverterClientStreamingClient struct {
	grpc.ClientStream
}

func (x *converterCurrencyConverterClientStreamingClient) Send(m *CurrencyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *converterCurrencyConverterClientStreamingClient) CloseAndRecv() (*CurrencyResponseList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CurrencyResponseList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *converterClient) CurrencyConverterBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (Converter_CurrencyConverterBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Converter_ServiceDesc.Streams[2], "/currencyProto.Converter/CurrencyConverterBidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &converterCurrencyConverterBidirectionalStreamingClient{stream}
	return x, nil
}

type Converter_CurrencyConverterBidirectionalStreamingClient interface {
	Send(*CurrencyRequest) error
	Recv() (*CurrencyResponse, error)
	grpc.ClientStream
}

type converterCurrencyConverterBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *converterCurrencyConverterBidirectionalStreamingClient) Send(m *CurrencyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *converterCurrencyConverterBidirectionalStreamingClient) Recv() (*CurrencyResponse, error) {
	m := new(CurrencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConverterServer is the server API for Converter service.
// All implementations must embed UnimplementedConverterServer
// for forward compatibility
type ConverterServer interface {
	CurrencyConverter(context.Context, *NoParam) (*Welcome, error)
	CurrencyConverterServerStreaming(*CurrencyRequestList, Converter_CurrencyConverterServerStreamingServer) error
	CurrencyConverterClientStreaming(Converter_CurrencyConverterClientStreamingServer) error
	CurrencyConverterBidirectionalStreaming(Converter_CurrencyConverterBidirectionalStreamingServer) error
	mustEmbedUnimplementedConverterServer()
}

// UnimplementedConverterServer must be embedded to have forward compatible implementations.
type UnimplementedConverterServer struct {
}

func (UnimplementedConverterServer) CurrencyConverter(context.Context, *NoParam) (*Welcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrencyConverter not implemented")
}
func (UnimplementedConverterServer) CurrencyConverterServerStreaming(*CurrencyRequestList, Converter_CurrencyConverterServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method CurrencyConverterServerStreaming not implemented")
}
func (UnimplementedConverterServer) CurrencyConverterClientStreaming(Converter_CurrencyConverterClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method CurrencyConverterClientStreaming not implemented")
}
func (UnimplementedConverterServer) CurrencyConverterBidirectionalStreaming(Converter_CurrencyConverterBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method CurrencyConverterBidirectionalStreaming not implemented")
}
func (UnimplementedConverterServer) mustEmbedUnimplementedConverterServer() {}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_CurrencyConverter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).CurrencyConverter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/currencyProto.Converter/CurrencyConverter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).CurrencyConverter(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_CurrencyConverterServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyRequestList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConverterServer).CurrencyConverterServerStreaming(m, &converterCurrencyConverterServerStreamingServer{stream})
}

type Converter_CurrencyConverterServerStreamingServer interface {
	Send(*CurrencyResponse) error
	grpc.ServerStream
}

type converterCurrencyConverterServerStreamingServer struct {
	grpc.ServerStream
}

func (x *converterCurrencyConverterServerStreamingServer) Send(m *CurrencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Converter_CurrencyConverterClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConverterServer).CurrencyConverterClientStreaming(&converterCurrencyConverterClientStreamingServer{stream})
}

type Converter_CurrencyConverterClientStreamingServer interface {
	SendAndClose(*CurrencyResponseList) error
	Recv() (*CurrencyRequest, error)
	grpc.ServerStream
}

type converterCurrencyConverterClientStreamingServer struct {
	grpc.ServerStream
}

func (x *converterCurrencyConverterClientStreamingServer) SendAndClose(m *CurrencyResponseList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *converterCurrencyConverterClientStreamingServer) Recv() (*CurrencyRequest, error) {
	m := new(CurrencyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Converter_CurrencyConverterBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConverterServer).CurrencyConverterBidirectionalStreaming(&converterCurrencyConverterBidirectionalStreamingServer{stream})
}

type Converter_CurrencyConverterBidirectionalStreamingServer interface {
	Send(*CurrencyResponse) error
	Recv() (*CurrencyRequest, error)
	grpc.ServerStream
}

type converterCurrencyConverterBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *converterCurrencyConverterBidirectionalStreamingServer) Send(m *CurrencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *converterCurrencyConverterBidirectionalStreamingServer) Recv() (*CurrencyRequest, error) {
	m := new(CurrencyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currencyProto.Converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrencyConverter",
			Handler:    _Converter_CurrencyConverter_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CurrencyConverterServerStreaming",
			Handler:       _Converter_CurrencyConverterServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CurrencyConverterClientStreaming",
			Handler:       _Converter_CurrencyConverterClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CurrencyConverterBidirectionalStreaming",
			Handler:       _Converter_CurrencyConverterBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/currencyExchange.proto",
}
