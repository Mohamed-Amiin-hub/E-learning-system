package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OrganizationAdminController defines the controller with its service
type OrganizationAdminController struct {
	OrganizationAdminService service.OrganizationAdminService
}

// NewOrganizationAdminController creates a new OrganizationAdminController instance
func NewOrganizationAdminController(adminService service.OrganizationAdminService) *OrganizationAdminController {
	return &OrganizationAdminController{OrganizationAdminService: adminService}
}

// CreateAdmin handles the creation of a new organization admin
func (c *OrganizationAdminController) CreateOrganizationAdmin(ctx *gin.Context) {
	var admin model.OrganizationAdmin

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAdmin, err := c.OrganizationAdminService.CreateAdmin(&admin)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdAdmin)
}

// UpdateAdmin handles updating an existing admin
func (c *OrganizationAdminController) UpdateOrganizationAdmin(ctx *gin.Context) {
	var admin model.OrganizationAdmin

	adminIDParam := ctx.Param("id")
	adminID, err := uuid.FromString(adminIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin.ID = adminID

	if err := c.OrganizationAdminService.UpdateAdmin(&admin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "admin updated successfully"})
}

// DeleteAdmin handles the soft delete of an admin
func (c *OrganizationAdminController) DeleteOrganizationAdmin(ctx *gin.Context) {
	adminIDParam := ctx.Param("id")
	adminID, err := uuid.FromString(adminIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin ID"})
		return
	}

	if err := c.OrganizationAdminService.DeleteAdmin(adminID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "admin deleted successfully"})
}

// GetAdminByID retrieves a single admin by its ID
func (c *OrganizationAdminController) GetOrganizationAdminByID(ctx *gin.Context) {
	adminIDParam := ctx.Param("id")
	adminID, err := uuid.FromString(adminIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin ID format"})
		return
	}

	admin, err := c.OrganizationAdminService.GetAdminByID(adminID)
	if err != nil {
		if err.Error() == "admin not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			log.Printf("Unexpected error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

// GetAllAdmins retrieves all organization admins
func (c *OrganizationAdminController) GetAllOrganizationAdmins(ctx *gin.Context) {
	admins, err := c.OrganizationAdminService.GetAllAdmins()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, admins)
}
