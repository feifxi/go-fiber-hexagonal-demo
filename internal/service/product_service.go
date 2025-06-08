package service

import (
	"chanombude/super-hexagonal/internal/domain"
	"chanombude/super-hexagonal/internal/repository"
)

type ProductService interface {
	CreateProduct(product *domain.Product) error
	GetProduct(id uint) (*domain.Product, error)
	GetAllProducts() ([]domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *domain.Product) error {
	// Add any business logic here before creating
	return s.repo.Create(product)
}

func (s *productService) GetProduct(id uint) (*domain.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) GetAllProducts() ([]domain.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) UpdateProduct(product *domain.Product) error {
	// Add any business logic here before updating
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
