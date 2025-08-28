package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationBillingRepository interface with required methods
type OrganizationBillingRepository interface {
	Create(OrganizationBilling *model.OrganizationBilling) error
	Update(OrganizationBilling *model.OrganizationBilling) error
	Delete(OrganizationBrandingID uuid.UUID) error
	GetByID(OrganizationBrandingID uuid.UUID) (*model.OrganizationBilling, error)
	GetAll() ([]*model.OrganizationBilling, error)
}
