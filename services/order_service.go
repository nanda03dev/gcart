package services

import (
	"context"
	"github.com/nanda03dev/oms/models"
	"github.com/nanda03dev/oms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	CreateOrder(order models.Order) error
	GetAllOrders() ([]models.Order, error)
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

func (s *orderService) CreateOrder(order models.Order) error {
	return s.orderRepository.Create(context.Background(), order)
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.orderRepository.GetAll(context.Background(), nil)
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
