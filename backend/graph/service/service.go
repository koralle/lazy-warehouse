package service

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func New(exec boil.ContextExecutor) Services {
	return &services{
		roleService:  &roleService{exec: exec},
		groupService: &groupService{exec: exec},
		userService:  &userService{exec: exec},
	}
}
