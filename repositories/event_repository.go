package repositories

import (
	"github.com/nanda03dev/go2ms/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	*BaseRepository[models.Event]
}

func NewEventRepository(db *mongo.Database) *EventRepository {
	return &EventRepository{
		BaseRepository: NewBaseRepository[models.Event](db, "events"),
	}
}

// Additional methods specific to Event repository can be added here
