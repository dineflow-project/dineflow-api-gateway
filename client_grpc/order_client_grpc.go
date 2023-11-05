package client_grpc

import (
	"context"
	"io"
	"log"

	"api-gateway/model"
	pb "api-gateway/pb"
)

type OrderClient struct {
	client *pb.OrderServiceClient
}

type OrdergRPCClient interface {
	GetAllOrders(ctx context.Context) ([]*model.Order, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error)
	UpdateOrderByID(ctx context.Context, id string, order *model.Order) (*model.Order, error)
	DeleteOrderByID(ctx context.Context, id string) (*model.Order, error)
	GetOrderByVendorID(ctx context.Context, vid string) ([]*model.Order, error)
	GetOrderByUserID(ctx context.Context, uid string) ([]*model.Order, error)
}

func (t *OrderClient) GetAllOrders(ctx context.Context) ([]*model.Order, error) {
	stream, err := (*t.client).GetOrders(ctx, &pb.GetOrdersRequest{})
	if err != nil {
		return nil, err
	}
	var orders []*model.Order

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListOrders: %v", err)
		}

		order := &model.Order{
			ID:         res.Id,
			Status:     res.Status,
			OrderMenus: make([]*model.OrderMenu, 0),
			VendorId:   res.VendorId,
			Price:      res.Price,
			UserId:     res.UserId,
			CreateAt:   res.CreatedAt.AsTime(),
			UpdatedAt:  res.UpdatedAt.AsTime(),
		}
		for _, orderMenu := range res.OrderMenus {
			orderMenu := &model.OrderMenu{
				MenuId:  orderMenu.MenuId,
				Amount:  int(orderMenu.Amount),
				Price:   orderMenu.Price,
				Request: orderMenu.Request,
			}
			order.OrderMenus = append(order.OrderMenus, orderMenu)
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (t *OrderClient) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	res, err := (*t.client).GetOrder(ctx, &pb.OrderRequest{Id: id})
	if err != nil {
		return nil, err
	}
	order := &model.Order{
		ID:         res.Order.Id,
		Status:     res.Order.Status,
		OrderMenus: make([]*model.OrderMenu, 0),
		VendorId:   res.Order.VendorId,
		Price:      res.Order.Price,
		UserId:     res.Order.UserId,
		CreateAt:   res.Order.CreatedAt.AsTime(),
		UpdatedAt:  res.Order.UpdatedAt.AsTime(),
	}
	for _, orderMenu := range res.Order.OrderMenus {
		orderMenu := &model.OrderMenu{
			MenuId:  orderMenu.MenuId,
			Amount:  int(orderMenu.Amount),
			Price:   orderMenu.Price,
			Request: orderMenu.Request,
		}
		order.OrderMenus = append(order.OrderMenus, orderMenu)
	}
	return order, nil
}

func (t *OrderClient) CreateOrder(ctx context.Context, order *model.Order) (*model.Order, error) {
	orderMenus := make([]*pb.CreateOrderRequest_OrderMenu, 0)
	for _, menu := range order.OrderMenus {
		pbMenu := &pb.CreateOrderRequest_OrderMenu{
			MenuId:  menu.MenuId,
			Amount:  int32(menu.Amount),
			Price:   menu.Price,
			Request: menu.Request,
		}
		orderMenus = append(orderMenus, pbMenu)
	}

	res, err := (*t.client).CreateOrder(ctx, &pb.CreateOrderRequest{
		Status:     order.Status,
		VendorId:   order.VendorId,
		UserId:     order.UserId,
		OrderMenus: orderMenus,
	})
	if err != nil {
		return nil, err
	}
	newOrder := &model.Order{
		ID:         res.Order.Id,
		Status:     res.Order.Status,
		OrderMenus: make([]*model.OrderMenu, 0),
		VendorId:   res.Order.VendorId,
		Price:      res.Order.Price,
		UserId:     res.Order.UserId,
		CreateAt:   res.Order.CreatedAt.AsTime(),
		UpdatedAt:  res.Order.UpdatedAt.AsTime(),
	}
	for _, orderMenu := range res.Order.OrderMenus {
		orderMenu := &model.OrderMenu{
			MenuId:  orderMenu.MenuId,
			Amount:  int(orderMenu.Amount),
			Price:   orderMenu.Price,
			Request: orderMenu.Request,
		}
		order.OrderMenus = append(order.OrderMenus, orderMenu)
	}
	return newOrder, nil
}

