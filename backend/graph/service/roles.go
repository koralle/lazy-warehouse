package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type roleService struct {
	pool *pgxpool.Pool
}

func (r *roleService) GetAllAvailableRole(ctx context.Context) ([]*model.Role, error) {
	rows, err := r.pool.Query(ctx, `SELECT id, name FROM lazy_warehouse.roles;`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var res []*model.Role

	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		res = append(res, &role)
	}

	return res, nil
}
