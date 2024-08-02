package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
)

type RefundPaymentService interface {
	CreateRefundPayment(refundPayment models.RefundPayment) (models.RefundPayment, error)
	GetAllRefundPayments(requestFilterBody common.RequestFilterBodyType) ([]models.RefundPayment, error)
	GetAllRefundPaymentsByOrderId(orderId string) ([]models.RefundPayment, error)
	GetRefundPaymentByID(docId string) (models.RefundPayment, error)
	UpdateRefundPayment(refundPayment models.RefundPayment) error
	UpdateRefundPaymentTimeout(docId string) bool
	DeleteRefundPayment(docId string) error
	DeleteOrderRefundPayments(orderId string) error
	ConfirmRefundPayment(refundPaymentConfirmBody common.RefundPaymentConfirmBody) error
}

type refundPaymentService struct {
	refundPaymentRepository *repositories.RefundPaymentRepository
}

func NewRefundPaymentService(refundPaymentRepository *repositories.RefundPaymentRepository) RefundPaymentService {
	return &refundPaymentService{refundPaymentRepository}
}

func (s *refundPaymentService) CreateRefundPayment(refundPayment models.RefundPayment) (models.RefundPayment, error) {
	refundPayment.DocId = models.Generate16DigitUUID()
	refundPayment.StatusCode = global_constant.REFUND_PAYMENT_INITIATED
	createError := s.refundPaymentRepository.Create(context.Background(), refundPayment)

	event := refundPayment.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return refundPayment, createError
}

func (s *refundPaymentService) GetAllRefundPayments(requestFilterBody common.RequestFilterBodyType) ([]models.RefundPayment, error) {
	return s.refundPaymentRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *refundPaymentService) GetAllRefundPaymentsByOrderId(orderId string) ([]models.RefundPayment, error) {
	return s.refundPaymentRepository.GetAllRefundPaymentsByOrderId(orderId)
}

func (s *refundPaymentService) GetRefundPaymentByID(docId string) (models.RefundPayment, error) {
	return s.refundPaymentRepository.GetByID(context.Background(), docId)
}

func (s *refundPaymentService) UpdateRefundPayment(updateRefundPayment models.RefundPayment) error {
	refundPayment, getByIdError := s.refundPaymentRepository.GetByID(context.Background(), updateRefundPayment.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}

	updateError := s.refundPaymentRepository.Update(context.Background(), refundPayment.DocId, refundPayment.ToUpdatedDocument(updateRefundPayment))

	event := refundPayment.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *refundPaymentService) UpdateRefundPaymentTimeout(docId string) bool {
	refundPayment, _ := s.refundPaymentRepository.GetByID(context.TODO(), docId)

	if refundPayment.StatusCode == global_constant.REFUND_PAYMENT_INITIATED {
		refundPayment.StatusCode = global_constant.REFUND_PAYMENT_TIMEOUT
		s.refundPaymentRepository.Update(context.TODO(), refundPayment.DocId, refundPayment)

		event := refundPayment.ToEvent(global_constant.OPERATION_UPDATE)
		common.AddToChanCRUD(event)
		return true
	}

	return false
}

func (s *refundPaymentService) DeleteRefundPayment(docId string) error {
	refundPayment, getByIdError := s.refundPaymentRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}
	deleteError := s.refundPaymentRepository.Delete(context.Background(), docId)

	event := refundPayment.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}

func (s *refundPaymentService) DeleteOrderRefundPayments(orderId string) error {
	orderIdFilter := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	refundPayments, _ := s.refundPaymentRepository.GetAllRefundPaymentsByOrderId(orderId)

	deleteError := s.refundPaymentRepository.DeleteMany(context.Background(), orderIdFilter)

	for _, refundPayment := range refundPayments {
		event := refundPayment.ToEvent(global_constant.OPERATION_DELETE)
		common.AddToChanCRUD(event)
	}
	return deleteError
}

func (s *refundPaymentService) ConfirmRefundPayment(refundPaymentConfirmBody common.RefundPaymentConfirmBody) error {
	refundPayment, getByError := s.refundPaymentRepository.GetByID(context.TODO(), refundPaymentConfirmBody.RefundPaymentId)

	if getByError == nil && refundPayment.StatusCode == global_constant.REFUND_PAYMENT_INITIATED {

		refundPayment.StatusCode = global_constant.REFUND_PAYMENT_CONFIRMED
		confirmError := s.refundPaymentRepository.Update(context.TODO(), refundPayment.DocId, refundPayment)

		event := refundPayment.ToEvent(global_constant.OPERATION_CONFIRMED)
		common.AddToChanCRUD(event)
		return confirmError
	}

	return getByError
}
