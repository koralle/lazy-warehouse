package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type userService struct {
	pool *pgxpool.Pool
}

func (r *userService) User(ctx context.Context, id model.UUID) (*model.User, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, name, email FROM users WHERE id = $1;`, id.UUID)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
