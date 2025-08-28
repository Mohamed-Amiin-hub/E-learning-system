package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationBrandingRepository interface with required methods
type OrganizationBrandingRepository interface {
	Create(OrganizationBranding *model.OrganizationBranding) error
	Update(OrganizationBranding *model.OrganizationBranding) error
	Delete(OrganizationBrandingID uuid.UUID) error
	GetByID(OrganizationBrandingID uuid.UUID) (*model.OrganizationBranding, error)
	GetAll() ([]*model.OrganizationBranding, error)
}
