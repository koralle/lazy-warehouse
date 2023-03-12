package service

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/db"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type roleService struct {
	exec boil.ContextExecutor
}

func convertRole(role *db.Role) *model.Role {
	return &model.Role{
		ID:   role.ID,
		Name: model.RoleCategory(role.Name),
	}
}

func (r *roleService) GetAllAvailableRole(ctx context.Context) ([]*model.Role, error) {
	roles, err := db.Roles(qm.Select(db.RoleColumns.ID, db.RoleColumns.Name)).All(ctx, r.exec)
	if err != nil {
		return nil, err
	}

	var res []*model.Role

	for idx := 0; idx < len(roles); idx++ {
		res = append(res, convertRole(roles[idx]))
	}

	return res, nil
}
