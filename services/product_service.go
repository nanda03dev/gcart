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

type ProductService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetAllProducts(requestFilterBody common.RequestFilterBodyType) ([]models.Product, error)
	GetProductByID(docId string) (models.Product, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(docId string) error
}

type productService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	product.DocId = utils.Generate16DigitUUID()
	createError := s.productRepository.Create(context.Background(), product)

	event := product.ToEvent(global_constant.OPERATION_CREATE)
	common.AddToChanCRUD(event)

	return product, createError
}

func (s *productService) GetAllProducts(requestFilterBody common.RequestFilterBodyType) ([]models.Product, error) {
	return s.productRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *productService) GetProductByID(docId string) (models.Product, error) {
	return s.productRepository.GetByID(context.Background(), docId)
}

func (s *productService) UpdateProduct(updateProduct models.Product) error {
	product, getByIdError := s.productRepository.GetByID(context.Background(), updateProduct.DocId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}

	updateError := s.productRepository.Update(context.Background(), product.DocId, product.ToUpdatedDocument(updateProduct))

	event := product.ToEvent(global_constant.OPERATION_UPDATE)
	common.AddToChanCRUD(event)

	return updateError
}

func (s *productService) DeleteProduct(docId string) error {
	product, getByIdError := s.productRepository.GetByID(context.Background(), docId)

	if getByIdError != nil {
		return errors.New(global_constant.ENTITY_NOT_FOUND)
	}
	deleteError := s.productRepository.Delete(context.Background(), docId)

	event := product.ToEvent(global_constant.OPERATION_DELETE)
	common.AddToChanCRUD(event)

	return deleteError
}
