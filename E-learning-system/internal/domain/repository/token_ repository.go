package repository

import "e-learning-system/internal/domain/model"

// interface token
type TokenRepository interface {
	FindByToken(token string) (*model.Token, error )
	Create(token *model.Token) error

}