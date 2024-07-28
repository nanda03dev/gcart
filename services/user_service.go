package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
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
	user.DocId = utils.Generate16DigitUUID()

	return user, s.userRepository.Create(context.Background(), user)
}

func (s *userService) GetAllUsers(requestFilterBody common.RequestFilterBodyType) ([]models.User, error) {
	return s.userRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *userService) GetUserByID(docId string) (models.User, error) {
	return s.userRepository.GetByID(context.Background(), docId)
}

func (s *userService) UpdateUser(user models.User) error {
	return s.userRepository.Update(context.Background(), user.DocId, user)
}

func (s *userService) DeleteUser(docId string) error {
	return s.userRepository.Delete(context.Background(), docId)
}
