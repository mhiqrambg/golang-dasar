package service

import (
	"time"

	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/repository"
)

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetSalesReportToday() (*model.SalesReport, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	return s.repo.GetSalesReport(startOfDay, endOfDay)
}

func (s *ReportService) GetSalesReportByDateRange(startDate, endDate time.Time) (*model.SalesReport, error) {
	endDate = endDate.Add(24 * time.Hour)
	return s.repo.GetSalesReport(startDate, endDate)
}
