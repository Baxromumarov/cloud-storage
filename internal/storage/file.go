package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
)

type File struct {
	log logger.Logger

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

func NewFileRepo(log logger.Logger) FileRepo {
	return &File{
		log: log,
	}
}

func (f *File) Create(ctx context.Context, file *File) error {
	return nil
}

