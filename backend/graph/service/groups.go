package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type groupService struct {
	pool *pgxpool.Pool
}

func (r *groupService) Group(ctx context.Context, id string) (*model.Group, error) {
	row := r.pool.QueryRow(ctx, `SELECT id, name, description FROM groups WHERE id = $1;`, id)

	var group model.Group
	if err := row.Scan(&group.ID, &group.Name, &group.Description); err != nil {
		return nil, err
	}

	return &group, nil
}
