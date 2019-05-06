package repository

import (
	"github.com/tjeerddie/basic-go-api/entities"
)

type Repository interface {
	Users() ([]entities.User, error)
	UserSingle(id int) (*entities.User, error)
	UserStore(user *entities.User) (error)
	// UserUpdate(user entities.User) (entities.User, error)
	// UserDelete(user entities.User) error
}
