package main

import (
	"go-clean-arch/config"
	route "go-clean-arch/controller"
	"go-clean-arch/repository"
	"go-clean-arch/service"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	getConnection := config.NewArangoDBDatabase()
	getRepository := repository.NewOrderRepository(getConnection)
	getRepositoryAudit := repository.NewAuditRepository(getConnection)
	getService := service.NewOrderService(getRepository, getRepositoryAudit)
	initRoute := route.NewOrderRoute(getService)
	e := echo.New()
	initRoute.Route(e)
	e.Logger.Fatal(e.Start(":3003"))
}
