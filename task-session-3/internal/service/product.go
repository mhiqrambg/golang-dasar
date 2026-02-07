package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]model.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) GetByID(id uuid.UUID) (model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Create(product *model.Product) error {
	product.ID = uuid.New()
	product.CreatedAt = time.Now()
	_, err := s.repo.Create(*product)
	return err
}

func (s *ProductService) Update(product *model.Product) error {
	fmt.Printf("DEBUG - Incoming product: Name=%s, Price=%d, Stock=%d\n",
		product.Name, product.Price, product.Stock)

	existing, err := s.repo.GetByID(product.ID)
	if err != nil {
		return err
	}

	fmt.Printf("DEBUG - Existing product: Name=%s, Price=%d, Stock=%d\n",
		existing.Name, existing.Price, existing.Stock)

	if product.Name != "" {
		existing.Name = product.Name
	}
	if product.Price != 0 {
		existing.Price = product.Price
	}
	if product.Stock != 0 {
		existing.Stock = product.Stock
	}

	fmt.Printf("DEBUG - After merge: Name=%s, Price=%d, Stock=%d\n",
		existing.Name, existing.Price, existing.Stock)

	updated, err := s.repo.Update(existing)
	if err != nil {
		return err
	}

	*product = updated

	fmt.Printf("DEBUG - Final product: Name=%s, Price=%d, Stock=%d\n",
		product.Name, product.Price, product.Stock)

	return nil
}

func (s *ProductService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
