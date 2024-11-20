package handler

import (
	"github.com/gin-gonic/gin"
	"testCase/pkg/service"
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
		api.GET("/users/:id/status", h.getUserStatus)

		api.GET("/users/leaderboard", h.getLeaderboard)

		api.POST("/users/:id/task/complete", h.completeTask)

		api.POST("/users/:id/referrer", h.addReferrer)

	}

	return router
}
