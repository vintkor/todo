package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vintkor/todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.authSignUp)
		auth.POST("/sign-in", h.authSignIn)
	}

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", h.usersList)
			users.POST("/", h.usersCreate)
			users.GET("/:id", h.usersRetrieve)
		}
	}

	return router
}
