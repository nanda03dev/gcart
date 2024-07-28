package services

import (
	"context"

	"github.com/nanda03dev/go2ms/common"
	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"github.com/nanda03dev/go2ms/utils"
)

type ProductService interface {
	CreateProduct(product models.Product) (models.Product, error)
	GetAllProducts(requestFilterBody common.RequestFilterBodyType) ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
	UpdateProduct(product models.Product) error
	DeleteProduct(id string) error
}

type productService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	product.DocId = utils.Generate16DigitUUID()

	return product, s.productRepository.Create(context.Background(), product)
}

func (s *productService) GetAllProducts(requestFilterBody common.RequestFilterBodyType) ([]models.Product, error) {
	return s.productRepository.GetAll(context.Background(), requestFilterBody.ListOfFilter, requestFilterBody.SortBody, requestFilterBody.Size)
}

func (s *productService) GetProductByID(docId string) (models.Product, error) {
	return s.productRepository.GetByID(context.Background(), docId)
}

func (s *productService) UpdateProduct(product models.Product) error {
	return s.productRepository.Update(context.Background(), product.DocId, product)
}

func (s *productService) DeleteProduct(docId string) error {
	return s.productRepository.Delete(context.Background(), docId)
}
