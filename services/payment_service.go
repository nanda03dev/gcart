package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
	"github.com/nanda03dev/go2ms/workers"
)

type PaymentService interface {
	CreatePayment(payment models.Payment) (models.Payment, error)
	GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error)
	GetPaymentByID(docId string) (models.Payment, error)
	UpdatePayment(payment models.Payment) error
	DeletePayment(docId string) error
}

type paymentService struct {
	paymentRepository *repositories.PaymentRepository
}

func NewPaymentService(paymentRepository *repositories.PaymentRepository) PaymentService {
	return &paymentService{paymentRepository}
}

func (s *paymentService) CreatePayment(payment models.Payment) (models.Payment, error) {
	payment.DocId = utils.Generate16DigitUUID()
	payment.Code = global_constant.PAYMENT_INITIATED
	createError := s.paymentRepository.Create(context.Background(), payment)

	event := payment.ToEvent(global_constant.OPERATION_CREATE)
	workers.AddToChanCRUD(event)

	return payment, createError
}

func (s *paymentService) GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error) {
	return s.paymentRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *paymentService) GetPaymentByID(docId string) (models.Payment, error) {
	return s.paymentRepository.GetByID(context.Background(), docId)
}

func (s *paymentService) UpdatePayment(payment models.Payment) error {
	updateError := s.paymentRepository.Update(context.Background(), payment.DocId, payment)

	event := payment.ToEvent(global_constant.OPERATION_UPDATE)
	workers.AddToChanCRUD(event)

	return updateError
}

func (s *paymentService) DeletePayment(docId string) error {
	payment, getByIdError := s.paymentRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New("entity not found")
	}
	deleteError := s.paymentRepository.Delete(context.Background(), docId)

	event := payment.ToEvent(global_constant.OPERATION_DELETE)
	workers.AddToChanCRUD(event)

	return deleteError
}
