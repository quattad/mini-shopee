package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Product struct {
	ID          uint32    `gorm:"primary key;auto_increment" json:"id"`
	Name        string    `gorm:"size:100;not null;unique" json:"name"`
	Description string    `gorm:"size:500;not null" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	CreatedAt   time.Time `gorm:"default:current_timestamp()" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp()" json:"updated_at"`
}

type productsServiceInterface interface {
	Prepare() *Product
	Validate(string) error
}

// Prepare removes any whitespaces from fields except for Description
func (p *Product) Prepare() *Product {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return p

}

// Validate acts as a sanitizer for fields
func (p *Product) Validate(action string) error {
	switch action {
	case "update":
		if p.Name == "" {
			return errors.New("Product must have defined property 'name'")
		}

		if p.Description == "" {
			return errors.New("Product must have defined property 'description'")
		}

		if p.Price == 0 {
			return errors.New("Product cannot be 0 for defined property 'price'")
		}

		return nil

	default:
		if p.Name == "" {
			return errors.New("Product must have defined property 'name'")
		}

		if p.Description == "" {
			return errors.New("Product must have defined property 'description'")
		}

		if p.Price == 0 {
			return errors.New("Product cannot be 0 for defined property 'price'")
		}

		return nil

	}
}
