package v1

import (
	"github.com/baxromumarov/cloud-storage/config"
	"github.com/baxromumarov/cloud-storage/internal/db"
	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
	// "github.com/jmoiron/sqlx"
)

type handlerV1 struct {
	log             logger.Logger
	cfg             *config.Config
	storagePostgres db.PostgresI
}

type HandlerV1Options struct {
	Log logger.Logger
	Cfg *config.Config
	Db  *sqlx.DB
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:             options.Log,
		cfg:             options.Cfg,
		storagePostgres: db.NewStorage(options.Db, options.Log),
	}
}

func (h *handlerV1) Log() logger.Logger {
	return h.log
}

func (h *handlerV1) Config() *config.Config {
	return h.cfg
}
