package repositories

import (
	"context"

	"github.com/nanda03dev/gcart/common"
	"github.com/nanda03dev/gcart/models"
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

func (r *ItemRepository) GetAllItemsByOrderId(orderId string) ([]models.Item, error) {
	filters := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	results, getAllError := r.GetAll(context.Background(), filters, nil, nil)

	return results, getAllError
}
