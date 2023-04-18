package repository_impl

import (
	"errors"
	"log"

	"github.com/haviz000/API-multi-level-user-golang/model"
	"github.com/haviz000/API-multi-level-user-golang/repository"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) CreateProduct(product model.Product) (*model.Product, error) {
	newProduct := model.Product{
		ProductID:   product.ProductID,
		Title:       product.Title,
		Description: product.Description,
		UserID:      product.UserID,
	}

	err := repository.DB.Create(&newProduct).Error
	if err != nil {
		log.Fatal("error")
		return nil, err
	}

	return &newProduct, nil
}

func (repository *ProductRepositoryImpl) FindProduct(productID string) (*model.Product, error) {
	productResult := model.Product{}

	err := repository.DB.Debug().Where("product_id = ?", productID).Take(&productResult).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}

	return &productResult, nil
}

func (repository *ProductRepositoryImpl) GetByUserID(userID string) ([]model.Product, error) {
	products := make([]model.Product, 0)
	tx := repository.DB.Where("user_id = ?", userID).Find(&products)
	return products, tx.Error
}

func (repository *ProductRepositoryImpl) GetAllProduct() ([]model.Product, error) {
	products := []model.Product{}

	err := repository.DB.Find(&products).Error
	if err != nil {
		return []model.Product{}, err
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) DeleteProduct(product model.Product) error {
	err := repository.DB.Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *ProductRepositoryImpl) UpdateProduct(product model.Product) (model.Product, error) {
	productUpdated := product

	err := repository.DB.Model(&productUpdated).Updates(model.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error
	if err != nil {
		return model.Product{}, err
	}

	return productUpdated, nil
}
