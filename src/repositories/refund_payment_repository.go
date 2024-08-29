package repositories

import (
	"context"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefundPaymentRepository struct {
	*BaseRepository[models.RefundPayment]
}

func NewRefundPaymentRepository(db *mongo.Database) *RefundPaymentRepository {
	return &RefundPaymentRepository{
		BaseRepository: NewBaseRepository[models.RefundPayment](db, "refundPayments"), // "refundPayments" is the collection name
	}
}

// Additional methods specific to refundPayment repository can be added here

func (r *RefundPaymentRepository) GetAllRefundPaymentsByOrderId(orderId string) ([]models.RefundPayment, error) {
	filters := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	results, getAllError := r.GetAll(context.Background(), filters, nil, nil)

	return results, getAllError
}
