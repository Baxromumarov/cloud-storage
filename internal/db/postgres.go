package db

import (
	"fmt"

	"github.com/baxromumarov/cloud-storage/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(cfg *config.Config) (*sqlx.DB, error) {

	// connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgresDatabase,
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// )

	db, err := sqlx.Connect("postgres", cfg.PostgresUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to Postgres")
	return db, nil

}
