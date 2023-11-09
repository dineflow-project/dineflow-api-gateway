package routes

import (
	"dineflow-api-gateway/controllers"
	"dineflow-api-gateway/middleware"
	"dineflow-api-gateway/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(userService), uc.userController.GetMe)
}
