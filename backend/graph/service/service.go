package service

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Services interface {
	RoleService
}

type services struct {
	*roleService
}

type RoleService interface {
	GetAllAvailableRole(ctx context.Context) ([]*model.Role, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		roleService: &roleService{exec: exec},
	}
}
