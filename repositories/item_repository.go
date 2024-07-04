package repositories

import (
	"github.com/nanda03dev/go2ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	*BaseRepository[models.Item]
}

func NewItemRepository(db *mongo.Database) *ItemRepository {
	return &ItemRepository{
		BaseRepository: NewBaseRepository[models.Item](db, "cities"), // "items" is the collection name
	}
}

// Additional methods specific to item repository can be added here
