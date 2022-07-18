package main

import (
	"find_competitor/configs"
	"find_competitor/models"
	"fmt"
)

func main() {
	isScript := true
	db, err := configs.GetDbInstance(isScript)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	} else {
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Competiton{})
	}
}
