package repository

import "go-clean-arch/entity"

type (
	AuditRepository interface {
		InsertLog(model entity.Audit) error
	}
)
