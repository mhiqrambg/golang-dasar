package repository

import (
	"github.com/google/uuid"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
)

type ProductRepository interface {
	Create(product model.Product) (model.Product, error)
	GetByID(id uuid.UUID) (model.Product, error)
	GetAll(name string) ([]model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(id uuid.UUID) error
}
