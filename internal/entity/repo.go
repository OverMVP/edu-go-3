package entity

type UserRepository interface {
	Save(user User) error
	FindById(id string) (User, error)
	FindAll() []User
	DeleteById(id string) error
}
