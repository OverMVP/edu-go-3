package repo

import (
	"log/slog"

	"l3/internal/entity"
)

type (
	mockUserRepo struct{}
)

var _ entity.UserRepository = (*mockUserRepo)(nil)

func NewMockUserRepo() *mockUserRepo {
	return &mockUserRepo{}
}

func (m *mockUserRepo) DeleteById(id string) error {
	slog.Info("User with id: %v was successfully deleted\n", "id", id)
	return nil
}

func (m *mockUserRepo) FindAll() []entity.User {
	slog.Info("FindAll was successfully called\n")
	return nil
}

func (m *mockUserRepo) FindById(id string) (entity.User, error) {
	slog.Info("FindById was successfully called\n")
	return entity.User{}, nil
}

func (m *mockUserRepo) FindByRole(role string) ([]entity.User, error) {
	slog.Info("FindByRole was successfully called\n")
	return nil, nil
}

func (m *mockUserRepo) Save(user entity.User) error {
	slog.Info("Save was successfully called\n")
	return nil
}
