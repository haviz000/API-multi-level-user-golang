package repository_impl

import (
	"errors"

	"github.com/haviz000/API-multi-level-user-golang/model"
	"github.com/haviz000/API-multi-level-user-golang/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) CreateUser(user model.User) (*model.User, error) {
	newUser := model.User{
		UserID:   user.UserID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}

	err := repository.DB.Create(&newUser).Error
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (repository *UserRepositoryImpl) UserCheck(userId string) (*model.User, error) {
	userResult := model.User{}

	err := repository.DB.Debug().Where("user_id = ?", userId).Take(&userResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &userResult, nil
}

func (repository *UserRepositoryImpl) UserCheckByEmail(email string) (*model.User, error) {
	userResult := model.User{}

	err := repository.DB.Debug().Where("email = ?", email).Take(&userResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &userResult, nil
}
