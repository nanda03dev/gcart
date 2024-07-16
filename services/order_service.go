package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/workers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrders(requestFilterBody common.RequestFilterBodyType) ([]models.Order, error)
	GetOrderByID(id string) (models.Order, error)
	UpdateOrder(order models.Order) error
	DeleteOrder(id string) error
}

type orderService struct {
	orderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) OrderService {
	return &orderService{orderRepository}
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	order.ID = primitive.NewObjectID()
	order.Code = global_constant.OrderSuccessCode.ORDER_INITIATED
	err := s.orderRepository.Create(context.Background(), order)
	event := common.EventType{
		EntityId:      order.ID,
		EntityType:    global_constant.Entities.Order,
		OperationType: global_constant.Operations.Create,
	}
	workers.AddToChanCRUD(event)

	return order, err
}

func (s *orderService) GetAllOrders(requestFilterBody common.RequestFilterBodyType) ([]models.Order, error) {
	return s.orderRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *orderService) GetOrderByID(id string) (models.Order, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Order{}, err
	}

	return s.orderRepository.GetByID(context.Background(), objectId)
}

func (s *orderService) UpdateOrder(order models.Order) error {
	return s.orderRepository.Update(context.Background(), order.ID, order)
}

func (s *orderService) DeleteOrder(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.orderRepository.Delete(context.Background(), objectId)
}
