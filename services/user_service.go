package services

import (
	"github.com/nanda03dev/oms/models"
	"github.com/nanda03dev/oms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id primitive.ObjectID) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id primitive.ObjectID) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(user models.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) GetUserByID(id primitive.ObjectID) (models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.userRepository.Update(user)
}

func (s *userService) DeleteUser(id primitive.ObjectID) error {
	return s.userRepository.Delete(id)
}
