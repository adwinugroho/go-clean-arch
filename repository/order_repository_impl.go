package repository

import (
	"context"
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/entity"
	"go-clean-arch/logger"
	"log"

	"github.com/arangodb/go-driver"
)

type OrderRepository struct {
	DBLive driver.Database
}

func NewOrderRepository(conn *config.ArangoDB) OrderRepositoryUsecase {
	return &OrderRepository{
		DBLive: conn.DBLive,
	}
}

func (db *OrderRepository) Insert(model entity.Order) error {
	ctx := context.Background()
	order := entity.Order{}

	col, err := db.DBLive.Collection(ctx, "order_test")
	if err != nil {
		log.Printf("[Insert] Error open connection to collection, cause: %+v\n", err)
		return err
	}

	driverCtx := driver.WithReturnNew(ctx, &order)
	meta, err := col.CreateDocument(driverCtx, model)
	if err != nil {
		log.Printf("[Insert] Error while creating document, cause: %+v\n", err)
		return err
	}
	order.ID = meta.Key
	fmt.Printf("Created document with key '%s', revision '%s'\n", meta.Key, meta.Rev)

	return nil
}

func (db *OrderRepository) GetByID(id string) (*entity.Order, error) {
	ctx := context.Background()
	order := entity.Order{}

	col, err := db.DBLive.Collection(ctx, "order_test")
	if err != nil {
		log.Printf("[GetByID] Error open connection to collection, cause: %+v\n", err)
		return nil, err
	}
	_, err = col.ReadDocument(ctx, id, &order)
	if err != nil {
		logger.ErrorLogger.Println("Error reading document", err)
		return nil, err
	}

	return &order, nil
}
