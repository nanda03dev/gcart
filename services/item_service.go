package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
)

type ItemService interface {
	CreateItem(item models.Item) (models.Item, error)
	GetAllItems(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error)
	GetItemByID(docId string) (models.Item, error)
	UpdateItem(item models.Item) error
	UpdateItemsTimeout(requestFilterBody common.FiltersBodyType) error
	DeleteItem(docId string) error
	DeleteOrderItems(orderId string) error
	ConfirmOrderItems(orderId string) error
}

type itemService struct {
	itemRepository *repositories.ItemRepository
}

func NewItemService(itemRepository *repositories.ItemRepository) ItemService {
	return &itemService{itemRepository}
}

func (s *itemService) CreateItem(item models.Item) (models.Item, error) {
	item.DocId = models.Generate16DigitUUID()
	item.StatusCode = global_constant.ITEM_INITIATED
	createError := s.itemRepository.Create(context.Background(), item)

	order, _ := AppServices.Order.GetOrderByID(item.OrderId)
	order.Amount = order.Amount + item.Amount
	AppServices.Order.UpdateOrder(order)

	event := item.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return item, createError
}

func (s *itemService) GetAllItems(requestFilterBody common.RequestFilterBodyType) ([]models.Item, error) {
	return s.itemRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *itemService) GetAllItemsByOrderId(orderId string) ([]models.Item, error) {
	return s.itemRepository.GetAllItemsByOrderId(orderId)
}

func (s *itemService) GetItemByID(docId string) (models.Item, error) {
	return s.itemRepository.GetByID(context.Background(), docId)
}

func (s *itemService) UpdateItem(updateItem models.Item) error {
	item, getByIdError := s.itemRepository.GetByID(context.Background(), updateItem.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}

	updateError := s.itemRepository.Update(context.Background(), item.DocId, item.ToUpdatedDocument(updateItem))

	event := item.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *itemService) UpdateItemsTimeout(filter common.FiltersBodyType) error {
	items, _ := s.itemRepository.GetAll(context.Background(), filter, nil, nil)

	updateDocument := map[string]interface{}{
		"statusCode": global_constant.ITEM_TIMEOUT,
	}

	updateManyError := s.itemRepository.UpdateMany(context.TODO(), filter, updateDocument)

	for _, item := range items {
		event := item.ToEvent(global_constant.OPERATION_UPDATE)
		common.AddToChanCRUD(event)
	}
	return updateManyError
}

func (s *itemService) DeleteItem(docId string) error {
	item, getByIdError := s.itemRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ERROR_ENTITY_NOT_FOUND)
	}
	deleteError := s.itemRepository.Delete(context.Background(), docId)

	event := item.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}

func (s *itemService) DeleteOrderItems(orderId string) error {

	orderIdFilter := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}
	items, _ := s.itemRepository.GetAllItemsByOrderId(orderId)

	deleteError := s.itemRepository.DeleteMany(context.Background(), orderIdFilter)

	for _, item := range items {
		event := item.ToEvent(global_constant.OPERATION_DELETE)
		common.AddToChanCRUD(event)
	}

	return deleteError
}

func (s *itemService) ConfirmOrderItems(orderId string) error {

	orderIdFilter := common.FiltersBodyType{
		{Key: "orderId", Value: orderId},
	}

	items, _ := s.itemRepository.GetAllItemsByOrderId(orderId)

	updateDocument := map[string]string{
		"statusCode": string(global_constant.ITEM_CONFIRMED),
	}

	updateError := s.itemRepository.UpdateMany(context.Background(), orderIdFilter, updateDocument)

	if updateError == nil {
		for _, item := range items {
			event := item.ToEvent(global_constant.OPERATION_CONFIRMED)
			common.AddToChanCRUD(event)
		}
	}

	return updateError
}
