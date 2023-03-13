package service

import (
	"context"
	"database/sql"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type groupService struct {
	db *sql.DB
}

func (r *groupService) Group(ctx context.Context, id string) (*model.Group, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, name, description FROM lazy_warehouse.groups WHERE id =  $1;`, id)

	var group model.Group
	if err := row.Scan(&group.ID, &group.Name, &group.Description); err != nil {
		return nil, err
	}

	return &group, nil
}
