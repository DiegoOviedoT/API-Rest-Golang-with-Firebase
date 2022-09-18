package models

import (
	"context"
	"log"
	"os"
	postEntity "servidorWeb/models/post/entity"

	firestore "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type PostRepository interface {
	Save(post *postEntity.Post) (*postEntity.Post, error)
	FindAll() ([]postEntity.Post, error)
}

type repository struct{}

func NewPostRepository() PostRepository {
	return &repository{}
}

var (
	collectionName string = "posts"
)

func (*repository) Save(post *postEntity.Post) (*postEntity.Post, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firestore.NewApp(ctx, nil, sa)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return post, nil
}

func (*repository) FindAll() ([]postEntity.Post, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firestore.NewApp(ctx, nil, sa)

	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}
	client, err := app.Firestore(ctx)

	if err != nil {
		log.Println(err)
	}

	defer client.Close()

	var posts []postEntity.Post
	itr := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := itr.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to iterate: %v", err)
		}

		posts = append(posts, postEntity.Post{
			Id:    doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		})
	}

	return posts, nil
}
