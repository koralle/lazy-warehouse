package service

import (
	"context"
	"database/sql"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
)

type Services interface {
	RoleService
	GroupService
	UserService
}

type services struct {
	*roleService
	*groupService
	*userService
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

func New(db *sql.DB) Services {
	return &services{
		roleService:  &roleService{db: db},
		groupService: &groupService{db: db},
		userService:  &userService{db: db},
	}
}
