package main

import (
	"find_competitor/configs"
	"find_competitor/models"
)

func main() {
	db := configs.GetDbInstance()
	db.AutoMigrate(&models.User{})
}
