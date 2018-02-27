package repository

import (
	"github.com/RisingStack/almandite-user-service/model"

	"github.com/go-pg/pg"
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int) (*model.User, error)
}

type userRepository struct {
	Db *pg.DB
}

// NewUserRepository returns a repository that implements the UserRepository interface
func NewUserRepository(dbConn *pg.DB) UserRepository {
	return &userRepository{
		Db: dbConn,
	}
}

func (u *userRepository) GetByID(id int) (*model.User, error) {
	user := model.User{ID: id}

	if err := u.Db.Select(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