func (t *OrderClient) UpdateOrderByID(ctx context.Context, id string, order *model.Order) (*model.Order, error) {
	orderMenus := make([]*pb.UpdateOrderRequest_OrderMenu, 0)
	for _, menu := range order.OrderMenus {
		pbMenu := &pb.UpdateOrderRequest_OrderMenu{
			MenuId:  menu.MenuId,
			Amount:  int32(menu.Amount),
			Price:   menu.Price,
			Request: menu.Request,
		}
		orderMenus = append(orderMenus, pbMenu)
	}

	res, err := (*t.client).UpdateOrder(ctx, &pb.UpdateOrderRequest{
		Id:         id,
		Status:     order.Status,
		VendorId:   order.VendorId,
		UserId:     order.UserId,
		OrderMenus: orderMenus,
	})
	if err != nil {
		return nil, err
	}
	updatedOrder := &model.Order{
		ID:         res.Order.Id,
		Status:     res.Order.Status,
		OrderMenus: make([]*model.OrderMenu, 0),
		VendorId:   res.Order.VendorId,
		Price:      res.Order.Price,
		UserId:     res.Order.UserId,
		CreateAt:   res.Order.CreatedAt.AsTime(),
		UpdatedAt:  res.Order.UpdatedAt.AsTime(),
	}
	for _, orderMenu := range res.Order.OrderMenus {
		orderMenu := &model.OrderMenu{
			MenuId:  orderMenu.MenuId,
			Amount:  int(orderMenu.Amount),
			Price:   orderMenu.Price,
			Request: orderMenu.Request,
		}
		order.OrderMenus = append(order.OrderMenus, orderMenu)
	}
	return updatedOrder, nil
}

func (t *OrderClient) DeleteOrderByID(ctx context.Context, id string) (*model.Order, error) {
	_, err := (*t.client).DeleteOrder(ctx, &pb.OrderRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *OrderClient) GetOrderByVendorID(ctx context.Context, vid string) ([]*model.Order, error) {
	stream, err := (*t.client).GetOrdersByVendorId(ctx, &pb.GetOrdersByVendorIdRequest{VendorId: vid})
	if err != nil {
		return nil, err
	}
	var orders []*model.Order

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListOrders: %v", err)
		}

		order := &model.Order{
			ID:         res.Id,
			Status:     res.Status,
			OrderMenus: make([]*model.OrderMenu, 0),
			VendorId:   res.VendorId,
			Price:      res.Price,
			UserId:     res.UserId,
			CreateAt:   res.CreatedAt.AsTime(),
			UpdatedAt:  res.UpdatedAt.AsTime(),
		}
		for _, orderMenu := range res.OrderMenus {
			orderMenu := &model.OrderMenu{
				MenuId:  orderMenu.MenuId,
				Amount:  int(orderMenu.Amount),
				Price:   orderMenu.Price,
				Request: orderMenu.Request,
			}
			order.OrderMenus = append(order.OrderMenus, orderMenu)
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func (t *OrderClient) GetOrderByUserID(ctx context.Context, uid string) ([]*model.Order, error) {
	stream, err := (*t.client).GetOrdersByUserId(ctx, &pb.GetOrdersByUserIdRequest{UserId: uid})
	if err != nil {
		return nil, err
	}
	var orders []*model.Order

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("ListOrders: %v", err)
		}

		order := &model.Order{
			ID:         res.Id,
			Status:     res.Status,
			OrderMenus: make([]*model.OrderMenu, 0),
			VendorId:   res.VendorId,
			Price:      res.Price,
			UserId:     res.UserId,
			CreateAt:   res.CreatedAt.AsTime(),
			UpdatedAt:  res.UpdatedAt.AsTime(),
		}
		for _, orderMenu := range res.OrderMenus {
			orderMenu := &model.OrderMenu{
				MenuId:  orderMenu.MenuId,
				Amount:  int(orderMenu.Amount),
				Price:   orderMenu.Price,
				Request: orderMenu.Request,
			}
			order.OrderMenus = append(order.OrderMenus, orderMenu)
		}

		orders = append(orders, order)
	}
	return orders, nil
}

func ProvideOrderClient(client *pb.OrderServiceClient) *OrderClient {
	return &OrderClient{client: client}
}
