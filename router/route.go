package router

import (
	postController "servidorWeb/controllers/post"

	"github.com/gorilla/mux"
)

type Route struct{}

func (*Route) ListRoute(router *mux.Router) {
	router.HandleFunc("/posts", postController.GetPosts).Methods("GET")
	router.HandleFunc("/posts", postController.AddPost).Methods("POST")
}

func NewRoute() Route {
	return Route{}
}
