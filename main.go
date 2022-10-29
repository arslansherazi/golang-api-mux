package main

import (
	"find_competitor/common"
	router "find_competitor/routing"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(common.ENVIRONMENT_VARIBALES_ERROR_MESSAGE)
	}

	router := router.RouterV1()
	fmt.Println(http.ListenAndServe(":4000", router))
}
