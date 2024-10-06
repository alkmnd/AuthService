package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/refresh", h.refresh)
	}

	return router
}

func (h *Handler) signUp(c *gin.Context) {

}

func (h *Handler) refresh(c *gin.Context) {

}
