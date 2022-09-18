package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	postEntity "servidorWeb/models/post/entity"
	postRepository "servidorWeb/models/post/repository"
)

var (
	repo postRepository.PostRepository = postRepository.NewPostRepository()
)

// Resolve the request for the route /posts and return all posts

func GetPosts(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "aplication/json")

	posts, err := repo.FindAll()

	if err != nil {

		response.WriteHeader(http.StatusInternalServerError)

		response.Write([]byte(`"error:" "error getting the posts"`))

		return
	}

	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(posts)
}

func AddPost(response http.ResponseWriter, resquest *http.Request) {
	var post postEntity.Post
	response.Header().Set("Content-Type", "aplication/json")
	err := json.NewDecoder(resquest.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		log.Printf("Error while decoding the request body for addPost: %v", err)
	}
	post.Id = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
