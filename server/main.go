package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saimahesh88/golang-react-todo/router"
)

func main() {
	r := router.Router() //router is the package, Router is the function
	fmt.Println("Starting server on port 9000....")
	log.Fatal(http.ListenAndServe(":9000", r))
}
