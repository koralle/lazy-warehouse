package service

import (
	"context"
	"database/sql"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type userService struct {
	db *sql.DB
}

func (r *userService) User(ctx context.Context, id string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, name, email FROM users WHERE id = $1`, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
