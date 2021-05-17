package service

import (
	"context"
	"go-clean-arch/models/request"
	"go-clean-arch/models/response"
)

type (
	OrderService interface {
		AddData(ctx context.Context, req request.CreateOrderLRequest) *response.GeneralResponse
		GetDataByID(req request.GetByIDorderRequest) *response.GeneralResponse
	}
)
