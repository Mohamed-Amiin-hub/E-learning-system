package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

// OrganizationTutorRepository interface with required methods
type OrganizationTutorRepository interface {
	Create(OrganizationTutor *model.OrganizationTutor) error
	Update(OrganizationTutor *model.OrganizationTutor) error
	Delete(OrganizationTutorID uuid.UUID) error
	GetByID(OrganizationTutorID uuid.UUID) (*model.OrganizationTutor, error)
	GetAll() ([]*model.OrganizationTutor, error)
}
