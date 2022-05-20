package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=gorm-test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
