package repositories

import (
	"github.com/nanda03dev/gcart/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	*BaseRepository[models.Order]
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		BaseRepository: NewBaseRepository[models.Order](db, "orders"), // "orders" is the collection name
	}
}

// Additional methods specific to order repository can be added here
