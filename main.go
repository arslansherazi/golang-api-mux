package main

import (
	router "find_competitor/routing"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := router.RouterV1()
	fmt.Println("Server is listening at port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
