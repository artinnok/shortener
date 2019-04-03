package web

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5435 user=short dbname=short password=short sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
