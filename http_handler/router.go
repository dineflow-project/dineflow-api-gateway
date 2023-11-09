package http_handler

import (
	"dineflow-api-gateway/middleware"
	"dineflow-api-gateway/services"

	"github.com/gin-gonic/gin"
)

func ProvideRouter(
	r *gin.Engine,
	userService services.UserService,
	orderHandler IOrderHandler,
	reviewHandler IReviewHandler,
	notificationHandler INotificationHandler,
	menuHandler IMenuHandler,
) {
	// order service
	r.GET("/order", orderHandler.GetAllOrders)
	r.GET("/order/:id", orderHandler.GetOrderByID)
	r.POST("/order", middleware.DeserializeUser(userService), middleware.Authorize("user"), orderHandler.CreateOrder)
	r.PUT("/order/:id", orderHandler.UpdateOrderByID)
	r.DELETE("/order/:id", orderHandler.DeleteOrderByID)
	r.GET("/order/byVendor/:vendorId", middleware.DeserializeUser(userService), middleware.Authorize("vendor"), orderHandler.GetOrderByVendorID)
	r.GET("/order/byUser/:userId", middleware.DeserializeUser(userService), middleware.Authorize("user"), orderHandler.GetOrderByUserID)

	// review service
	r.GET("/review/:id", reviewHandler.GetReviewByID)
	r.GET("/review/avgScore/:vendorId", reviewHandler.GetAvgReviewScoreByVendorID)
	r.GET("/review", reviewHandler.GetAllReviews)
	r.GET("/review/byVendor/:vendorId", reviewHandler.GetReviewByVendorID)
	r.POST("/review", middleware.DeserializeUser(userService), middleware.Authorize("user"), reviewHandler.CreateReview)
	r.PUT("/review/:id", reviewHandler.UpdateReviewByID)
	r.DELETE("/review/:id", reviewHandler.DeleteReviewByID)

	// notification service
	r.GET("/notification/:recipientID", middleware.DeserializeUser(userService), notificationHandler.GetAllNotifiactions)
	r.GET("/notification/unread/:recipientID", middleware.DeserializeUser(userService), notificationHandler.GetUnreadNotification)

	//menu service
	r.GET("/menu/:id", menuHandler.GetMenuByID)
	r.GET("/menu", menuHandler.GetAllMenus)
	r.GET("/menu/byVendor/:vendorId", menuHandler.GetMenuByVendorID)
	r.POST("/menu", middleware.DeserializeUser(userService), middleware.Authorize("vendor"), menuHandler.CreateMenu)
	r.PUT("/menu/:id", middleware.DeserializeUser(userService), middleware.Authorize("vendor"), menuHandler.UpdateMenuByID)
	r.DELETE("/menu/:id", middleware.DeserializeUser(userService), middleware.Authorize("vendor"), menuHandler.DeleteMenuByID)

	// vendor
	r.GET("/vendor/:id", menuHandler.GetVendorByID)
	r.GET("/vendor/byOwner/:id", menuHandler.GetVendorByOwnerID)
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
