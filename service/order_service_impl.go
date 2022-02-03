package service

import (
	"context"
	"encoding/json"
	"go-clean-arch/config"
	"go-clean-arch/entity"
	"go-clean-arch/models"
	"go-clean-arch/models/request"
	"go-clean-arch/models/response"
	"go-clean-arch/pkg/helper"
	"go-clean-arch/repository"
	"log"
)

func NewOrderService(repository repository.OrderRepository, auditRepository repository.AuditRepository) OrderService {
	return &OrderServiceImp{
		OrderRepository: repository,
		AuditRepository: auditRepository,
	}
}

type OrderServiceImp struct {
	OrderRepository repository.OrderRepository
	AuditRepository repository.AuditRepository
}

func (service *OrderServiceImp) AddData(ctx context.Context, req request.CreateOrderLRequest) error {
	var userCtx config.UserContext
	var user *request.User = ctx.Value(userCtx).(*request.User)
	// validation
	if req.Number == "" || len(req.Menus) == 0 {
		log.Println("error cause number or menus empty")
		return models.NewError(400, "Invalid Data")
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
		Owner: user.ID,
		Data: &entity.Data{
			Number: req.Number,
			Menus:  entityMenus,
		},
		Audit: &newAudit,
	}
	err := service.OrderRepository.Insert(newOrder)
	if err != nil {
		log.Println("error when insert to DB")
		return models.NewError(500, "Internal Server Error, Please Contact Customer Service")
	}
	// pub to channel
	pub, err := config.ConnectNats()
	if err != nil {
		log.Println("error when connection to nats")
		return models.NewError(500, "Internal Server Error, Please Contact Customer Service")
	}
	plPublish := entity.Order{
		ID:   newOrder.ID,
		Data: newOrder.Data,
	}
	plBytes, err := json.Marshal(plPublish)
	if err != nil {
		log.Println("error when marshalling payload for publish")
		return models.NewError(500, "Internal Server Error, Please Contact Customer Service")
	}
	if err := pub.Stan.Publish(config.CH_ORDER, plBytes); err != nil {
		log.Println("error cause can't publish to channel order")
		return models.NewError(500, "Internal Server Error, Please Contact Customer Service")
	}

	return nil
}

func (service *OrderServiceImp) GetDataByID(ctx context.Context, req request.GetByIDorderRequest) (*response.GetByIDOrderResponse, error) {
	var userCtx config.UserContext
	var user *request.User = ctx.Value(userCtx).(*request.User)
	data, err := service.OrderRepository.GetByID(req.ID, user.ID)
	if err != nil {
		log.Println("error when get by id")
		return nil, models.NewError(500, "Internal Server Error, Please Contact Customer Service")
	} else if data == nil || req.ID == "" {
		log.Println("id not found")
		return nil, models.NewError(404, "Data Not Found")
	}
	resp := response.GetByIDOrderResponse{
		ID:     data.ID,
		Number: data.Data.Number,
		Menus:  data.Data.Menus,
	}
	return &resp, nil
}
