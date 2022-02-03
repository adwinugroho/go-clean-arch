package controller

import (
	"context"
	"go-clean-arch/config"
	"go-clean-arch/models"
	"go-clean-arch/models/request"
	"go-clean-arch/models/response"
	"go-clean-arch/service"
	"go-clean-arch/service/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

var userContext config.UserContext

type (
	OrderRoute struct {
		service service.OrderService
		user    *request.User
	}
)

func NewOrderRoute(service service.OrderService) OrderRoute {
	return OrderRoute{service: service}
}

func (route *OrderRoute) Route(e *echo.Echo) {
	orderRoute := e.Group("/order")
	orderRoute.Use(route.middlewareOrder)
	orderRoute.POST("/new", route.New)
	orderRoute.POST("/get", route.Get)
}

func (route *OrderRoute) New(c echo.Context) error {
	ctx := context.Background()
	body := new(request.CreateOrderLRequest)
	c.Bind(body)
	messageValidate, errValidate := validation.ValidateCreateOrder(*body)
	if errValidate != nil {
		return c.JSON(http.StatusOK, models.NewError(400, messageValidate))
	}
	ctxService := context.WithValue(ctx, userContext, route.user)
	err := route.service.AddData(ctxService, *body)
	if err != nil {
		return c.JSON(http.StatusOK, err.(*models.Error))
	}
	return c.JSON(http.StatusOK, response.Success(200))
}

func (route *OrderRoute) Get(c echo.Context) error {
	body := new(request.GetByIDorderRequest)
	c.Bind(body)
	messageValidate, errValidate := validation.ValidateGetOrder(*body)
	if errValidate != nil {
		return c.JSON(http.StatusOK, models.NewError(400, messageValidate))
	}
	ctxService := context.WithValue(context.Background(), userContext, route.user)
	respData, err := route.service.GetDataByID(ctxService, *body)
	if err != nil {
		return c.JSON(http.StatusOK, err.(*models.Error))
	}
	return c.JSON(http.StatusOK, response.Success(200).SetData(respData))
}
