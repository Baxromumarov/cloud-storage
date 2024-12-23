package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type filePostgres struct {
	db  *sqlx.DB
	log logger.Logger
}

func NewFilePostgresRepo(db *sqlx.DB, log logger.Logger) FilePostgresRepo {
	return &filePostgres{
		db:  db,
		log: log,
	}
}

func (c *filePostgres) Create(ctx context.Context, req *File) error {
	query := `INSERT INTO files (
					name, 
					size, 
					path, 
					bucket, 
					content_type,
					metadata,
					md5_hash,
					created_at,
					updated_at,
					deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := c.db.ExecContext(ctx, query, req.Name, req.Size, req.Path, req.Bucket, req.ContentType, req.Metadata, req.MD5Hash, req.CreatedAt, req.UpdatedAt, req.DeletedAt)
	if err != nil {
		return c.log.Error("Error while inserting file into database", logger.String("error", err.Error()))
	}
	return nil
}
