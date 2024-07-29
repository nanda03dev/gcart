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

type CityService interface {
	CreateCity(city models.City) (models.City, error)
	GetAllCities(requestFilterBody common.RequestFilterBodyType) ([]models.City, error)
	GetCityByID(docId string) (models.City, error)
	UpdateCity(city models.City) error
	DeleteCity(docId string) error
}

type cityService struct {
	cityRepository *repositories.CityRepository
}

func NewCityService(cityRepository *repositories.CityRepository) CityService {
	return &cityService{cityRepository}
}

func (s *cityService) CreateCity(city models.City) (models.City, error) {
	city.DocId = utils.Generate16DigitUUID()
	createError := s.cityRepository.Create(context.Background(), city)

	event := city.ToEvent(global_constant.OPERATION_CREATE)
	workers.AddToChanCRUD(event)

	return city, createError
}

func (s *cityService) GetAllCities(requestFilterBody common.RequestFilterBodyType) ([]models.City, error) {
	return s.cityRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *cityService) GetCityByID(docId string) (models.City, error) {
	return s.cityRepository.GetByID(context.Background(), docId)
}

func (s *cityService) UpdateCity(city models.City) error {

	updateError := s.cityRepository.Update(context.Background(), city.DocId, city)

	event := city.ToEvent(global_constant.OPERATION_UPDATE)
	workers.AddToChanCRUD(event)

	return updateError
}

func (s *cityService) DeleteCity(docId string) error {
	city, getByIdError := s.cityRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New("entity not found")
	}
	deleteError := s.cityRepository.Delete(context.Background(), docId)

	event := city.ToEvent(global_constant.OPERATION_DELETE)
	workers.AddToChanCRUD(event)

	return deleteError
}
