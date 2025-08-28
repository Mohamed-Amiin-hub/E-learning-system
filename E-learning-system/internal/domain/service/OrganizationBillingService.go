package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrganizationBillingService interface with CRUD methods
type OrganizationBillingService interface {
	CreateBilling(billing *model.OrganizationBilling) (*model.OrganizationBilling, error)
	UpdateBilling(billing *model.OrganizationBilling) error
	DeleteBilling(billingID uuid.UUID) error
	GetBillingByID(billingID uuid.UUID) (*model.OrganizationBilling, error)
	GetAllBillings() ([]*model.OrganizationBilling, error)
}

// organizationBillingServiceImpl struct implementing OrganizationBillingService
type organizationBillingServiceImpl struct {
	repo repository.OrganizationBillingRepository
}

// Constructor
func NewOrganizationBillingService(billingRepo repository.OrganizationBillingRepository) OrganizationBillingService {
	return &organizationBillingServiceImpl{
		repo: billingRepo,
	}
}

// CreateBilling creates a new billing record
func (s *organizationBillingServiceImpl) CreateBilling(billing *model.OrganizationBilling) (*model.OrganizationBilling, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	billing.ID = newID
	billing.CreatedAt = time.Now()
	billing.UpdatedAt = time.Now()

	log.Printf("Creating organization billing: %+v", billing)

	if err := s.repo.Create(billing); err != nil {
		return nil, fmt.Errorf("failed to create billing: %v", err)
	}

	return billing, nil
}

// UpdateBilling updates an existing billing record
func (s *organizationBillingServiceImpl) UpdateBilling(billing *model.OrganizationBilling) error {
	billing.UpdatedAt = time.Now()

	// Check if billing exists
	_, err := s.repo.GetByID(billing.ID)
	if err != nil {
		return fmt.Errorf("billing not found with ID %s: %v", billing.ID, err)
	}

	if err := s.repo.Update(billing); err != nil {
		return fmt.Errorf("failed to update billing with ID %s: %v", billing.ID, err)
	}

	log.Printf("Organization billing updated: %+v", billing)
	return nil
}

// DeleteBilling performs a soft delete
func (s *organizationBillingServiceImpl) DeleteBilling(billingID uuid.UUID) error {
	// Check if billing exists
	_, err := s.repo.GetByID(billingID)
	if err != nil {
		return fmt.Errorf("billing not found with ID %s: %v", billingID, err)
	}

	if err := s.repo.Delete(billingID); err != nil {
		return fmt.Errorf("failed to delete billing with ID %s: %v", billingID, err)
	}

	log.Printf("Organization billing deleted: %v", billingID)
	return nil
}

// GetBillingByID retrieves a single billing record by ID
func (s *organizationBillingServiceImpl) GetBillingByID(billingID uuid.UUID) (*model.OrganizationBilling, error) {
	billing, err := s.repo.GetByID(billingID)
	if err != nil {
		return nil, fmt.Errorf("failed to get billing by ID %s: %v", billingID, err)
	}
	return billing, nil
}

// GetAllBillings retrieves all billing records
func (s *organizationBillingServiceImpl) GetAllBillings() ([]*model.OrganizationBilling, error) {
	billings, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all billings: %v", err)
	}
	return billings, nil
}
