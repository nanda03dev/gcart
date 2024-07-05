package repositories

import (
	"github.com/nanda03dev/go2ms/models"
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
