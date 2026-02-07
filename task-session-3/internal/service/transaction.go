package service

import (
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
