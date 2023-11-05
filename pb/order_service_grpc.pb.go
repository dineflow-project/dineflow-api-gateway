// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: order_service.proto

package pb

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (OrderService_GetOrdersClient, error)
	GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdRequest, opts ...grpc.CallOption) (OrderService_GetOrdersByUserIdClient, error)
	GetOrdersByVendorId(ctx context.Context, in *GetOrdersByVendorIdRequest, opts ...grpc.CallOption) (OrderService_GetOrdersByVendorIdClient, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	DeleteOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (OrderService_GetOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], "/pb.OrderService/GetOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceGetOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) GetOrdersByUserId(ctx context.Context, in *GetOrdersByUserIdRequest, opts ...grpc.CallOption) (OrderService_GetOrdersByUserIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[1], "/pb.OrderService/GetOrdersByUserId", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrdersByUserIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetOrdersByUserIdClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceGetOrdersByUserIdClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrdersByUserIdClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) GetOrdersByVendorId(ctx context.Context, in *GetOrdersByVendorIdRequest, opts ...grpc.CallOption) (OrderService_GetOrdersByVendorIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[2], "/pb.OrderService/GetOrdersByVendorId", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceGetOrdersByVendorIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_GetOrdersByVendorIdClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceGetOrdersByVendorIdClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrdersByVendorIdClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, "/pb.OrderService/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error)
	GetOrder(context.Context, *OrderRequest) (*OrderResponse, error)
	GetOrders(*GetOrdersRequest, OrderService_GetOrdersServer) error
	GetOrdersByUserId(*GetOrdersByUserIdRequest, OrderService_GetOrdersByUserIdServer) error
	GetOrdersByVendorId(*GetOrdersByVendorIdRequest, OrderService_GetOrdersByVendorIdServer) error
	UpdateOrder(context.Context, *UpdateOrderRequest) (*OrderResponse, error)
	DeleteOrder(context.Context, *OrderRequest) (*DeleteOrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServiceServer) GetOrders(*GetOrdersRequest, OrderService_GetOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrderServiceServer) GetOrdersByUserId(*GetOrdersByUserIdRequest, OrderService_GetOrdersByUserIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrdersByUserId not implemented")
}
func (UnimplementedOrderServiceServer) GetOrdersByVendorId(*GetOrdersByVendorIdRequest, OrderService_GetOrdersByVendorIdServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOrdersByVendorId not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrder(context.Context, *OrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOrdersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetOrders(m, &orderServiceGetOrdersServer{stream})
}

type OrderService_GetOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceGetOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_GetOrdersByUserId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOrdersByUserIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetOrdersByUserId(m, &orderServiceGetOrdersByUserIdServer{stream})
}

type OrderService_GetOrdersByUserIdServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceGetOrdersByUserIdServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrdersByUserIdServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_GetOrdersByVendorId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetOrdersByVendorIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).GetOrdersByVendorId(m, &orderServiceGetOrdersByVendorIdServer{stream})
}

type OrderService_GetOrdersByVendorIdServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceGetOrdersByVendorIdServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrdersByVendorIdServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.OrderService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetOrders",
			Handler:       _OrderService_GetOrders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrdersByUserId",
			Handler:       _OrderService_GetOrdersByUserId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOrdersByVendorId",
			Handler:       _OrderService_GetOrdersByVendorId_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "order_service.proto",
}
