package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type File struct {
	log logger.Logger
	db  *sqlx.DB

	Name        string                 `json:"name"`
	Size        int64                  `json:"size"`
	Path        string                 `json:"path"`
	Bucket      string                 `json:"bucket"`
	ContentType string                 `json:"content_type"`
	Metadata    map[string]interface{} `json:"metadata"`
	MD5Hash     string                 `json:"md5_hash"` // md5 hash of the file
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	DeletedAt   string                 `json:"deleted_at"`
}

func NewFileRepo(db *sqlx.DB, log logger.Logger) *File {
	return &File{
		log: log,
		db:  db,
	}
}

// Save file in postgres
func (f *File) Create(ctx context.Context, file *File) error {
	
	return nil
}
