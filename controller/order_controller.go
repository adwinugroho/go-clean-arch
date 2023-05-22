package controller

import (
	"context"
	"go-clean-arch/config"
	"go-clean-arch/models"
	"go-clean-arch/models/request"
	"go-clean-arch/models/response"
	"go-clean-arch/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var userContext config.UserContext

type (
	OrderRoute struct {
		service   service.OrderService
		user      *request.User
		Validator *validator.Validate
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
	errValidate := route.Validate(body)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, errValidate)
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
	errValidate := route.Validate(body)
	if errValidate != nil {
		return c.JSON(http.StatusBadRequest, errValidate)
	}
	ctxService := context.WithValue(context.Background(), userContext, route.user)
	respData, err := route.service.GetDataByID(ctxService, *body)
	if err != nil {
		return c.JSON(http.StatusOK, err.(*models.Error))
	}
	return c.JSON(http.StatusOK, response.Success(200).SetData(respData))
}
