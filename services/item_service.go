package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
	"github.com/nanda03dev/go2ms/workers"
)

type ItemService interface {
	CreateItem(item models.Item) (models.Item, error)
	GetAllItems(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error)
	GetItemByID(docId string) (models.Item, error)
	UpdateItem(item models.Item) error
	DeleteItem(docId string) error
}

type itemService struct {
	itemRepository *repositories.ItemRepository
}

func NewItemService(itemRepository *repositories.ItemRepository) ItemService {
	return &itemService{itemRepository}
}

func (s *itemService) CreateItem(item models.Item) (models.Item, error) {
	item.DocId = utils.Generate16DigitUUID()
	createError := s.itemRepository.Create(context.Background(), item)

	event := item.ToEvent(global_constant.OPERATION_CREATE)
	workers.AddToChanCRUD(event)

	return item, createError
}

func (s *itemService) GetAllItems(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error) {
	return s.itemRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *itemService) GetItemByID(docId string) (models.Item, error) {
	return s.itemRepository.GetByID(context.Background(), docId)
}

func (s *itemService) UpdateItem(item models.Item) error {
	updateError := s.itemRepository.Update(context.Background(), item.DocId, item)

	event := item.ToEvent(global_constant.OPERATION_UPDATE)
	workers.AddToChanCRUD(event)

	return updateError
}

func (s *itemService) DeleteItem(docId string) error {
	item, getByIdError := s.itemRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New("entity not found")
	}
	deleteError := s.itemRepository.Delete(context.Background(), docId)

	event := item.ToEvent(global_constant.OPERATION_DELETE)
	workers.AddToChanCRUD(event)

	return deleteError
}
