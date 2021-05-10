package repository

import "go-clean-arch/entity"

type (
	OrderRepositoryUsecase interface {
		Insert(model entity.Order) error
		GetByID(id string) (*entity.Order, error)
	}
)
