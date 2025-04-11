package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/domain"
	"github.com/Mnebezsaxara/KazakhExpress/inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

    // Проверка на количество на складе (если оно меньше 0)
    if product.Stock < 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Количество не может быть отрицательным"})
        return
    }

    // Создаем продукт
    err := h.usecase.CreateProduct(&product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении товара"})
        return
    }

    // Логируем созданный продукт
    log.Printf("Создан новый продукт: %+v", product)

    // Возвращаем созданный продукт с ID
    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) List(c *gin.Context) {
    filters := make(map[string]interface{})

    // Логируем все входящие параметры
    log.Printf("List request received with query params: %+v", c.Request.URL.Query())

    // Фильтрация по категории, если категория задана
    category := c.DefaultQuery("category", "")
    if category != "" {
        filters["category"] = category
        log.Printf("Category filter applied: %s", category)
    }

    // Фильтрация по минимальной цене
    minPrice := c.DefaultQuery("minPrice", "")
    if minPrice != "" {
        filters["minPrice"] = minPrice
        log.Printf("MinPrice filter applied: %s", minPrice)
    }

    // Фильтрация по максимальной цене
    maxPrice := c.DefaultQuery("maxPrice", "")
    if maxPrice != "" {
        filters["maxPrice"] = maxPrice
        log.Printf("MaxPrice filter applied: %s", maxPrice)
    }

    // Пагинация
    page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")

    pageInt, _ := strconv.Atoi(page)
    limitInt, _ := strconv.Atoi(limit)

    log.Printf("Pagination params: page=%d, limit=%d", pageInt, limitInt)

    // Получаем товары с использованием фильтров и пагинации
    products, err := h.usecase.GetProducts(filters, limitInt, (pageInt-1)*limitInt)
    if err != nil {
        log.Printf("Error getting products: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении товаров"})
        return
    }

    // Логируем полученные данные
    log.Printf("Retrieved %d products", len(products))
    for i, product := range products {
        log.Printf("Product %d: ID=%s, Name=%s, Category=%s, Price=%f", 
            i+1, product.ID.Hex(), product.Name, product.Category, product.Price)
    }

    // Если данных нет, возвращаем пустой массив
    if len(products) == 0 {
        log.Println("No products found")
        c.JSON(http.StatusOK, []domain.Product{})
        return
    }

    c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
    id := c.Param("id")
    product, err := h.usecase.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
    id := c.Param("id")
    var update bson.M
    if err := c.ShouldBindJSON(&update); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
        return
    }

    err := h.usecase.UpdateProduct(id, update)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Обновлено успешно"})
}

func (h *ProductHandler) Delete(c *gin.Context) {
    id := c.Param("id")
    err := h.usecase.DeleteProduct(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Удалено успешно"})
}
