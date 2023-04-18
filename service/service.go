package service

import "github.com/haviz000/API-multi-level-user-golang/model"

type UserService interface {
	CreateNewUser(request model.UserRegisterRequest) (model.UserRegisterResponse, error)
	Login(request model.UserLoginRequest) (string, error)
}

type ProductService interface {
	CreateProduct(request model.ProductCreateRequest, userId string) (model.ProductCreateResponse, error)
	GetProductByUserID(userId string) ([]model.ProductResponse, error)
	GetProductByID(productID string, userID string, role bool) (model.ProductResponse, error)
	GetAllProduct() ([]model.ProductResponse, error)
	UpdatedProduct(productID string, request model.ProductUpdateRequest, userId string) (model.ProductResponse, error)
	DeleteProduct(productID string) error
}
