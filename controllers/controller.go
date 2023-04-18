package controllers

import "github.com/gin-gonic/gin"

type UserController interface {
	Registration(c *gin.Context)
	Login(c *gin.Context)
}

type ProductController interface {
	CreateProduct(c *gin.Context)
	GetAllProduct(c *gin.Context)
	GetProductByRole(c *gin.Context)
	GetOneProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
}
