package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/config"
)

func New(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, func(), error) {

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
	dbpool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	return dbpool, func() { dbpool.Close() }, nil
}
