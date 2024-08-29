package repositories

import (
	"github.com/nanda03dev/gcart/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CityRepository struct {
	*BaseRepository[models.City]
}

func NewCityRepository(db *mongo.Database) *CityRepository {
	return &CityRepository{
		BaseRepository: NewBaseRepository[models.City](db, "cities"), // "citys" is the collection name
	}
}

// Additional methods specific to city repository can be added here
