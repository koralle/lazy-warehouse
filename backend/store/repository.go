package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/koralle/lazy-warehouse/backend/config"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=%s search_path=%s",
		cfg.DBHost, cfg.DBPort,
		cfg.DBUser, cfg.DBPassword,
		cfg.DBName,
		cfg.SSLMode,
		cfg.SearchPath,
	)
	log.Println(dsn)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	return db, func() { _ = db.Close() }, nil
}
