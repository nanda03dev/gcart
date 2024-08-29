package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/global_constant"
	"github.com/nanda03dev/gcart/src/models"
	"github.com/nanda03dev/gcart/src/repositories"
)

type OrderService interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetAllOrders(requestFilterBody common.RequestFilterBodyType) ([]models.Order, error)
	GetOrderByID(docId string) (models.Order, error)
	UpdateOrder(order models.Order) error
	ConfirmOrder(orderConfirmBody common.OrderConfirmBody) error
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
	order.DocId = models.Generate16DigitUUID()
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
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
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
		common.AddToChanPaymentRefund(event)
		return true
	}

	return false
}

func (s *orderService) DeleteOrder(docId string) error {
	order, getByIdError := s.orderRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}

	deleteError := s.orderRepository.Delete(context.Background(), docId)

	event := order.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	AppServices.Item.DeleteOrderItems(order.DocId)
	AppServices.Payment.DeleteOrderPayments(order.DocId)

	return deleteError
}

func (s *orderService) ConfirmOrder(orderConfirmBody common.OrderConfirmBody) error {
	order, getByIdError := s.orderRepository.GetByID(context.Background(), orderConfirmBody.OrderId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}

	orderPayments, paymentGetError := AppServices.Payment.GetAllPaymentsByOrderId(order.DocId)

	if paymentGetError != nil {
		return errors.New(global_constant.ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_PAYMENT_PENDING)
	}

	var orderAmount = order.Amount
	var paymentTotalAmount int

	for _, payment := range orderPayments {
		if payment.StatusCode == global_constant.PAYMENT_CONFIRMED {
			paymentTotalAmount = paymentTotalAmount + payment.Amount
		}
	}

	if paymentTotalAmount < orderAmount {
		return errors.New(global_constant.ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_PAYMENT_PENDING)
	}

	confirmItemsError := AppServices.Item.ConfirmOrderItems(order.DocId)

	if confirmItemsError != nil {
		return errors.New(global_constant.ERROR_ORDER_CANNOT_BE_CONFIRMED_DUE_TO_ITEM_CONFIRM_ISSUE)
	}

	order.StatusCode = global_constant.ORDER_CONFIRMED
	updateError := s.orderRepository.Update(context.Background(), order.DocId, order)

	event := order.ToEvent(global_constant.OPERATION_CONFIRMED)
	common.AddToChanCRUD(event)
	common.AddToChanPaymentRefund(event)

	return updateError
}
