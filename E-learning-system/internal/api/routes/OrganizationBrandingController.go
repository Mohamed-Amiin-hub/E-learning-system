package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationBrandingRoutes registers organization branding-related endpoints
func RegisterOrganizationBrandingRoutes(
	routes *gin.Engine,
	brandingController *controller.OrganizationBrandingController,
	tokenRepo repository.TokenRepository,
) {
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	brandingGroup := routes.Group("/organization-brandings")
	{
		// All organization branding routes are protected by authentication
		brandingGroup.Use(authMiddleware)
		{
			brandingGroup.POST("", brandingController.CreateOrganizationBranding)       // Create new branding
			brandingGroup.PUT("/:id", brandingController.UpdateOrganizationBranding)    // Update branding by ID
			brandingGroup.DELETE("/:id", brandingController.DeleteOrganizationBranding) // Soft delete branding by ID
			brandingGroup.GET("/:id", brandingController.GetOrganizationBrandingByID)   // Get branding by ID
			brandingGroup.GET("", brandingController.GetAllOrganizationBrandings)       // Get all brandings
		}
	}
}
