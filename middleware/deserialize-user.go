package middleware

import (
	"dineflow-api-gateway/configs"
	"dineflow-api-gateway/model"
	"dineflow-api-gateway/services"
	"dineflow-api-gateway/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "error": "You are not logged in"})
			return
		}

		// config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, configs.EnvAccessTokenPublicK())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "error": err.Error()})
			return
		}

		user, err := userService.FindUserById(fmt.Sprint(sub))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "error": "The user belonging to this token no logger exists"})
			return
		}
		ctx.Set("currentUser", user)
		ctx.Next()
	}
}

func Authorize(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := ctx.Get("currentUser")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "401", "error": "User not authenticated"})
			return
		}

		// Assuming the user has a 'role' field in the user object
		userRole := user.(*model.DBResponse).Role
		if userRole == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "500", "error": "Failed to get user role"})
			return
		}

		// Check if the user's role is in the allowed roles
		if !contains(roles, userRole) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": "403", "error": "Insufficient permissions"})
			return
		}

		ctx.Next()
	}
}

// Helper function to check if a string is in a slice of strings
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
