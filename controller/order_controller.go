package route

import (
	"context"
	"go-clean-arch/models/request"
	"go-clean-arch/service"

	"github.com/labstack/echo/v4"
)

type (
	OrderRoute struct {
		service service.OrderServiceUsecase
	}
)

func NewOrderRoute(service service.OrderServiceUsecase) OrderRoute {
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
	response := route.service.AddData(ctx, *body)
	return c.JSON(response.Code, response)
}

func (route *OrderRoute) Get(c echo.Context) error {
	body := new(request.GetByIDorderRequest)
	c.Bind(body)
	response := route.service.GetDataByID(*body)
	return c.JSON(response.Code, response)
}
