package routes

import (
	"api-gateway/controllers"
	"api-gateway/middleware"
	"api-gateway/services"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	//router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), rc.authController.LogoutUser)
	router.PATCH("/resetpassword", middleware.DeserializeUser(userService), rc.authController.ResetPassword)

}
