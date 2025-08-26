package repo

import (
	"errors"

	"l3/internal/constants"
	"l3/internal/entity"
)

type (
	InMemoryUserRepo struct {
		users map[string]entity.User
	}
)

var _ entity.UserRepository = (*InMemoryUserRepo)(nil)

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]entity.User),
	}
}

func (r *InMemoryUserRepo) checkIsNewUser(email string) bool {
	for _, v := range r.users {
		if v.Email == email {
			return false
		}
	}
	return true
}

func (r *InMemoryUserRepo) Save(user entity.User) error {
	if !r.checkIsNewUser(user.Email) {
		return errors.New(constants.ErrUserAlreadyExists)
	}

	r.users[user.ID] = user

	return nil
}

func (r *InMemoryUserRepo) FindAll() []entity.User {
	users := make([]entity.User, 0, len(r.users))

	for _, v := range r.users {
		users = append(users, v)
	}

	return users
}

func (r *InMemoryUserRepo) DeleteById(id string) error {
	if _, ok := r.users[id]; !ok {
		return errors.New(constants.ErrUserNotFound)
	}

	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepo) FindById(id string) (entity.User, error) {
	if u, ok := r.users[id]; !ok {
		return entity.User{}, errors.New(constants.ErrUserNotFound)
	} else {
		return u, nil
	}
}

func (r *InMemoryUserRepo) FindByRole(role string) ([]entity.User, error) {
	if !entity.UserRole(role).Valid() {
		return nil, errors.New(constants.ErrUserNotExistingRole)
	}

	users := make([]entity.User, 0, len(r.users))

	for _, v := range r.users {
		if v.Role == entity.UserRole(role) {
			users = append(users, v)
		}
	}

	return users, nil
}
