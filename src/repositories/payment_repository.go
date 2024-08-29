package repositories

import (
	"context"

	"github.com/nanda03dev/gcart/src/common"
	"github.com/nanda03dev/gcart/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	*BaseRepository[models.Payment]
}

func NewPaymentRepository(db *mongo.Database) *PaymentRepository {
	return &PaymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db, "payments"), // "payments" is the collection name
	}
}

// Additional methods specific to payment repository can be added here

func (r *PaymentRepository) GetAllPaymentsByOrderId(orderId string) ([]models.Payment, error) {
	filters := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	results, getAllError := r.GetAll(context.Background(), filters, nil, nil)

	return results, getAllError
}
