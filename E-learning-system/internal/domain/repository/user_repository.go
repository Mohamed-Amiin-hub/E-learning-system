package repository

import (
	"e-learning-system/internal/domain/model"

	"github.com/gofrs/uuid"
)

type UserRepository interface {
	Create(user *model.User) error
	Get(user uuid.UUID) (*model.User, error)
	Update(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	Delete(user uuid.UUID) error
	List() ([]*model.User, error)

	// password reset 
	SetResetToken(email string, token uuid.UUID, expiry string) error
	FindByResetToken(token uuid.UUID) (*model.User, error)
	UpdatePassword(userID uuid.UUID, hashedPassword string) error
	ClearResetToken(userID uuid.UUID) error
}
