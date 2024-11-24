package http

import (
	"github.com/karthiknarayan07/IAM-System/http/handlers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

// RegisterHandlers allows registering handlers for different routes.
func (r *Router) RegisterHandlers(userHandler *handlers.UserHandler) {
	// User routes
	userRoutes := r.Engine.Group("/users")
	{
		userRoutes.POST("", userHandler.RegisterUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
	}

	// Future handlers for other domains can be added here
	// For example:
	// roleRoutes := r.Engine.Group("/roles")
	// roleRoutes.POST("", roleHandler.CreateRole)
}
