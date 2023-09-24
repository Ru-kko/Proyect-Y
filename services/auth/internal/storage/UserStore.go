package storage

import "Proyect-Y/typo"

type UserStore interface {
	Get(string) (*typo.AuthData, error)
	GetByUserTag(string) (*typo.AuthData, error)
	Save(typo.AuthData) (*typo.AuthData, error)
	Edit(typo.AuthData) (*typo.AuthData, error)
	Delete(string) error
	Close() error
}
