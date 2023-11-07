package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	orderHandler IOrderHandler,
	reviewHandler IReviewHandler,
	notificationHandler INotificationHandler,

) {
	// order service
	r.GET("/order", orderHandler.GetAllOrders)
	r.GET("/order/:id", orderHandler.GetOrderByID)
	r.POST("/order", orderHandler.CreateOrder)
	r.PUT("/order/:id", orderHandler.UpdateOrderByID)
	r.DELETE("/order/:id", orderHandler.DeleteOrderByID)
	r.GET("/order/byVendor/:vendorId", orderHandler.GetOrderByVendorID)
	r.GET("/order/byUser/:userId", orderHandler.GetOrderByUserID)

	// professor service
	r.GET("/review/:id", reviewHandler.GetReviewByID)
	r.GET("/review", reviewHandler.GetAllReviews)
	r.GET("/review/byVendor/:vendorId", reviewHandler.GetReviewByVendorID)
	r.POST("/review", reviewHandler.CreateReview)
	r.PUT("/review/:id", reviewHandler.UpdateReviewByID)
	r.DELETE("/review/:id", reviewHandler.DeleteReviewByID)

	r.GET("/notification/:recipientID", notificationHandler.GetAllNotifiactions)
	r.GET("/notification/unread/:recipientID", notificationHandler.GetUnreadNotification)

}
