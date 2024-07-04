package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/nanda03dev/oms/models"
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
