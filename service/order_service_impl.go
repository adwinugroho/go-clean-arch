package service

import (
	"context"
	"encoding/json"
	"go-clean-arch/config"
	"go-clean-arch/entity"
	"go-clean-arch/logger"
	"go-clean-arch/models/request"
	"go-clean-arch/models/response"
	"go-clean-arch/pkg/helper"
	"go-clean-arch/repository"
)

func NewOrderService(repository repository.OrderRepositoryUsecase, auditRepository repository.AuditRepositoryUsecase) OrderServiceUsecase {
	return &OrderService{
		OrderRepository: repository,
		AuditRepository: auditRepository,
	}
}

type OrderService struct {
	OrderRepository repository.OrderRepositoryUsecase
	AuditRepository repository.AuditRepositoryUsecase
}

func (service *OrderService) AddData(ctx context.Context, req request.CreateOrderLRequest) *response.GeneralResponse {
	// validation
	if req.Number == "" || len(req.Menus) == 0 {
		logger.ErrorLogger.Println("error cause number or menus empty")
		return response.Error(400, "Invalid Data")
	}

	// convert menu in request to menu in entity
	var entityMenus []entity.Menu
	for _, menuInReq := range req.Menus {
		var entityMenu = entity.Menu{
			Name:                  menuInReq.Name,
			Qty:                   menuInReq.Qty,
			AdditionalInformation: menuInReq.AdditionalInformation,
		}

		entityMenus = append(entityMenus, entityMenu)
	}

	newAudit := entity.Audit{
		CreatedAt: helper.TimeHostNow().Format("20060102150405"),
		UpdatedAt: "",
		DeletedAt: "",
		CurrNo:    1,
	}

	newOrder := entity.Order{
		Data: &entity.Data{
			Number: req.Number,
			Menus:  entityMenus,
		},
		Audit: &newAudit,
	}
	err := service.OrderRepository.Insert(newOrder)
	if err != nil {
		logger.ErrorLogger.Println("error when insert to DB")
		return response.Error(500, "Internal Server Error, Please Contact Customer Service")
	}
	// pub to channel
	pub, err := config.ConnectNats()
	if err != nil {
		logger.ErrorLogger.Println("error when connection to nats")
		return response.Error(500, "Internal Server Error, Please Contact Customer Service")
	}
	plPublish := entity.Order{
		ID:   newOrder.ID,
		Data: newOrder.Data,
	}
	plBytes, err := json.Marshal(plPublish)
	if err != nil {
		logger.ErrorLogger.Println("error when marshalling payload for publish")
		return response.Error(500, "Internal Server Error, Please Contact Customer Service")
	}
	if err := pub.Stan.Publish(config.CH_ORDER, plBytes); err != nil {
		logger.ErrorLogger.Println("error cause can't publish to channel order")
		return response.Error(500, "Internal Server Error, Please Contact Customer Service")
	}

	return response.Success(200, nil)
}

func (service *OrderService) GetDataByID(req request.GetByIDorderRequest) *response.GeneralResponse {
	data, err := service.OrderRepository.GetByID(req.ID)
	if err != nil {
		logger.ErrorLogger.Println("error when get by id")
		return response.Error(500, "Internal Server Error, Please Contact Customer Service")
	} else if data == nil || req.ID == "" {
		logger.ErrorLogger.Println("id not found")
		return response.Error(404, "Data Not Found")
	}
	resp := response.GetByIDOrderResponse{
		ID:     data.ID,
		Number: data.Data.Number,
		Menus:  data.Data.Menus,
	}
	return response.Success(200, resp)
}
