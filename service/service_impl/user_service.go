package service_impl

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/haviz000/API-multi-level-user-golang/helpers"
	"github.com/haviz000/API-multi-level-user-golang/model"
	"github.com/haviz000/API-multi-level-user-golang/repository"
	"github.com/haviz000/API-multi-level-user-golang/service"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) CreateNewUser(request model.UserRegisterRequest) (model.UserRegisterResponse, error) {
	userID := uuid.New()
	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	var role bool
	if request.Role == "admin" {
		role = true
	} else {
		role = false
	}

	user := model.User{
		UserID:   userID,
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     role,
	}

	result, err := service.UserRepository.CreateUser(user)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	var roleString string
	if result.Role == true {
		roleString = "admin"
	} else {
		roleString = "user"
	}
	response := model.UserRegisterResponse{
		UserID: result.UserID,
		Name:   result.Name,
		Email:  result.Email,
		Role:   roleString,
	}

	return response, nil
}

func (service *UserService) Login(request model.UserLoginRequest) (string, error) {
	result, err := service.UserRepository.UserCheckByEmail(request.Email)
	if err != nil {
		return "", err
	}

	isMatch := helpers.PasswordIsMatch(request.Password, result.Password)
	if isMatch == false {
		return "", errors.New(fmt.Sprintf("Invalid username or password"))
	}

	myClaim := helpers.MyClaims{
		UserID: result.UserID.String(),
		Role:   result.Role,
	}
	jwtToken, err := helpers.GenerateToken(myClaim)

	return jwtToken, err
}
