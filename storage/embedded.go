package storage

import "github.com/zefrenchwan/m3.git/entities"

type EmbeddedDao struct {
}

// Info returns that it is an SQLITE impl
func (e EmbeddedDao) Info() string {
	return "SQLITE implementation"
}

func (e EmbeddedDao) Users() ([]entities.User, error) {
	return nil, nil
}

func (e EmbeddedDao) UpsertUser(user entities.User) error {
	return nil
}

func (e EmbeddedDao) DeleteUser(user entities.User) error {
	return nil
}

func NewEmbeddedDao(path string) (Dao, error) {
	return EmbeddedDao{}, nil
}
