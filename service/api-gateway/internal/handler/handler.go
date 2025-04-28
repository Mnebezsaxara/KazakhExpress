package handler

import (
	"net/http"
	"strconv"

	"api-gateway/internal/client"
	"api-gateway/proto"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	inventoryClient *client.InventoryClient
}

func NewHandler(inventoryClient *client.InventoryClient) *Handler {
	return &Handler{
		inventoryClient: inventoryClient,
	}
}

func (h *Handler) HandleProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")
	category := c.Query("category")

	// Convert price strings to float64
	var minPrice, maxPrice float64
	var err error
	if minPriceStr != "" {
		minPrice, err = strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minPrice format"})
			return
		}
	}
	if maxPriceStr != "" {
		maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maxPrice format"})
			return
		}
	}

	products, err := h.inventoryClient.GetProducts(c.Request.Context(), int32(page), int32(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Фильтрация по цене и категории
	filteredProducts := make([]*proto.Product, 0)
	for _, p := range products {
		if minPriceStr != "" && p.Price < minPrice {
			continue
		}
		if maxPriceStr != "" && p.Price > maxPrice {
			continue
		}
		if category != "" && p.Category != category {
			continue
		}
		filteredProducts = append(filteredProducts, p)
	}

	c.JSON(http.StatusOK, gin.H{
		"products": filteredProducts,
	})
}

func (h *Handler) HandleCategories(c *gin.Context) {
	categories, err := h.inventoryClient.GetCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *Handler) HandleOrders(c *gin.Context) {
	// Здесь будет логика создания заказа
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created",
	})
}

func (h *Handler) HandleGetOrders(c *gin.Context) {
	// Здесь будет логика получения заказов
	c.JSON(http.StatusOK, gin.H{
		"orders": []gin.H{},
	})
}

func (h *Handler) HandleCancelOrder(c *gin.Context) {
	// Здесь будет логика отмены заказа
	c.JSON(http.StatusOK, gin.H{
		"message": "Order cancelled",
	})
} 