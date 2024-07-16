package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers(requestFilterBody common.RequestFilterBodyType) ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id string) error
}

type userService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	user.ID = primitive.NewObjectID()

	return user, s.userRepository.Create(context.Background(), user)
}

func (s *userService) GetAllUsers(requestFilterBody common.RequestFilterBodyType) ([]models.User, error) {
	return s.userRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *userService) GetUserByID(id string) (models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	return s.userRepository.GetByID(context.Background(), objectId)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.userRepository.Update(context.Background(), user.ID, user)
}

func (s *userService) DeleteUser(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.userRepository.Delete(context.Background(), objectId)
}
