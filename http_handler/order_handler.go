package http_handler

import (
	grpcClient "dineflow-api-gateway/client_grpc"
	"dineflow-api-gateway/model"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderClientgRPC grpcClient.OrdergRPCClient
}

type IOrderHandler interface {
	GetAllOrders(c *gin.Context)
	GetOrderByID(c *gin.Context)
	CreateOrder(c *gin.Context)
	UpdateOrderByID(c *gin.Context)
	DeleteOrderByID(c *gin.Context)
	GetOrderByVendorID(c *gin.Context)
	GetOrderByUserID(c *gin.Context)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderClientgRPC.GetAllOrders(c)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": orders,
	})
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderClientgRPC.GetOrderByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": order,
	})
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}

	//attach user id that logged in
	order.UserId = c.MustGet("currentUser").(*model.DBResponse).ID.Hex()

	createdOrder, err := h.orderClientgRPC.CreateOrder(c, &order)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": createdOrder,
	})
}

func (h *OrderHandler) UpdateOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	updatedOrder, err := h.orderClientgRPC.UpdateOrderByID(c, id, &order)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": updatedOrder,
	})
}

func (h *OrderHandler) DeleteOrderByID(c *gin.Context) {
	id := c.Param("id")
	deletedOrder, err := h.orderClientgRPC.DeleteOrderByID(c, id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": deletedOrder,
	})
}

func (h *OrderHandler) GetOrderByVendorID(c *gin.Context) {
	vid := c.Param("vendorId")
	order, err := h.orderClientgRPC.GetOrderByVendorID(c, vid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": order,
	})
}

func (h *OrderHandler) GetOrderByUserID(c *gin.Context) {
	uid := c.Param("userId")
	order, err := h.orderClientgRPC.GetOrderByUserID(c, uid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": order,
	})
}

func ProvideOrderHandler(orderClientgRPC grpcClient.OrdergRPCClient) *OrderHandler {
	return &OrderHandler{
		orderClientgRPC: orderClientgRPC,
	}
}
