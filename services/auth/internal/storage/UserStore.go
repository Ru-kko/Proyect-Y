package storage

import (
	"Proyect-Y/auth-service/internal/domain"
)

type UserStore interface {
	Get(string) (*domain.StoredUser, error)
	GetByUserTag(string) (*domain.StoredUser, error)
	Save(domain.StoredUser) (*domain.StoredUser, error)
	Edit(domain.StoredUser) (*domain.StoredUser, error)
	Delete(string) error
	Close() error
}
