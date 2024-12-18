package db

import (
	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// register all repositories

type StorageI interface {
	// Company() CompanyRepo

}

type storagePg struct {
	// company CompanyRepo

}

func NewStorage(db *sqlx.DB, log logger.Logger) StorageI {
	return &storagePg{
		// company: NewCompanyRepo(db, log),
	}
}

// func (s *storagePg) Company() CompanyRepo {
// 	return s.company
// }



type CompanyRepo interface {
	// Create(ctx context.Context, req *models.Company) error
	// Update(ctx context.Context, req *models.Company) error
	// Delete(ctx context.Context, id string) error
	// Get(ctx context.Context, id string) (*models.Company, error)
}
