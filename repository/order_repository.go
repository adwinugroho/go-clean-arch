package repository

import "go-clean-arch/entity"

type (
	OrderRepository interface {
		Insert(model entity.Order) error
		GetByID(id string) (*entity.Order, error)
	}
)
