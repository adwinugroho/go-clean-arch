package repository

import (
	"context"
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/entity"
	"log"

	"github.com/arangodb/go-driver"
)

type AuditRepository struct {
	DBLog driver.Database
}

func NewAuditRepository(conn *config.ArangoDB) AuditRepositoryUsecase {
	return &AuditRepository{
		DBLog: conn.DBLog,
	}
}

func (db *AuditRepository) InsertLog(model entity.Audit) error {
	ctx := context.Background()
	audit := entity.Audit{}
	audit.ID = fmt.Sprintf("%s-%d", audit.ID, audit.CurrNo)
	col, err := db.DBLog.Collection(ctx, "order_test_log")
	if err != nil {
		log.Printf("[InsertLog] Error open connection to collection, cause: %+v\n", err)
		return err
	}

	driverCtx := driver.WithReturnNew(ctx, &audit)
	meta, err := col.CreateDocument(driverCtx, model)
	if err != nil {
		log.Printf("[Insertlog] Error while creating document, cause: %+v\n", err)
		return err
	}
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)

	return nil
}
