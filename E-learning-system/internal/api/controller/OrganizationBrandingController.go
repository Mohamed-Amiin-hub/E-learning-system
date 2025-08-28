package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OrganizationBrandingController defines the controller with its service
type OrganizationBrandingController struct {
	OrganizationBrandingService service.OrganizationBrandingService
}

// NewOrganizationBrandingController creates a new OrganizationBrandingController instance
func NewOrganizationBrandingController(brandingService service.OrganizationBrandingService) *OrganizationBrandingController {
	return &OrganizationBrandingController{OrganizationBrandingService: brandingService}
}

// CreateBranding handles the creation of a new organization branding
func (c *OrganizationBrandingController) CreateOrganizationBranding(ctx *gin.Context) {
	var branding model.OrganizationBranding

	if err := ctx.ShouldBindJSON(&branding); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBranding, err := c.OrganizationBrandingService.CreateBranding(&branding)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdBranding)
}

// UpdateBranding handles the update of an existing branding
func (c *OrganizationBrandingController) UpdateOrganizationBranding(ctx *gin.Context) {
	var branding model.OrganizationBranding

	brandingIDParam := ctx.Param("id")
	brandingID, err := uuid.FromString(brandingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branding ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&branding); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branding.ID = brandingID

	if err := c.OrganizationBrandingService.UpdateBranding(&branding); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "branding updated successfully"})
}

// DeleteBranding handles the soft delete of a branding
func (c *OrganizationBrandingController) DeleteOrganizationBranding(ctx *gin.Context) {
	brandingIDParam := ctx.Param("id")
	brandingID, err := uuid.FromString(brandingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branding ID"})
		return
	}

	if err := c.OrganizationBrandingService.DeleteBranding(brandingID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "branding deleted successfully"})
}

// GetBrandingByID retrieves a single branding by its ID
func (c *OrganizationBrandingController) GetOrganizationBrandingByID(ctx *gin.Context) {
	brandingIDParam := ctx.Param("id")
	brandingID, err := uuid.FromString(brandingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid branding ID format"})
		return
	}

	branding, err := c.OrganizationBrandingService.GetBrandingByID(brandingID)
	if err != nil {
		if err.Error() == "branding not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			log.Printf("Unexpected error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, branding)
}

// GetAllBrandings retrieves all organization brandings
func (c *OrganizationBrandingController) GetAllOrganizationBrandings(ctx *gin.Context) {
	brandings, err := c.OrganizationBrandingService.GetAllBrandings()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, brandings)
}
