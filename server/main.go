package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/saimahesh88/golang-react-todo/router"
)

func main() {
	r := router.Router() //router is the package, Router is the function
	fmt.Println("Starting server on port 9000....")
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type"}),
		)(r)))
}
