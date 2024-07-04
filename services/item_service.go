package services

import (
	"context"

	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemService interface {
	CreateItem(item models.Item) error
	GetAllCities() ([]models.Item, error)
	GetItemByID(id string) (models.Item, error)
	UpdateItem(item models.Item) error
	DeleteItem(id string) error
}

type itemService struct {
	itemRepository *repositories.ItemRepository
}

func NewItemService(itemRepository *repositories.ItemRepository) ItemService {
	return &itemService{itemRepository}
}

func (s *itemService) CreateItem(item models.Item) error {
	return s.itemRepository.Create(context.Background(), item)
}

func (s *itemService) GetAllCities() ([]models.Item, error) {
	return s.itemRepository.GetAll(context.Background(), nil)
}

func (s *itemService) GetItemByID(id string) (models.Item, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Item{}, err
	}

	return s.itemRepository.GetByID(context.Background(), objectId)
}

func (s *itemService) UpdateItem(item models.Item) error {
	return s.itemRepository.Update(context.Background(), item.ID, item)
}

func (s *itemService) DeleteItem(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.itemRepository.Delete(context.Background(), objectId)
}
