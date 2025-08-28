package routes

import (
	"e-learning-system/internal/api/controller"
	"e-learning-system/internal/api/middleware"
	"e-learning-system/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

// RegisterOrganizationBillingRoutes registers organization billing-related endpoints
func RegisterOrganizationBillingRoutes(
	routes *gin.Engine,
	billingController *controller.OrganizationBillingController,
	tokenRepo repository.TokenRepository,
) {
	authMiddleware := middleware.AuthMiddleware(tokenRepo)

	billingGroup := routes.Group("/organization-billings")
	{
		// All organization billing routes are protected by authentication
		billingGroup.Use(authMiddleware)
		{
			billingGroup.POST("", billingController.CreateOrganizationBilling)       // Create new billing record
			billingGroup.PUT("/:id", billingController.UpdateOrganizationBilling)    // Update billing by ID
			billingGroup.DELETE("/:id", billingController.DeleteOrganizationBilling) // Soft delete billing by ID
			billingGroup.GET("/:id", billingController.GetOrganizationBillingByID)   // Get billing by ID
			billingGroup.GET("", billingController.GetAllOrganizationBillings)       // Get all billing records
		}
	}
}
