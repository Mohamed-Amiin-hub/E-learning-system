package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationAdminRepository interface with required methods
type OrganizationAdminRepository interface {
	Create(organization *model.OrganizationAdmin) error
	Update(organization *model.OrganizationAdmin) error
	Delete(OrganizationAdminID uuid.UUID) error
	GetByID(OrganizationAdminID uuid.UUID) (*model.OrganizationAdmin, error)
	GetAll() ([]*model.OrganizationAdmin, error)
}
