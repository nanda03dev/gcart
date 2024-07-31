package services

import (
	"context"
	"errors"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/global_constant"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers(requestFilterBody common.RequestFilterBodyType) ([]models.User, error)
	GetUserByID(docId string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(docId string) error
}

type userService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	user.DocId = utils.Generate16DigitUUID()
	createError := s.userRepository.Create(context.Background(), user)

	event := user.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return user, createError
}

func (s *userService) GetAllUsers(requestFilterBody common.RequestFilterBodyType) ([]models.User, error) {
	return s.userRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *userService) GetUserByID(docId string) (models.User, error) {
	return s.userRepository.GetByID(context.Background(), docId)
}

func (s *userService) UpdateUser(updateUser models.User) error {
	user, getByIdError := s.userRepository.GetByID(context.Background(), updateUser.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}

	updateError := s.userRepository.Update(context.Background(), user.DocId, user.ToUpdatedDocument(updateUser))

	event := user.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *userService) DeleteUser(docId string) error {
	user, getByIdError := s.userRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}
	deleteError := s.userRepository.Delete(context.Background(), docId)

	event := user.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}
