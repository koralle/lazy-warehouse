package service

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type setRoleService struct {
	pool *pgxpool.Pool
}

func (s *setRoleService) SetRoleToUserInGroup(ctx context.Context, input model.SetRoleInput) (*model.SetRoleResult, error) {
	var user model.User
	err := s.pool.QueryRow(ctx, `SELECT id, name, email FROM lazy_warehouse.users WHERE id = $1;`, input.UserID).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	var group model.Group
	err = s.pool.QueryRow(ctx, `SELECT id, name, description FROM lazy_warehouse.groups WHERE id = $1;`, input.GroupID).Scan(&group.ID, &group.Name, &group.Description)
	if err != nil {
		return nil, err
	}

	var role model.Role
	err = s.pool.QueryRow(ctx, `SELECT id, name FROM lazy_warehouse.roles WHERE name = $1;`, input.Role).Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}

	upsertQuery := `
    INSERT INTO lazy_warehouse.administration 
      (id, user_id, group_id, role_id)
    VALUES ($1, $2, $3, $4) 
    ON CONFLICT ON CONSTRAINT user_and_group
    DO UPDATE SET role_id=$4;
  `

	u6, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	_, err = s.pool.Query(ctx, upsertQuery, u6, user.ID, group.ID, role.ID)
	if err != nil {
		return nil, err
	}

	return &model.SetRoleResult{User: &user, Group: &group, Role: role.Name}, nil
}
