package configs

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbInstance() (*gorm.DB, error) {

	dsn := "host=" + os.Getenv("DB_HOST") + "user=" + os.Getenv("DB_USER") + "password=" + os.Getenv("DB_PASSWORD") +
		"dbname=" + os.Getenv("DB_NAME") + "port=" + os.Getenv("PORT") + "sslmode=disable TimeZone=Asia/Karachi"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
