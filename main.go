package main

import (
	"go-clean-arch/config"
	route "go-clean-arch/controller"
	"go-clean-arch/repository"
	"go-clean-arch/service"
	"go-clean-arch/service/subscribe"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	getConnection := config.NewArangoDBDatabase()
	getRepository := repository.NewOrderRepository(getConnection)
	getRepositoryAudit := repository.NewAuditRepository(getConnection)
	getService := service.NewOrderService(getRepository, getRepositoryAudit)
	initRoute := route.NewOrderRoute(getService)
	e := echo.New()
	initRoute.Route(e)
	// pubsub keyspace notif expire redis
	subscribe.ValidateExpireUID()
	// pub sub channel nats (send to other services)
	subscribe.ListenNats(getRepository)
	e.Logger.Fatal(e.Start(":3003"))
}
