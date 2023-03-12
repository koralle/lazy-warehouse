package service

import (
	"context"

	"github.com/koralle/lazy-warehouse/backend/graph/db"
	"github.com/koralle/lazy-warehouse/backend/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func convertUser(group *db.User) *model.User {
	return &model.User{
		ID:   group.ID,
		Name: group.Name,
	}
}

func (r *userService) User(ctx context.Context, id string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserColumns.ID, db.UserColumns.Name),
		db.UserWhere.ID.EQ(id),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}
