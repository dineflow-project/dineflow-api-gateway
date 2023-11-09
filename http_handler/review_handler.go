package http_handler

import (
	"net/http"

	restClient "dineflow-api-gateway/client_rest"
	"dineflow-api-gateway/model"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewClientRest restClient.ReviewClientRest
}

type IReviewHandler interface {
	GetReviewByID(c *gin.Context)
	GetAllReviews(c *gin.Context)
	GetReviewByVendorID(c *gin.Context)
	CreateReview(c *gin.Context)
	UpdateReviewByID(c *gin.Context)
	DeleteReviewByID(c *gin.Context)
}

func (h *ReviewHandler) GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	review, err := h.reviewClientRest.GetReviewByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": review})
}

func (h *ReviewHandler) GetReviewByVendorID(c *gin.Context) {
	vendor_id := c.Param("vendorId")
	review, err := h.reviewClientRest.GetReviewByVendorID(vendor_id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": review})
}

func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.reviewClientRest.GetAllReviews() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": reviews})
}

func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var params model.Review
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}

	//attach user id that logged in
	params.User_id = c.MustGet("currentUser").(*model.DBResponse).ID.Hex()

	err := h.reviewClientRest.CreateReview(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		// "data": review,
		"message": "Review created successfully",
	})
}

func (h *ReviewHandler) UpdateReviewByID(c *gin.Context) {
	id := c.Param("id")
	var params model.Review
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.reviewClientRest.UpdateReviewByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		// "data": review,
		"message": "Review updated successfully",
	})
}

func (h *ReviewHandler) DeleteReviewByID(c *gin.Context) {
	id := c.Param("id")
	err := h.reviewClientRest.DeleteReviewByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		// "data": review,
		"message": "Review deleted successfully",
	})
}

func ProvideReviewHandler(reviewClientRest restClient.ReviewClientRest) *ReviewHandler {
	return &ReviewHandler{
		reviewClientRest: reviewClientRest,
	}
}
