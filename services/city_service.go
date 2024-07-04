package services

import (
	"context"

	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CityService interface {
	CreateCity(city models.City) (models.City, error)
	GetAllCities() ([]models.City, error)
	GetCityByID(id string) (models.City, error)
	UpdateCity(city models.City) error
	DeleteCity(id string) error
}

type cityService struct {
	cityRepository *repositories.CityRepository
}

func NewCityService(cityRepository *repositories.CityRepository) CityService {
	return &cityService{cityRepository}
}

func (s *cityService) CreateCity(city models.City) (models.City, error) {
	city.ID = primitive.NewObjectID()
	return city, s.cityRepository.Create(context.Background(), city)
}

func (s *cityService) GetAllCities() ([]models.City, error) {
	return s.cityRepository.GetAll(context.Background(), nil)
}

func (s *cityService) GetCityByID(id string) (models.City, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.City{}, err
	}

	return s.cityRepository.GetByID(context.Background(), objectId)
}

func (s *cityService) UpdateCity(city models.City) error {
	return s.cityRepository.Update(context.Background(), city.ID, city)
}

func (s *cityService) DeleteCity(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.cityRepository.Delete(context.Background(), objectId)
}
