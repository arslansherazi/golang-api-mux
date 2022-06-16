package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbInstance() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=fnd_comp_db port=5432 sslmode=disable TimeZone=Asia/Karachi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
