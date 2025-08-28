package controller

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// OrganizationTutorController defines the tutor controller with its service
type OrganizationTutorController struct {
	OrganizationTutorService service.OrganizationTutorService
}

// NewOrganizationTutorController creates a new OrganizationTutorController instance
func NewOrganizationTutorController(tutorService service.OrganizationTutorService) *OrganizationTutorController {
	return &OrganizationTutorController{OrganizationTutorService: tutorService}
}

// CreateTutor handles the creation of a new organization tutor
func (c *OrganizationTutorController) CreateTutorOrganization(ctx *gin.Context) {
	var tutor model.OrganizationTutor

	if err := ctx.ShouldBindJSON(&tutor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTutor, err := c.OrganizationTutorService.CreateTutor(&tutor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdTutor)
}

// UpdateTutor handles the update of an existing tutor
func (c *OrganizationTutorController) UpdateTutorOrganization(ctx *gin.Context) {
	var tutor model.OrganizationTutor

	tutorIDParam := ctx.Param("id")
	tutorID, err := uuid.FromString(tutorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutor ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&tutor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tutor.ID = tutorID

	if err := c.OrganizationTutorService.UpdateTutor(&tutor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "tutor updated successfully"})
}

// DeleteTutor handles the soft delete of a tutor
func (c *OrganizationTutorController) DeleteTutorOrganization(ctx *gin.Context) {
	tutorIDParam := ctx.Param("id")
	tutorID, err := uuid.FromString(tutorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutor ID"})
		return
	}

	if err := c.OrganizationTutorService.DeleteTutor(tutorID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "tutor deleted successfully"})
}

// GetTutorByID retrieves a single tutor by its ID
func (c *OrganizationTutorController) GetTutorOrganizationByID(ctx *gin.Context) {
	tutorIDParam := ctx.Param("id")
	tutorID, err := uuid.FromString(tutorIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutor ID"})
		return
	}

	tutor, err := c.OrganizationTutorService.GetTutorByID(tutorID)
	if err != nil {
		if err.Error() == "tutor not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			log.Printf("Unexpected error: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, tutor)
}

// GetAllTutors retrieves all organization tutors
func (c *OrganizationTutorController) GetAllTutorsOrganization(ctx *gin.Context) {
	tutors, err := c.OrganizationTutorService.GetAllTutors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tutors)
}
