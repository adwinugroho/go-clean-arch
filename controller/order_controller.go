package controller

import (
	"context"
	"go-clean-arch/models/request"
	resModel "go-clean-arch/models/response"
	"go-clean-arch/service"
	"go-clean-arch/service/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	OrderRoute struct {
		service service.OrderService
	}
)

func NewOrderRoute(service service.OrderService) OrderRoute {
	return OrderRoute{service: service}
}

func (route *OrderRoute) Route(e *echo.Echo) {
	orderRoute := e.Group("/order")
	orderRoute.POST("/new", route.New)
	orderRoute.POST("/get", route.Get)
}

func (route *OrderRoute) New(c echo.Context) error {
	ctx := context.Background()
	body := new(request.CreateOrderLRequest)
	c.Bind(body)
	messageValidate, errValidate := validation.ValidateCreateOrder(*body)
	if errValidate != nil {
		return c.JSON(http.StatusOK, resModel.Error(400, messageValidate))
	}
	response := route.service.AddData(ctx, *body)
	return c.JSON(response.Code, response)
}

func (route *OrderRoute) Get(c echo.Context) error {
	body := new(request.GetByIDorderRequest)
	c.Bind(body)
	messageValidate, errValidate := validation.ValidateGetOrder(*body)
	if errValidate != nil {
		return c.JSON(http.StatusOK, resModel.Error(400, messageValidate))
	}
	response := route.service.GetDataByID(*body)
	return c.JSON(response.Code, response)
}
