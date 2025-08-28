package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrganizationBrandingService interface with CRUD methods
type OrganizationBrandingService interface {
	CreateBranding(branding *model.OrganizationBranding) (*model.OrganizationBranding, error)
	UpdateBranding(branding *model.OrganizationBranding) error
	DeleteBranding(brandingID uuid.UUID) error
	GetBrandingByID(brandingID uuid.UUID) (*model.OrganizationBranding, error)
	GetAllBrandings() ([]*model.OrganizationBranding, error)
}

// organizationBrandingServiceImpl struct implementing OrganizationBrandingService
type organizationBrandingServiceImpl struct {
	repo repository.OrganizationBrandingRepository
}

// Constructor
func NewOrganizationBrandingService(brandingRepo repository.OrganizationBrandingRepository) OrganizationBrandingService {
	return &organizationBrandingServiceImpl{
		repo: brandingRepo,
	}
}

// CreateBranding creates a new branding record
func (s *organizationBrandingServiceImpl) CreateBranding(branding *model.OrganizationBranding) (*model.OrganizationBranding, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	branding.ID = newID
	branding.CreatedAt = time.Now()
	branding.UpdatedAt = time.Now()

	log.Printf("Creating organization branding: %+v", branding)

	if err := s.repo.Create(branding); err != nil {
		return nil, fmt.Errorf("failed to create branding: %v", err)
	}

	return branding, nil
}

// UpdateBranding updates an existing branding record
func (s *organizationBrandingServiceImpl) UpdateBranding(branding *model.OrganizationBranding) error {
	branding.UpdatedAt = time.Now()

	// Check if branding exists
	_, err := s.repo.GetByID(branding.ID)
	if err != nil {
		return fmt.Errorf("branding not found with ID %s: %v", branding.ID, err)
	}

	if err := s.repo.Update(branding); err != nil {
		return fmt.Errorf("failed to update branding with ID %s: %v", branding.ID, err)
	}

	log.Printf("Organization branding updated: %+v", branding)
	return nil
}

// DeleteBranding performs a soft delete
func (s *organizationBrandingServiceImpl) DeleteBranding(brandingID uuid.UUID) error {
	// Check if branding exists
	_, err := s.repo.GetByID(brandingID)
	if err != nil {
		return fmt.Errorf("branding not found with ID %s: %v", brandingID, err)
	}

	if err := s.repo.Delete(brandingID); err != nil {
		return fmt.Errorf("failed to delete branding with ID %s: %v", brandingID, err)
	}

	log.Printf("Organization branding deleted: %v", brandingID)
	return nil
}

// GetBrandingByID retrieves a single branding record by ID
func (s *organizationBrandingServiceImpl) GetBrandingByID(brandingID uuid.UUID) (*model.OrganizationBranding, error) {
	branding, err := s.repo.GetByID(brandingID)
	if err != nil {
		return nil, fmt.Errorf("failed to get branding by ID %s: %v", brandingID, err)
	}
	return branding, nil
}

// GetAllBrandings retrieves all branding records
func (s *organizationBrandingServiceImpl) GetAllBrandings() ([]*model.OrganizationBranding, error) {
	brandings, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all brandings: %v", err)
	}
	return brandings, nil
}
