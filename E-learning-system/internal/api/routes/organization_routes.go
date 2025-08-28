package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationRoutes registers organization-related endpoints
func RegisterOrganizationRoutes(routes *gin.Engine, orgController *controller.OrganizationController, tokenRepo repository.TokenRepository,) {
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	orgGroup := routes.Group("/organizations")
	{
		// All organization routes are protected by authentication
		orgGroup.Use(authMiddleware)
		{
			orgGroup.POST("", orgController.CreateOrganization)       // Create new organization
			orgGroup.PUT("/:id", orgController.UpdateOrganization)    // Update organization by ID
			orgGroup.DELETE("/:id", orgController.DeleteOrganization) // Soft delete organization by ID
			orgGroup.GET("/:id", orgController.GetOrganizationByID)   // Get organization by ID
			orgGroup.GET("", orgController.GetAllOrganizations)       // Get all organizations
		}
	}
}
