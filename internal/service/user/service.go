package user

import (
	"context"

	"alex/test/internal/model"
)

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUserWithSomeLogic(ctx context.Context, name string, age int) (*model.User, error) {
	user := &model.User{
		Name: name,
		Age:  uint8(age),
	}

	return user, s.repository.CreateUser(ctx, user)
}

func (s *Service) GetUserWithSomeLogic(ctx context.Context, userID int) (*model.User, error) {
	return s.repository.GetUser(ctx, userID)
}

func (s *Service) DeleteUserWithSomeLogic(ctx context.Context, userID int) error {
	return nil
}

func (s *Service) GetUsersByAge(ctx context.Context, age int) ([]model.User, error) {
	return []model.User{}, nil
}

func (s *Service) SaveUser(ctx context.Context, user *model.User) error {
	return s.repository.Save(ctx, user)
}
