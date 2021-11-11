package repository

import "go-clean-arch/entity"

type (
	OrderRepository interface {
		Insert(model entity.Order) error
		GetByID(id, owner string) (*entity.Order, error)
		DeleteByID(id string) error
	}
)
