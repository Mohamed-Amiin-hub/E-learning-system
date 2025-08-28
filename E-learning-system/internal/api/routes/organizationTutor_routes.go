package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationTutorRoutes registers organization tutor-related endpoints
func RegisterOrganizationTutorRoutes(
	routes *gin.Engine,
	orgTutorController *controller.OrganizationTutorController,
	tokenRepo repository.TokenRepository,
) {
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	tutorGroup := routes.Group("/organization-tutors")
	{
		// All organization tutor routes are protected by authentication
		tutorGroup.Use(authMiddleware)
		{
			tutorGroup.POST("", orgTutorController.CreateTutorOrganization)       // Add a tutor to organization
			tutorGroup.PUT("/:id", orgTutorController.UpdateTutorOrganization)    // Update tutor info (approve, etc.)
			tutorGroup.DELETE("/:id", orgTutorController.DeleteTutorOrganization) // Remove tutor from organization
			tutorGroup.GET("/:id", orgTutorController.GetTutorOrganizationByID)   // Get tutor by ID
			tutorGroup.GET("", orgTutorController.GetAllTutorsOrganization)       // Get all tutors
		}
	}
}
