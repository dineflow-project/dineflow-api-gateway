package controllers

import (
	"api-gateway/model"
	"api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*model.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "data": model.FilteredResponse(currentUser)})
}
