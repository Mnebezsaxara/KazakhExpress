package handler

import (
	"net/http"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(u *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: u}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	createdCategory, err := h.usecase.Create(c.Request.Context(), &category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении категории"})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, total, err := h.usecase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении категорий"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"total":      total,
	})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	category, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.usecase.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ошибка при удалении категории"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}
