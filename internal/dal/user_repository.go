package dal

import (
	"errors"

	"github.com/RisingStack/almandite-user-service/internal/dal/models"
	"github.com/go-pg/pg"
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Fetch() (*[]models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id int) error
}

type userRepository struct {
	DB *pg.DB
}

// NewUserRepository returns a repository that implements the UserRepository interface
func newUserRepository(dbConn *pg.DB) UserRepository {
	return &userRepository{
		DB: dbConn,
	}
}

func (u *userRepository) GetByUsername(username string) (*models.User, error) {
	var users []models.User

	if err := u.DB.Model(&users).Where("username = ?", username).Select(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("User not found")
	}

	return &users[0], nil
}

func (u *userRepository) Fetch() (*[]models.User, error) {
	var users []models.User

	if err := u.DB.Model(&users).Select(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *userRepository) GetByID(id int) (*models.User, error) {
	user := models.User{ID: id}

	if err := u.DB.Select(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Create(user *models.User) error {
	return u.DB.Insert(user)
}

func (u *userRepository) Update(user *models.User) error {
	return u.DB.Update(user)
}

func (u *userRepository) Delete(id int) error {
	user := models.User{ID: id}

	return u.DB.Delete(&user)
}
