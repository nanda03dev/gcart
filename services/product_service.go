package services

import (
	"context"

	"github.com/nanda03dev/go2ms/models"
	"github.com/nanda03dev/go2ms/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService interface {
	CreateProduct(product models.Product) error
	GetAllProducts() ([]models.Product, error)
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

func (s *productService) CreateProduct(product models.Product) error {
	return s.productRepository.Create(context.Background(), product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.productRepository.GetAll(context.Background(), nil)
}

func (s *productService) GetProductByID(id string) (models.Product, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Product{}, err
	}

	return s.productRepository.GetByID(context.Background(), objectId)
}

func (s *productService) UpdateProduct(product models.Product) error {
	return s.productRepository.Update(context.Background(), product.ID, product)
}

func (s *productService) DeleteProduct(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return s.productRepository.Delete(context.Background(), objectId)
}
