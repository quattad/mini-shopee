package repository

import "github/quattad/mini-shopee/products-service/models"

type ProductRepository interface {
	Save(product models.Product) (models.Product, error)
	FindAll() ([]models.Product, error)
	FindById(pid uint32) (models.Product, error)
	Update(pid uint32, product models.Product) (int64, error)
	Delete(pid uint32) (int64, error)
}
