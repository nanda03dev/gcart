package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
)

type EventService interface {
	CreateEvent(event models.Event) (models.Event, error)
	GetAllEvents(requestFilterBody common.RequestFilterBodyType) ([]models.Event, error)
	GetEventByID(docId string) (models.Event, error)
	UpdateEvent(event models.Event) error
	DeleteEvent(docId string) error
}

type eventService struct {
	eventRepository *repositories.EventRepository
}

func NewEventService(eventRepository *repositories.EventRepository) EventService {
	return &eventService{eventRepository}
}

func (s *eventService) CreateEvent(event models.Event) (models.Event, error) {
	event.DocId = models.Generate16DigitUUID()
	return event, s.eventRepository.Create(context.Background(), event)
}

func (s *eventService) GetAllEvents(requestFilterBody common.RequestFilterBodyType) ([]models.Event, error) {
	return s.eventRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *eventService) GetEventByID(docId string) (models.Event, error) {
	return s.eventRepository.GetByID(context.Background(), docId)
}

func (s *eventService) UpdateEvent(event models.Event) error {
	return s.eventRepository.Update(context.Background(), event.DocId, event)
}

func (s *eventService) DeleteEvent(docId string) error {
	return s.eventRepository.Delete(context.Background(), docId)
}
