package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationRepository interface with required methods
type AdminRepository interface {
	Create(Admin *model.Admin) error
	Update(Admin *model.Admin) error
	Delete(AdminID uuid.UUID) error
	GetByID(AdminID uuid.UUID) (*model.Admin, error)
	GetAll() ([]*model.Admin, error)
}
