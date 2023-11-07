package http_handler

import (
	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	orderHandler IOrderHandler,
	reviewHandler IReviewHandler,
	userHandler IUserHandler,
	notificationHandler INotificationHandler,
	menuHandler IMenuHandler,
) {
	// order service
	r.GET("/order", orderHandler.GetAllOrders)
	r.GET("/order/:id", orderHandler.GetOrderByID)
	r.POST("/order", orderHandler.CreateOrder)
	r.PUT("/order/:id", orderHandler.UpdateOrderByID)
	r.DELETE("/order/:id", orderHandler.DeleteOrderByID)
	r.GET("/order/byVendor/:vendorId", orderHandler.GetOrderByVendorID)
	r.GET("/order/byUser/:userId", orderHandler.GetOrderByUserID)

	// review service
	r.GET("/review/:id", reviewHandler.GetReviewByID)
	r.GET("/review", reviewHandler.GetAllReviews)
	r.GET("/review/byVendor/:vendorId", reviewHandler.GetReviewByVendorID)
	r.POST("/review", reviewHandler.CreateReview)
	r.PUT("/review/:id", reviewHandler.UpdateReviewByID)
	r.DELETE("/review/:id", reviewHandler.DeleteReviewByID)

	// user service
	r.GET("/users/me", userHandler.GetMe)
	r.POST("/auth/login", userHandler.Login)
	r.POST("/auth/register", userHandler.SignUp)

	// notification service
	r.GET("/notification/:recipientID", notificationHandler.GetAllNotifiactions)
	r.GET("/notification/unread/:recipientID", notificationHandler.GetUnreadNotification)

	//menu service
	r.GET("/menu/:id", menuHandler.GetMenuByID)
	r.GET("/menu", menuHandler.GetAllMenus)
	r.GET("/menu/byVendor/:vendorId", menuHandler.GetMenuByVendorID)
	r.POST("/menu", menuHandler.CreateMenu)
	r.PUT("/menu/:id", menuHandler.UpdateMenuByID)
	r.DELETE("/menu/:id", menuHandler.DeleteMenuByID)

	// vendor
	r.GET("/vendor/:id", menuHandler.GetVendorByID)
	r.GET("/vendor", menuHandler.GetAllVendors)
	r.GET("/vendor/canteen/:id", menuHandler.GetAllVendorsByCanteenID)
	r.POST("/vendor", menuHandler.CreateVendor)
	r.PUT("/vendor/:id", menuHandler.UpdateVendorByID)
	r.DELETE("/vendor/:id", menuHandler.DeleteVendorByID)

	// canteen
	r.GET("/canteen/:id", menuHandler.GetCanteenByID)
	r.GET("/canteen", menuHandler.GetAllCanteens)
	r.POST("/canteen", menuHandler.CreateCanteen)
	r.PUT("/canteen/:id", menuHandler.UpdateCanteenByID)
	r.DELETE("/canteen/:id", menuHandler.DeleteCanteenByID)

}
