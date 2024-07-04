package repositories

import (
	"github.com/nanda03dev/go2ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	*BaseRepository[models.User]
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[models.User](db, "users"), // "users" is the collection name
	}
}

// Additional methods specific to user repository can be added here
