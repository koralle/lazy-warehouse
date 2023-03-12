package service

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/db"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type groupService struct {
	exec boil.ContextExecutor
}

func convertGroup(group *db.Group) *model.Group {
	return &model.Group{
		ID:   group.ID,
		Name: group.Name,
	}
}

func (r *groupService) Group(ctx context.Context, id string) (*model.Group, error) {
	group, err := db.Groups(
		qm.Select(db.GroupColumns.ID, db.GroupColumns.Name),
		db.GroupWhere.ID.EQ(id),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertGroup(group), nil
}
