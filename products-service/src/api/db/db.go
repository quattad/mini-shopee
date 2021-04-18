package db

import (
	_ "github.com/go-sql-driver/mysql" // database/sql package must be used together with db driver. import only for side effects
	"github.com/jinzhu/gorm"
)

var (
	DBService dbServiceInterface
)

func init() {
	DBService = &dbService{}
}

type dbService struct {
	Name string
}

type dbServiceInterface interface {
	Connect(string, string) (*gorm.DB, error)
}

// Connect tos to the db and returns a gorm.db instance
func (d *dbService) Connect(DBDRIVER, DBURL string) (*gorm.DB, error) {
	db, err := gorm.Open(DBDRIVER, DBURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}
