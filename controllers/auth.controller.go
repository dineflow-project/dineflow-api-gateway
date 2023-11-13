package controllers

import (
	"dineflow-api-gateway/configs"
	"dineflow-api-gateway/model"
	"dineflow-api-gateway/services"
	"dineflow-api-gateway/utils"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
	ctx         context.Context
	collection  *mongo.Collection
}

func NewAuthController(authService services.AuthService, userService services.UserService, ctx context.Context, collection *mongo.Collection) AuthController {
	return AuthController{authService, userService, ctx, collection}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var user *model.SignUpInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": err.Error()})
		return
	}

	if user.Password != user.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "Passwords do not match"})
		return
	}

	newUser, err := ac.authService.SignUpUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"code": "409", "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"code": "502", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"code": "200", "data": model.FilteredResponse(newUser)})
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var credentials *model.SignInInput

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": err.Error()})
		return
	}

	user, err := ac.userService.FindUserByEmail(credentials.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "Invalid email or password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": err.Error()})
		return
	}

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "Invalid email or Password"})
		return
	}

	// configs, _ := configs.LoadConfigs(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(configs.EnvAccessTokenExpire(), user.ID, configs.EnvAccessTokenPK())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": err.Error()})
		return
	}

	// refresh_token, err := utils.CreateToken(configs.RefreshTokenExpiresIn, user.ID, configs.RefreshTokenPrivateKey)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"code": "500", "error": err.Error()})
	// 	return
	// }

	ctx.SetCookie("access_token", access_token, configs.EnvAccessTokenMaxAge()*60, "/", "localhost", false, true)
	//ctx.SetCookie("refresh_token", refresh_token, configs.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", configs.EnvAccessTokenMaxAge()*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "token": access_token, "data": model.FilteredResponse(user)})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	//ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "message": "Logged out"})
}

func (ac *AuthController) ResetPassword(ctx *gin.Context) {
	var userCredential *model.ResetPasswordInput

	if err := ctx.ShouldBindJSON(&userCredential); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": err.Error()})
		return
	}

	if userCredential.Password != userCredential.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "Passwords do not match"})
		return
	}

	// Hash the new password.
	hashedPassword, err := utils.HashPassword(userCredential.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "500", "error": "Password hashing failed"})
		return
	}

	// Retrieve the current user using the GetMe function.
	currentUser := ctx.MustGet("currentUser").(*model.DBResponse)
	if currentUser == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "error": "User not authenticated"})
		return
	}

	// Update the user's password in the database.
	// You should update this logic to match your database query to update the user's password.
	query := bson.D{{Key: "_id", Value: currentUser.ID}} // Replace with the appropriate query.

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "password", Value: hashedPassword},
			{Key: "updated_at", Value: time.Now()},
		}},
	}
	result, err := ac.collection.UpdateOne(ac.ctx, query, update)

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "400", "error": "User not found or update failed"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "500", "error": "Password update failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": "200", "message": "Password updated successfully"})
}
