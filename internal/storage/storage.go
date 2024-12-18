package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
)

// register all repositories

type StorageI interface {
	File() FileRepo
}

type storage struct {
	file FileRepo
}

func NewStorage(log logger.Logger) StorageI {
	return &storage{
		file: NewFileRepo(log),
	}
}

func (s *storage) File() FileRepo {
	return s.file
}

type FileRepo interface {
	Create(ctx context.Context, req *File) error
	// Delete(ctx context.Context, id string) error
	// Get(ctx context.Context, id string) (*File, error)
	// Update(ctx context.Context, req *File) error
}

var _ FileRepo = &File{}
