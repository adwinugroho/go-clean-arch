package repository

import "go-clean-arch/entity"

type (
	AuditRepositoryUsecase interface {
		InsertLog(model entity.Audit) error
	}
)
