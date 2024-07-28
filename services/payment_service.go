package services

import (
	"context"

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
	GetPaymentByID(id string) (models.Payment, error)
	UpdatePayment(payment models.Payment) error
	DeletePayment(id string) error
}

type paymentService struct {
	paymentRepository *repositories.PaymentRepository
}

func NewPaymentService(paymentRepository *repositories.PaymentRepository) PaymentService {
	return &paymentService{paymentRepository}
}

func (s *paymentService) CreatePayment(payment models.Payment) (models.Payment, error) {
	payment.DocId = utils.Generate16DigitUUID()

	event := common.EventType{
		EntityId:      payment.DocId,
		EntityType:    global_constant.Entities.Payment,
		OperationType: global_constant.Operations.Create,
	}
	workers.AddToChanCRUD(event)

	return payment, s.paymentRepository.Create(context.Background(), payment)
}

func (s *paymentService) GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error) {
	return s.paymentRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *paymentService) GetPaymentByID(docId string) (models.Payment, error) {
	return s.paymentRepository.GetByID(context.Background(), docId)
}

func (s *paymentService) UpdatePayment(payment models.Payment) error {
	event := common.EventType{
		EntityId:      payment.DocId,
		EntityType:    global_constant.Entities.Payment,
		OperationType: global_constant.Operations.Update,
	}
	workers.AddToChanCRUD(event)

	return s.paymentRepository.Update(context.Background(), payment.DocId, payment)
}

func (s *paymentService) DeletePayment(docId string) error {
	event := common.EventType{
		EntityId:      docId,
		EntityType:    global_constant.Entities.Payment,
		OperationType: global_constant.Operations.Delete,
	}
	workers.AddToChanCRUD(event)
	return s.paymentRepository.Delete(context.Background(), docId)
}
