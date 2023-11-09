package http_handler

import (
	"net/http"

	restClient "dineflow-api-gateway/client_rest"
	"dineflow-api-gateway/model"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	menuClientRest restClient.MenuClientRest
}

type IMenuHandler interface {
	//menu
	GetMenuByID(c *gin.Context)
	GetAllMenus(c *gin.Context)
	GetMenuByVendorID(c *gin.Context)
	CreateMenu(c *gin.Context)
	UpdateMenuByID(c *gin.Context)
	DeleteMenuByID(c *gin.Context)

	//vendor
	GetVendorByID(c *gin.Context)
	GetVendorByOwnerID(c *gin.Context)
	GetAllVendors(c *gin.Context)
	GetAllVendorsByCanteenID(c *gin.Context)
	CreateVendor(c *gin.Context)
	UpdateVendorByID(c *gin.Context)
	DeleteVendorByID(c *gin.Context)

	//canteen
	GetCanteenByID(c *gin.Context)
	GetAllCanteens(c *gin.Context)
	CreateCanteen(c *gin.Context)
	UpdateCanteenByID(c *gin.Context)
	DeleteCanteenByID(c *gin.Context)
}

// menu ---------------------------------------------------------------------------------------------------------------------------------------
func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	id := c.Param("id")
	menu, err := h.menuClientRest.GetMenuByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": menu})
}

func (h *MenuHandler) GetMenuByVendorID(c *gin.Context) {
	vendor_id := c.Param("vendorId")
	menu, err := h.menuClientRest.GetMenuByVendorID(vendor_id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": menu})
}

func (h *MenuHandler) GetAllMenus(c *gin.Context) {
	canteenId := c.Query("canteenId")
	vendorId := c.Query("vendorId")
	minprice := c.Query("minprice")
	maxprice := c.Query("maxprice")
	menus, err := h.menuClientRest.GetAllMenus(canteenId, vendorId, minprice, maxprice) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": menus})
}

func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var params model.Menu
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.CreateMenu(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Menu created successfully",
	})
}

func (h *MenuHandler) UpdateMenuByID(c *gin.Context) {
	id := c.Param("id")
	var params model.Menu
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.UpdateMenuByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Menu updated successfully",
	})
}

func (h *MenuHandler) DeleteMenuByID(c *gin.Context) {
	id := c.Param("id")
	err := h.menuClientRest.DeleteMenuByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Menu deleted successfully",
	})
}

// vendor ---------------------------------------------------------------------------------------------------------------------------------------
func (h *MenuHandler) GetVendorByID(c *gin.Context) {
	id := c.Param("id")
	vendor, err := h.menuClientRest.GetVendorByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendor})
}

func (h *MenuHandler) GetVendorByOwnerID(c *gin.Context) {
	id := c.Param("id")
	vendor, err := h.menuClientRest.GetVendorByOwnerID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendor})
}

func (h *MenuHandler) GetAllVendors(c *gin.Context) {
	vendors, err := h.menuClientRest.GetAllVendors() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendors})
}

func (h *MenuHandler) GetAllVendorsByCanteenID(c *gin.Context) {
	id := c.Param("id")
	vendors, err := h.menuClientRest.GetAllVendorsByCanteenID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendors})
}

func (h *MenuHandler) CreateVendor(c *gin.Context) {
	var params model.Vendor
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.CreateVendor(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Vendor created successfully",
	})
}

func (h *MenuHandler) UpdateVendorByID(c *gin.Context) {
	id := c.Param("id")
	var params model.Vendor
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.UpdateVendorByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Vendor updated successfully",
	})
}

func (h *MenuHandler) DeleteVendorByID(c *gin.Context) {
	id := c.Param("id")
	err := h.menuClientRest.DeleteVendorByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Vendor deleted successfully",
	})
}

// canteen ---------------------------------------------------------------------------------------------------------------------------------------
func (h *MenuHandler) GetCanteenByID(c *gin.Context) {
	id := c.Param("id")
	vendor, err := h.menuClientRest.GetCanteenByID(id) // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendor})
}

func (h *MenuHandler) GetAllCanteens(c *gin.Context) {
	vendors, err := h.menuClientRest.GetAllCanteens() // Method on your rest client
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "400", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": vendors})
}

func (h *MenuHandler) CreateCanteen(c *gin.Context) {
	var params model.Canteen
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.CreateCanteen(params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Canteen created successfully",
	})
}

func (h *MenuHandler) UpdateCanteenByID(c *gin.Context) {
	id := c.Param("id")
	var params model.Canteen
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	err := h.menuClientRest.UpdateCanteenByID(id, params)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Canteen updated successfully",
	})
}

func (h *MenuHandler) DeleteCanteenByID(c *gin.Context) {
	id := c.Param("id")
	err := h.menuClientRest.DeleteCanteenByID(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  "400",
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "Canteen deleted successfully",
	})
}

// ------------------------------------------------------------------------------------
func ProvideMenuHandler(menuClientRest restClient.MenuClientRest) *MenuHandler {
	return &MenuHandler{
		menuClientRest: menuClientRest,
	}
}
