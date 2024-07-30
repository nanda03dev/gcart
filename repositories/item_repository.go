package repositories

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	*BaseRepository[models.Item]
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{
		BaseRepository: NewBaseRepository[models.Item](db, "items"), // "items" is the collection name
	}
}

// Additional methods specific to item repository can be added here

func (r *ItemRepository) GetAllItemsByOrderId(orderId string) []models.Item {
	filters := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	results, _ := r.GetAll(context.Background(), filters, nil, nil)

	return results
}
