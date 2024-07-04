package repositories

import (
	"github.com/nanda03dev/go2ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	*BaseRepository[models.Product]
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		BaseRepository: NewBaseRepository[models.Product](db, "products"), // "products" is the collection name
	}
}

// Additional methods specific to product repository can be added here
