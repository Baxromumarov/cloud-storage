package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type company struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewCompanyRepo(db *sqlx.DB, log logger.Logger) FilePostgresRepo {
	return &company{
		db:  db,
		log: log,
	}
}

func (c *company) Create(ctx context.Context, req *File) error {
	return nil
}
