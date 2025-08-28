package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrganizationService interface with CRUD methods
type OrganizationService interface {
	CreateOrganization(org *model.Organization) (*model.Organization, error)
	UpdateOrganization(org *model.Organization) error
	DeleteOrganization(orgID uuid.UUID) error
	GetOrganizationByID(orgID uuid.UUID) (*model.Organization, error)
	GetAllOrganizations() ([]*model.Organization, error)
}

// organizationServiceImpl struct implementing OrganizationService
type organizationServiceImpl struct {
	repo repository.OrganizationRepository
}

// Constructor
func NewOrganizationService(orgRepo repository.OrganizationRepository) OrganizationService {
	return &organizationServiceImpl{
		repo: orgRepo,
	}
}

// CreateOrganization creates a new organization
func (s *organizationServiceImpl) CreateOrganization(org *model.Organization) (*model.Organization, error) {
	// Generate a new UUID for the organization
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	org.ID = newID
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	// Log the creation attempt
	log.Printf("Creating organization: %+v", org)

	// Save to repository
	if err := s.repo.Create(org); err != nil {
		return nil, fmt.Errorf("failed to create organization: %v", err)
	}

	return org, nil
}

// UpdateOrganization updates an existing organization
func (s *organizationServiceImpl) UpdateOrganization(org *model.Organization) error {
	org.UpdatedAt = time.Now()

	// Check if organization exists
	_, err := s.repo.GetByID(org.ID)
	if err != nil {
		return fmt.Errorf("organization not found with ID %s: %v", org.ID, err)
	}

	if err := s.repo.Update(org); err != nil {
		return fmt.Errorf("failed to update organization with ID %s: %v", org.ID, err)
	}

	log.Printf("Organization updated: %+v", org)
	return nil
}

// DeleteOrganization performs a soft delete
func (s *organizationServiceImpl) DeleteOrganization(orgID uuid.UUID) error {
	// Check if organization exists
	_, err := s.repo.GetByID(orgID)
	if err != nil {
		return fmt.Errorf("organization not found with ID %s: %v", orgID, err)
	}

	if err := s.repo.Delete(orgID); err != nil {
		return fmt.Errorf("failed to delete organization with ID %s: %v", orgID, err)
	}

	log.Printf("Organization soft-deleted: %v", orgID)
	return nil
}

// GetOrganizationByID retrieves a single organization
func (s *organizationServiceImpl) GetOrganizationByID(orgID uuid.UUID) (*model.Organization, error) {
	org, err := s.repo.GetByID(orgID)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization by ID %s: %v", orgID, err)
	}
	return org, nil
}

// GetAllOrganizations retrieves all organizations
func (s *organizationServiceImpl) GetAllOrganizations() ([]*model.Organization, error) {
	orgs, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all organizations: %v", err)
	}
	return orgs, nil
}
