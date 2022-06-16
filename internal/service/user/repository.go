package user

import (
	"context"

	"alex/test/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID int) (*model.User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetAllUsersWithAge(ctx context.Context, age int) ([]model.User, error)
	Save(context.Context, *model.User) error
}
