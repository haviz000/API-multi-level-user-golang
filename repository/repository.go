package repository

import "github.com/haviz000/API-multi-level-user-golang/model"

type UserRepository interface {
	CreateUser(user model.User) (*model.User, error)
	UserCheck(userId string) (*model.User, error)
	UserCheckByEmail(email string) (*model.User, error)
}

type ProductRepository interface {
	CreateProduct(product model.Product) (*model.Product, error)
	FindProduct(productID string) (*model.Product, error)
	GetByUserID(userID string) ([]model.Product, error)
	GetAllProduct() ([]model.Product, error)
	DeleteProduct(product model.Product) error
	UpdateProduct(product model.Product) (model.Product, error)
}
