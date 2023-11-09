package http_handler

import (
	"net/http"

	restClient "dineflow-api-gateway/client_rest"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	notificationClientRest restClient.NotificationClientRest
}

type INotificationHandler interface {
	GetUnreadNotification(c *gin.Context)
	GetAllNotifiactions(c *gin.Context)
}

func (h *NotificationHandler) GetUnreadNotification(c *gin.Context) {
	recipientID := c.Param("recipientID")
	notifications, err := h.notificationClientRest.GetUnreadNotification(recipientID) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": notifications})
}

func (h *NotificationHandler) GetAllNotifiactions(c *gin.Context) {
	recipientID := c.Param("recipientID")
	notifications, err := h.notificationClientRest.GetAllNotifiactions(recipientID) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": notifications})
}

func ProvideNotificationHandler(notificationClientRest restClient.NotificationClientRest) *NotificationHandler {
	return &NotificationHandler{
		notificationClientRest: notificationClientRest,
	}
}
