package service

import (
	"context"
	"database/sql"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type roleService struct {
	db *sql.DB
}

func (r *roleService) GetAllAvailableRole(ctx context.Context) ([]*model.Role, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM lazy_warehouse.roles;")

	if err != nil {
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
