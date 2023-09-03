package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cenkY00/mongoAPI/router"
)

func main() {
	fmt.Println("MongoDB api")

	r := router.Router()

	fmt.Println("Server is starting...")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server is running on port 8000")
}
