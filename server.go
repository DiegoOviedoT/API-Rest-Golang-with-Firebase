package main

import (
	"log"
	"net/http"
	"os"

	"servidorWeb/router"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	route router.Route = router.NewRoute()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("SERVER_PORT")
	router := mux.NewRouter()

	route.ListRoute(router)

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
