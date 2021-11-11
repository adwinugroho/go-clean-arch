package repository

import (
	"context"
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/entity"
	"log"

	"github.com/arangodb/go-driver"
)

type OrderRepositoryImp struct {
	DBLive driver.Database
}

func NewOrderRepository(conn *config.ArangoDB) OrderRepository {
	return &OrderRepositoryImp{
		DBLive: conn.DBLive,
	}
}

func (db *OrderRepositoryImp) Insert(model entity.Order) error {
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

func (db *OrderRepositoryImp) DeleteByID(id string) error {
	ctx := context.Background()
	col, err := db.DBLive.Collection(ctx, "order_test")
	if err != nil {
		log.Printf("[DeleteByID] Error open connection to collection, cause: %+v\n", err)
		return err
	}
	_, err = col.RemoveDocument(ctx, id)
	if err != nil {
		log.Println("Error reading document")
		return err
	}
	return nil
}

func (db *OrderRepositoryImp) GetByID(id, owner string) (*entity.Order, error) {
	ctx := context.Background()
	order := entity.Order{}
	var bindVars = map[string]interface{}{
		"owner": owner,
		"id":    id,
	}
	query := `FOR x IN order_test FILTER x.owner == @owner AND x._key == @id RETURN x`
	cursor, err := db.DBLive.Query(ctx, query, bindVars)
	if err != nil {
		log.Printf("[GetByID] Error open connection to collection, cause: %+v\n", err)
		return nil, err
	}
	_, err = cursor.ReadDocument(ctx, &order)
	if driver.IsNoMoreDocuments(err) {
		log.Println("data nil or id not found")
		return nil, nil
	} else if err != nil {
		log.Println("Error reading document")
		return nil, err
	}

	return &order, nil
}
