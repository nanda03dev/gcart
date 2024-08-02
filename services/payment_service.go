package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
)

type PaymentService interface {
	CreatePayment(payment models.Payment) (models.Payment, error)
	GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error)
	GetAllPaymentsByOrderId(orderId string) ([]models.Payment, error)
	GetPaymentByID(docId string) (models.Payment, error)
	UpdatePayment(payment models.Payment) error
	UpdatePaymentTimeout(docId string) bool
	DeletePayment(docId string) error
	DeleteOrderPayments(orderId string) error
	ConfirmPayment(paymentConfirmBody common.PaymentConfirmBody) error
}

type paymentService struct {
	paymentRepository *repositories.PaymentRepository
}

func NewPaymentService(paymentRepository *repositories.PaymentRepository) PaymentService {
	return &paymentService{paymentRepository}
}

func (s *paymentService) CreatePayment(payment models.Payment) (models.Payment, error) {
	payment.DocId = models.Generate16DigitUUID()
	payment.StatusCode = global_constant.PAYMENT_INITIATED
	createError := s.paymentRepository.Create(context.Background(), payment)

	event := payment.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return payment, createError
}

func (s *paymentService) GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error) {
	return s.paymentRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *paymentService) GetAllPaymentsByOrderId(orderId string) ([]models.Payment, error) {
	return s.paymentRepository.GetAllPaymentsByOrderId(orderId)
}

func (s *paymentService) GetPaymentByID(docId string) (models.Payment, error) {
	return s.paymentRepository.GetByID(context.Background(), docId)
}

func (s *paymentService) UpdatePayment(updatePayment models.Payment) error {
	payment, getByIdError := s.paymentRepository.GetByID(context.Background(), updatePayment.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}

	updateError := s.paymentRepository.Update(context.Background(), payment.DocId, payment.ToUpdatedDocument(updatePayment))

	event := payment.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *paymentService) UpdatePaymentTimeout(docId string) bool {
	payment, _ := s.paymentRepository.GetByID(context.TODO(), docId)

	if payment.StatusCode == global_constant.PAYMENT_INITIATED {
		payment.StatusCode = global_constant.PAYMENT_TIMEOUT
		s.paymentRepository.Update(context.TODO(), payment.DocId, payment)

		event := payment.ToEvent(global_constant.OPERATION_UPDATE)
		common.AddToChanCRUD(event)
		return true
	}

	return false
}

func (s *paymentService) DeletePayment(docId string) error {
	payment, getByIdError := s.paymentRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}
	deleteError := s.paymentRepository.Delete(context.Background(), docId)

	event := payment.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}

func (s *paymentService) DeleteOrderPayments(orderId string) error {
	orderIdFilter := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	payments, _ := s.paymentRepository.GetAllPaymentsByOrderId(orderId)

	deleteError := s.paymentRepository.DeleteMany(context.Background(), orderIdFilter)

	for _, payment := range payments {
		event := payment.ToEvent(global_constant.OPERATION_DELETE)
		common.AddToChanCRUD(event)
	}
	return deleteError
}

func (s *paymentService) ConfirmPayment(paymentConfirmBody common.PaymentConfirmBody) error {
	payment, getByError := s.paymentRepository.GetByID(context.TODO(), paymentConfirmBody.PaymentId)

	if getByError == nil && payment.StatusCode == global_constant.PAYMENT_INITIATED {

		payment.StatusCode = global_constant.PAYMENT_CONFIRMED
		confirmError := s.paymentRepository.Update(context.TODO(), payment.DocId, payment)

		event := payment.ToEvent(global_constant.OPERATION_CONFIRMED)
		common.AddToChanCRUD(event)
		return confirmError
	}

	return getByError
}
