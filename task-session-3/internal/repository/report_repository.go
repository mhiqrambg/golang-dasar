package repository

import (
	"database/sql"
	"time"

	"github.com/mhiqrambg/golang-dasar/task-session-3/internal/model"
)

type ReportRepository interface {
	GetSalesReport(startDate, endDate time.Time) (*model.SalesReport, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetSalesReport(startDate, endDate time.Time) (*model.SalesReport, error) {
	report := &model.SalesReport{}

	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0), COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`
		SELECT COALESCE(p.name, 'N/A'), COALESCE(SUM(td.quantity), 0)
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id::uuid = p.id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.name
		ORDER BY SUM(td.quantity) DESC
		LIMIT 1
	`, startDate, endDate).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err == sql.ErrNoRows {
		report.ProdukTerlaris = model.ProdukTerlaris{Nama: "N/A", QtyTerjual: 0}
	} else if err != nil {
		return nil, err
	}

	return report, nil
}
