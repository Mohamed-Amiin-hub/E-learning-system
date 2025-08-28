package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OrganizationBillingController defines the controller with its service
type OrganizationBillingController struct {
	OrganizationBillingService service.OrganizationBillingService
}

// NewOrganizationBillingController creates a new OrganizationBillingController instance
func NewOrganizationBillingController(billingService service.OrganizationBillingService) *OrganizationBillingController {
	return &OrganizationBillingController{OrganizationBillingService: billingService}
}

// CreateBilling handles the creation of a new organization billing
func (c *OrganizationBillingController) CreateOrganizationBilling(ctx *gin.Context) {
	var billing model.OrganizationBilling

	if err := ctx.ShouldBindJSON(&billing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBilling, err := c.OrganizationBillingService.CreateBilling(&billing)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdBilling)
}

// UpdateBilling handles the update of an existing billing
func (c *OrganizationBillingController) UpdateOrganizationBilling(ctx *gin.Context) {
	var billing model.OrganizationBilling

	billingIDParam := ctx.Param("id")
	billingID, err := uuid.FromString(billingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid billing ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&billing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	billing.ID = billingID

	if err := c.OrganizationBillingService.UpdateBilling(&billing); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "billing updated successfully"})
}

// DeleteBilling handles the soft delete of a billing record
func (c *OrganizationBillingController) DeleteOrganizationBilling(ctx *gin.Context) {
	billingIDParam := ctx.Param("id")
	billingID, err := uuid.FromString(billingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid billing ID"})
		return
	}

	if err := c.OrganizationBillingService.DeleteBilling(billingID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "billing deleted successfully"})
}

// GetBillingByID retrieves a single billing record by its ID
func (c *OrganizationBillingController) GetOrganizationBillingByID(ctx *gin.Context) {
	billingIDParam := ctx.Param("id")
	billingID, err := uuid.FromString(billingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid billing ID format"})
		return
	}

	billing, err := c.OrganizationBillingService.GetBillingByID(billingID)
	if err != nil {
		if err.Error() == "organization billing not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			log.Printf("Unexpected error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, billing)
}

// GetAllBillings retrieves all organization billing records
func (c *OrganizationBillingController) GetAllOrganizationBillings(ctx *gin.Context) {
	billings, err := c.OrganizationBillingService.GetAllBillings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, billings)
}
