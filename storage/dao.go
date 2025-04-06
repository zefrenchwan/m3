package storage

import "github.com/zefrenchwan/m3.git/entities"

type Dao interface {
	// Dislays details on the type of DAO
	Info() string
	// User management
	Users() ([]entities.User, error)
	UpsertUser(user entities.User) error
	DeleteUser(user entities.User) error
}
