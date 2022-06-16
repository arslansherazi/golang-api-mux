package main

import (
	"find_competitor/configs"
	"find_competitor/models"
	"fmt"
)

func main() {
	db, err := configs.GetDbInstance()
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		db.AutoMigrate(&models.User{})
	}
}
