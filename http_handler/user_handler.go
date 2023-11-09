package http_handler

import (
	"net/http"

	restClient "dineflow-api-gateway/client_rest"

	"dineflow-api-gateway/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userClientRest restClient.UserClientRest
}

type IUserHandler interface {
	GetMe(c *gin.Context)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	user, err := h.userClientRest.GetMe()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var params model.User
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	user, err := h.userClientRest.SignUp(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var params model.User
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	resp, err := h.userClientRest.Login(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": resp,
	})
}

func ProvideUserHandler(userClientRest restClient.UserClientRest) *UserHandler {
	return &UserHandler{
		userClientRest: userClientRest,
	}
}
