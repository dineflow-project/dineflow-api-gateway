package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	orderHandler IOrderHandler,

) {
	// order service
	r.GET("/order", orderHandler.GetAllOrders)
	r.GET("/order/:id", orderHandler.GetOrderByID)
	r.POST("/order", orderHandler.CreateOrder)
	r.PUT("/order/:id", orderHandler.UpdateOrderByID)
	r.DELETE("/order/:id", orderHandler.DeleteOrderByID)
	r.GET("/order/byVendor/:vendorId", orderHandler.GetOrderByVendorID)
	r.GET("/order/byUser/:userId", orderHandler.GetOrderByUserID)
}
