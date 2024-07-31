package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
)

type OrderService interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrders(requestFilterBody common.RequestFilterBodyType) ([]models.Order, error)
	GetOrderByID(docId string) (models.Order, error)
	UpdateOrder(order models.Order) error
	UpdateOrderTimeout(docId string) bool
	DeleteOrder(docId string) error
}

type orderService struct {
	orderRepository *repositories.OrderRepository
}

func NewOrderService(orderRepository *repositories.OrderRepository) OrderService {
	return &orderService{orderRepository}
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	order.DocId = utils.Generate16DigitUUID()
	order.StatusCode = global_constant.ORDER_INITIATED
	createError := s.orderRepository.Create(context.Background(), order)

	event := order.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return order, createError
}

func (s *orderService) GetAllOrders(requestFilterBody common.RequestFilterBodyType) ([]models.Order, error) {
	return s.orderRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *orderService) GetOrderByID(docId string) (models.Order, error) {
	return s.orderRepository.GetByID(context.Background(), docId)
}

func (s *orderService) UpdateOrder(updateOrder models.Order) error {
	order, getByIdError := s.orderRepository.GetByID(context.Background(), updateOrder.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}

	updateError := s.orderRepository.Update(context.Background(), order.DocId, order.ToUpdatedDocument(updateOrder))

	event := order.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *orderService) UpdateOrderTimeout(docId string) bool {
	order, _ := s.orderRepository.GetByID(context.TODO(), docId)

	if order.StatusCode == global_constant.ORDER_INITIATED {
		order.StatusCode = global_constant.ORDER_TIMEOUT
		s.orderRepository.Update(context.TODO(), order.DocId, order)

		event := order.ToEvent(global_constant.OPERATION_UPDATE)
		common.AddToChanCRUD(event)
		return true
	}

	return false
}

func (s *orderService) DeleteOrder(docId string) error {
	order, getByIdError := s.orderRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}
	deleteError := s.orderRepository.Delete(context.Background(), docId)

	event := order.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}
