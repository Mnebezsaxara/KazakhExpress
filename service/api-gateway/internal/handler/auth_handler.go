package handler

import (
	"net/http"

	"api-gateway/internal/usecase"
	"api-gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC *usecase.AuthUsecase
}

func NewAuthHandler(authUC *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
	}
}

func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", h.SignUp)
		auth.POST("/login", h.Login)
		auth.GET("/profile", h.GetUser)
		auth.PUT("/profile", h.UpdateUser)
		auth.DELETE("/profile", h.DeleteUser)
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.authUC.SignUp(c.Request.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, "User registered successfully")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.authUC.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response.Success(c, "Login successful", gin.H{"token": token})
}

func (h *AuthHandler) GetUser(c *gin.Context) {
	email := c.GetString("user_email")
	if email == "" {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.authUC.GetProfile(c.Request.Context(), email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Don't return the password
	user.Password = ""
	response.Success(c, "User profile retrieved", user)
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {
	email := c.GetString("user_email")
	if email == "" {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.authUC.UpdateUser(c.Request.Context(), email, req.Name, req.Password, req.Role)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "User updated successfully", nil)
}

func (h *AuthHandler) DeleteUser(c *gin.Context) {
	email := c.GetString("user_email")
	if email == "" {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := h.authUC.DeleteUser(c.Request.Context(), email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "User deleted successfully", nil)
}
