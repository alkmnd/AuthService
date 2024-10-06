package handler

import (
	"authService/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	services *service.AuthService
}

func NewHandler(services *service.AuthService) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/tokens", h.getTokens)
		auth.GET("/refresh", h.refresh)
	}

	return router
}

type userSignUpInput struct {
	Email string `json:"email"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input userSignUpInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "incorrect format")
		return
	}

	id, err := h.services.CreateUser(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getTokens(c *gin.Context) {

}

func (h *Handler) refresh(c *gin.Context) {

}
