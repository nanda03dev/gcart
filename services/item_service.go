package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
)

type ItemService interface {
	CreateItem(item models.Item) (models.Item, error)
	GetAllCities(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error)
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

func (s *itemService) CreateItem(item models.Item) (models.Item, error) {
	item.DocId = utils.Generate16DigitUUID()
	return item, s.itemRepository.Create(context.Background(), item)
}

func (s *itemService) GetAllCities(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error) {
	return s.itemRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *itemService) GetItemByID(docId string) (models.Item, error) {

	return s.itemRepository.GetByID(context.Background(), docId)
}

func (s *itemService) UpdateItem(item models.Item) error {
	return s.itemRepository.Update(context.Background(), item.DocId, item)
}

func (s *itemService) DeleteItem(docId string) error {
	return s.itemRepository.Delete(context.Background(), docId)
}
