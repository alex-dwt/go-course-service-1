package postgresql

import (
	"context"

	"alex/test/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, userID int) (*model.User, error) {
	user := model.User{}
	result := r.db.First(&user, userID)

	return &user, result.Error
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID int) error {
	return nil
}

func (r *UserRepository) GetAllUsersWithAge(ctx context.Context, age int) ([]model.User, error) {
	return []model.User{}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	result := r.db.Create(user)

	return result.Error
}

func (r *UserRepository) Save(ctx context.Context, user *model.User) error {
	result := r.db.Save(user)

	return result.Error
}
