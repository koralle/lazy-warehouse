package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type Services interface {
	RoleService
	GroupService
	UserService
	SetRoleService
}

type services struct {
	*roleService
	*groupService
	*userService
	*setRoleService
}

type RoleService interface {
	GetAllAvailableRole(ctx context.Context) ([]*model.Role, error)
}

type GroupService interface {
	Group(ctx context.Context, id string) (*model.Group, error)
}

type UserService interface {
	User(ctx context.Context, id string) (*model.User, error)
}

type SetRoleService interface {
	SetRoleToUserInGroup(ctx context.Context, input model.SetRoleInput) (*model.SetRoleResult, error)
}

func New(pool *pgxpool.Pool) Services {
	return &services{
		roleService:    &roleService{pool: pool},
		groupService:   &groupService{pool: pool},
		userService:    &userService{pool: pool},
		setRoleService: &setRoleService{pool: pool},
	}
}
