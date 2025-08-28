package service

import (
	"fmt"
	"time"

	"l3/internal/constants"
	"l3/internal/entity"
)

type UserService struct {
	repo entity.UserRepository
}

func NewUserService(repo entity.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUser(id string) (entity.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("looking for user with id: %s, error: %w", id, err)
	}

	return user, nil
}

func (u *UserService) CreateUser(name string, email string, role entity.UserRole) (entity.User, error) {
	user := entity.User{
		Name:      name,
		Email:     email,
		Role:      role,
		ID:        email,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Save(user); err != nil {
		return entity.User{}, fmt.Errorf("saving user: %w", err)
	}

	return user, nil
}

func (u *UserService) ListUsers() []entity.User {
	return u.repo.FindAll()
}

func (s *UserService) FindByRole(role string) ([]entity.User, error) {
	uRole := entity.UserRole(role)

	if !uRole.Valid() {
		return nil, fmt.Errorf("validating role %s, error: %w", role, constants.ErrUserNotExistingRole)
	}

	users := s.repo.FindAll()

	var list []entity.User

	for _, v := range users {
		if v.Role == uRole {
			list = append(list, v)
		}
	}

	return list, nil
}

func (u *UserService) RemoveUser(id string) error {
	err := u.repo.DeleteById(id)
	if err != nil {
		return fmt.Errorf("deleting user with id: %s, error: %w", id, err)
	}

	return nil
}
