package handler

import (
	"TODO/service"

	"github.com/gin-gonic/gin"
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			str := lists.Group(":id/str")
			{
				str.POST("/", h.createStr)
				str.GET("/", h.getAllStr)
			}
		}

		str := api.Group("str")
		{
			str.GET("/:id", h.getStrById)
			str.PUT("/:id", h.updateStr)
			str.DELETE("/:id", h.deleteStr)
		}
	}

	return router
}
