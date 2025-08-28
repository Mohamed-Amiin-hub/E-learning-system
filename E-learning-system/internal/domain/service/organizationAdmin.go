package service

import (
	"e-learning-system/internal/domain/model"
	"e-learning-system/internal/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// OrganizationAdminService interface with CRUD methods
type OrganizationAdminService interface {
	CreateAdmin(admin *model.OrganizationAdmin) (*model.OrganizationAdmin, error)
	UpdateAdmin(admin *model.OrganizationAdmin) error
	DeleteAdmin(adminID uuid.UUID) error
	GetAdminByID(adminID uuid.UUID) (*model.OrganizationAdmin, error)
	GetAllAdmins() ([]*model.OrganizationAdmin, error)
}

// organizationAdminServiceImpl struct implementing OrganizationAdminService
type organizationAdminServiceImpl struct {
	repo repository.OrganizationAdminRepository
}

// Constructor
func NewOrganizationAdminService(adminRepo repository.OrganizationAdminRepository) OrganizationAdminService {
	return &organizationAdminServiceImpl{
		repo: adminRepo,
	}
}

// CreateAdmin creates a new organization admin
func (s *organizationAdminServiceImpl) CreateAdmin(admin *model.OrganizationAdmin) (*model.OrganizationAdmin, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}

	admin.ID = newID
	admin.CreatedAt = time.Now()

	log.Printf("Creating organization admin: %+v", admin)

	if err := s.repo.Create(admin); err != nil {
		return nil, fmt.Errorf("failed to create admin: %v", err)
	}

	return admin, nil
}

// UpdateAdmin updates an existing admin
func (s *organizationAdminServiceImpl) UpdateAdmin(admin *model.OrganizationAdmin) error {
	// Check if admin exists
	_, err := s.repo.GetByID(admin.ID)
	if err != nil {
		return fmt.Errorf("admin not found with ID %s: %v", admin.ID, err)
	}

	if err := s.repo.Update(admin); err != nil {
		return fmt.Errorf("failed to update admin with ID %s: %v", admin.ID, err)
	}

	log.Printf("Organization admin updated: %+v", admin)
	return nil
}

// DeleteAdmin performs a soft delete
func (s *organizationAdminServiceImpl) DeleteAdmin(adminID uuid.UUID) error {
	// Check if admin exists
	_, err := s.repo.GetByID(adminID)
	if err != nil {
		return fmt.Errorf("admin not found with ID %s: %v", adminID, err)
	}

	if err := s.repo.Delete(adminID); err != nil {
		return fmt.Errorf("failed to delete admin with ID %s: %v", adminID, err)
	}

	log.Printf("Organization admin deleted: %v", adminID)
	return nil
}

// GetAdminByID retrieves a single admin by ID
func (s *organizationAdminServiceImpl) GetAdminByID(adminID uuid.UUID) (*model.OrganizationAdmin, error) {
	admin, err := s.repo.GetByID(adminID)
	if err != nil {
		return nil, fmt.Errorf("failed to get admin by ID %s: %v", adminID, err)
	}
	return admin, nil
}

// GetAllAdmins retrieves all organization admins
func (s *organizationAdminServiceImpl) GetAllAdmins() ([]*model.OrganizationAdmin, error) {
	admins, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all admins: %v", err)
	}
	return admins, nil
}
