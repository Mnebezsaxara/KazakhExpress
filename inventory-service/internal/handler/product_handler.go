package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if product.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Количество не может быть отрицательным"})
		return
	}

	createdProduct, err := h.usecase.Create(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении товара"})
		return
	}

	log.Printf("Создан новый продукт: %+v", createdProduct)
	c.JSON(http.StatusCreated, createdProduct)
}

func (h *ProductHandler) List(c *gin.Context) {
	filter := domain.ProductFilter{
		Category: c.DefaultQuery("category", ""),
		Page:     1,
		Limit:    10,
	}

	if minPrice := c.DefaultQuery("minPrice", ""); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filter.MinPrice = price
		}
	}

	if maxPrice := c.DefaultQuery("maxPrice", ""); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filter.MaxPrice = price
		}
	}

	if page := c.DefaultQuery("page", "1"); page != "" {
		if p, err := strconv.Atoi(page); err == nil {
			filter.Page = p
		}
	}

	if limit := c.DefaultQuery("limit", "10"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = l
		}
	}

	products, total, err := h.usecase.List(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товаров"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"total":    total,
	})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	product, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	product.ID = id
	updatedProduct, err := h.usecase.Update(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Удалено успешно"})
}
