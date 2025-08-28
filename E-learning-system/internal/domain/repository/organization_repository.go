package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationRepository interface with required methods
type OrganizationRepository interface {
	Create(organization *model.Organization) error
	Update(organization *model.Organization) error
	Delete(organizationID uuid.UUID) error
	GetByID(organizationID uuid.UUID) (*model.Organization, error)
	GetAll() ([]*model.Organization, error)
}

