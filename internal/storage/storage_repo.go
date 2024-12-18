package storage

import (
	"context"

	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

// register all repositories

type StorageI interface {
	File() FileRepo
	Postgres() FilePostgresRepo
}

type storage struct {
	file     FileRepo
	postgres FilePostgresRepo
}

func NewStorage(db *sqlx.DB, log logger.Logger) StorageI {
	return &storage{
		file: NewFileRepo(db, log),
	}
}

func (s *storage) File() FileRepo {
	return s.file
}
func (s *storage) Postgres() FilePostgresRepo {
	return s.postgres
}

// All file operations should be defined here
type FileRepo interface {
	Create(ctx context.Context, req *File) error
}

// FilePostgresRepo is a repository for file operations in postgres
type FilePostgresRepo interface {
	Create(ctx context.Context, req *File) error
}

// LogPostgresRepo is a repository for log operations in postgres
type LogPostgresRepo interface{}

// compile time check if the interfaces are implemented
var _ FileRepo = &File{}
var _ FilePostgresRepo = &File{}
var _ FilePostgresRepo = &File{}
