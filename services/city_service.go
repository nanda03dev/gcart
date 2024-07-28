package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
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
	return city, s.cityRepository.Create(context.Background(), city)
}

func (s *cityService) GetAllCities(requestFilterBody common.RequestFilterBodyType) ([]models.City, error) {
	return s.cityRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *cityService) GetCityByID(docId string) (models.City, error) {
	return s.cityRepository.GetByID(context.Background(), docId)
}

func (s *cityService) UpdateCity(city models.City) error {
	return s.cityRepository.Update(context.Background(), city.DocId, city)
}

func (s *cityService) DeleteCity(docId string) error {
	return s.cityRepository.Delete(context.Background(), docId)
}
