package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationAdminRoutes registers organization admin-related endpoints
func RegisterOrganizationAdminRoutes(
	routes *gin.Engine,
	orgAdminController *controller.OrganizationAdminController,
	tokenRepo repository.TokenRepository,
) {
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	adminGroup := routes.Group("/organization-admins")
	{
		// All organization admin routes are protected by authentication
		adminGroup.Use(authMiddleware)
		{
			adminGroup.POST("", orgAdminController.CreateOrganizationAdmin)       // Create new admin
			adminGroup.PUT("/:id", orgAdminController.UpdateOrganizationAdmin)    // Update admin by ID
			adminGroup.DELETE("/:id", orgAdminController.DeleteOrganizationAdmin) // Soft delete admin by ID
			adminGroup.GET("/:id", orgAdminController.GetOrganizationAdminByID)   // Get admin by ID
			adminGroup.GET("", orgAdminController.GetAllOrganizationAdmins)       // Get all admins
		}
	}
}
