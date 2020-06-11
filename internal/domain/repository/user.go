package repository

import "damihailov85/kurovo.run/internal/domain/entity"

type Users interface {
	Get(id string) (entity.User, error)
	ListAll() ([]entity.User, error)
	Save(u entity.User) error
	Block(id string, reason string) error
}
