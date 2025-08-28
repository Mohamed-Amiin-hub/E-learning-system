package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OrganizationController defines the organization controller with its service
type OrganizationController struct {
	OrganizationService service.OrganizationService
}

// NewOrganizationController creates a new OrganizationController instance
func NewOrganizationController(orgService service.OrganizationService) *OrganizationController {
	return &OrganizationController{OrganizationService: orgService}
}

// CreateOrganization handles the creation of a new organization
func (c *OrganizationController) CreateOrganization(ctx *gin.Context) {
	var org model.Organization

	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdOrg, err := c.OrganizationService.CreateOrganization(&org)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdOrg)
}

// UpdateOrganization handles the update of an existing organization
func (c *OrganizationController) UpdateOrganization(ctx *gin.Context) {
	var org model.Organization

	orgIDParam := ctx.Param("id")
	orgID, err := uuid.FromString(orgIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&org); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	org.ID = orgID

	if err := c.OrganizationService.UpdateOrganization(&org); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "organization updated successfully"})
}

// DeleteOrganization handles the soft delete of an organization
func (c *OrganizationController) DeleteOrganization(ctx *gin.Context) {
	orgIDParam := ctx.Param("id")
	orgID, err := uuid.FromString(orgIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	if err := c.OrganizationService.DeleteOrganization(orgID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "organization deleted successfully"})
}

// GetOrganizationByID retrieves a single organization by its ID
func (c *OrganizationController) GetOrganizationByID(ctx *gin.Context) {
	orgIDParam := ctx.Param("id")
	orgID, err := uuid.FromString(orgIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization ID"})
		return
	}

	org, err := c.OrganizationService.GetOrganizationByID(orgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, org)
}

// GetAllOrganizations retrieves all organizations
func (c *OrganizationController) GetAllOrganizations(ctx *gin.Context) {
	orgs, err := c.OrganizationService.GetAllOrganizations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orgs)
}
