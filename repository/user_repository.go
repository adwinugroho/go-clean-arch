package repository

import (
	"context"
	"go-clean-arch/entity"
)

type UserRepository interface {
	Add(ctx context.Context, user entity.User) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	DeleteWithReturnOld(ctx context.Context, id string) (*entity.User, error)
	GetDetailByID(ctx context.Context, id string) (*entity.User, error)
	GetList(ctx context.Context, vars entity.BindVarsUser) (*[]entity.User, error)
	Update(user entity.User) (*entity.User, error)
}
