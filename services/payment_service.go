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
	payment.ID = primitive.NewObjectID()

	event := common.EventType{
		EntityId:      payment.ID,
		EntityType:    global_constant.Entities.Payment,
		OperationType: global_constant.Operations.Create,
	}
	workers.AddToChanCRUD(event)

	return payment, s.paymentRepository.Create(context.Background(), payment)
}

func (s *paymentService) GetAllPayments(requestFilterBody common.RequestFilterBodyType) ([]models.Payment, error) {
	return s.paymentRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *paymentService) GetPaymentByID(id string) (models.Payment, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Payment{}, err
	}

	return s.paymentRepository.GetByID(context.Background(), objectId)
}

func (s *paymentService) UpdatePayment(payment models.Payment) error {
	return s.paymentRepository.Update(context.Background(), payment.ID, payment)
}

func (s *paymentService) DeletePayment(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.paymentRepository.Delete(context.Background(), objectId)
}
