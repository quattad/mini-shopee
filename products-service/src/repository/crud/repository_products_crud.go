package crud

import (
	"errors"
	"github/quattad/mini-shopee/products-service/models"
	"github/quattad/mini-shopee/products-service/src/repository"
	"github/quattad/mini-shopee/products-service/src/utils/channels"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	RepositoryProductsCrudService repository.ProductRepository
)

func init() {
	RepositoryProductsCrudService = &RepositoryProductsCrud{}
}

type RepositoryProductsCrud struct {
	db *gorm.DB
}

func NewRepositoryProductsCrud(db *gorm.DB) *RepositoryProductsCrud {
	return &RepositoryProductsCrud{db}
}

/*
SAVE
*/
func (r *RepositoryProductsCrud) Save(product models.Product) (models.Product, error) {
	var err error
	done := make(chan bool)

	go func(ch chan<- bool) {
		defer close(ch)
		err := r.db.Debug().Model(&models.Product{}).Create(&product).Error

		if err != nil {
			ch <- false
			return
		}

		ch <- true
	}(done)

	if channels.OK(done) {
		return product, nil
	}
	return models.Product{}, err
}

/*
FINDALL
*/
func (r *RepositoryProductsCrud) FindAll() ([]models.Product, error) {
	var err error
	done := make(chan bool)

	products := []models.Product{}

	go func(ch chan<- bool) {
		err := r.db.Debug().Model(&models.Product{}).Limit(100).Find(&products).Error

		if err != nil {
			ch <- false
			return
		}

		ch <- true
		return
	}(done)

	if channels.OK(done) {
		return products, nil
	}

	return []models.Product{}, err
}

/*
FIND BY ID
*/
func (r *RepositoryProductsCrud) FindById(pid uint32) (models.Product, error) {
	var err error

	product := models.Product{}
	done := make(chan bool)

	go func(ch chan<- bool) {
		err = r.db.Debug().Model(&models.Product{}).Where("id=?", pid).Take(&product).Error

		if err != nil {
			ch <- false
			return
		}

		ch <- true
		close(ch)
	}(done)

	if channels.OK(done) {
		return product, nil
	}

	if gorm.IsRecordNotFoundError(err) {
		return models.Product{}, errors.New("Product not found")
	}

	return models.Product{}, err
}

/*
UPDATE
*/
func (r *RepositoryProductsCrud) Update(pid uint32, product models.Product) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		rs = r.db.Debug().Model(&models.Product{}).Where("id=?", pid).Take(&models.Product{}).UpdateColumns(
			map[string]interface{}{
				"name":        product.Name,
				"description": product.Description,
				"updated_at":  time.Now(),
			},
		)

	}(done)

	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

/*
DELETE
*/
func (r *RepositoryProductsCrud) Delete(pid uint32) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)

	go func(ch chan<- bool) {
		rs = r.db.Debug().Model(&models.Product{}).Where("id=?", pid).Take(&models.Product{}).Delete(&models.Product{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}

	return 0, rs.Error
}
