package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventService interface {
	CreateEvent(event models.Event) (models.Event, error)
	GetAllEvents(requestFilterBody common.RequestFilterBody) ([]models.Event, error)
	GetEventByID(id string) (models.Event, error)
	UpdateEvent(event models.Event) error
	DeleteEvent(id string) error
}

type eventService struct {
	eventRepository *repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) EventService {
	return &eventService{eventRepository}
}

func (s *eventService) CreateEvent(event models.Event) (models.Event, error) {
	event.ID = primitive.NewObjectID()
	return event, s.eventRepository.Create(context.Background(), event)
}

func (s *eventService) GetAllEvents(requestFilterBody common.RequestFilterBody) ([]models.Event, error) {
	return s.eventRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *eventService) GetEventByID(id string) (models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Event{}, err
	}

	return s.eventRepository.GetByID(context.Background(), objectId)
}

func (s *eventService) UpdateEvent(event models.Event) error {
	return s.eventRepository.Update(context.Background(), event.ID, event)
}

func (s *eventService) DeleteEvent(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.eventRepository.Delete(context.Background(), objectId)
}
