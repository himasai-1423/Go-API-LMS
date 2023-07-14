package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookData struct {
	Name     string   `json:"name" bson:"name"`
	Author   string   `json:"author" bson:"author"`
	BookId   int32    `json:"bookId" bson:"bookId"`
	Genre    []string `json:"genre" bson:"genre"`
	Quantity int32    `json:"quantity" bson:"quantity"`
}

func BooksAvailable(coll *mongo.Collection, ctx context.Context, context *gin.Context) {
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bookList BookData
		if err = cursor.Decode(&bookList); err != nil {
			panic(err)
		}
		context.IndentedJSON(http.StatusOK, bookList)
	}
}
