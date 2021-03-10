package config

import (
	"context"
	"log"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ArangoDB struct {
	DBLive driver.Database
	DBLog  driver.Database
}

func NewArangoDBDatabase() *ArangoDB {
	ctx, cancel := NewArangoDBContext()
	defer cancel()

	// create a connection to DB
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBURL}, //DB Url
	})
	if err != nil {
		panic(err)
	}
	// create a new connection to DB client
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUsername, DBPassword),
	})
	if err != nil {
		panic(err)
	}
	// connect client and DB
	db, err := client.Database(ctx, DBName)
	if err != nil {
		log.Printf("Error connecting to database, cause: %+v\n", err)
		panic(err)
	}
	// same as DBLive, we create client for log
	connLog, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBLOGURL}, //DB Url
	})
	if err != nil {
		panic(err)
	}
	// create a new connection to DB client
	clientLog, err := driver.NewClient(driver.ClientConfig{
		Connection:     connLog,
		Authentication: driver.BasicAuthentication(DBUsername, DBPassword),
	})
	if err != nil {
		panic(err)
	}
	// connect client and DB
	dbLog, err := clientLog.Database(ctx, DBLogName)
	if err != nil {
		log.Printf("Error connecting to database, cause: %+v\n", err)
		panic(err)
	}
	return &ArangoDB{
		DBLive: db,
		DBLog:  dbLog,
	}
}

func NewArangoDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
