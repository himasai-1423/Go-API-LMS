package main

import (
	"context"
	"fmt"
	model "lib-mng-sys/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_URI string = "mongodb://localhost:27017"

func main() {
	// # Establishing connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URI))
	ctx := context.TODO()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connection sucessfully established!")

	// # Adding data
	coll := client.Database("ABCLibrary").Collection("Books")
	// AddData(coll, ctx)

	model.CreateColl(coll, client, ctx)

	router := gin.Default()

	router.GET("/FindBooks", func(c *gin.Context) {
		model.BooksAvailable(coll, ctx, c)
	})
	router.GET("/FindBooks/:genre", func(c *gin.Context) {
		genre := c.Param("genre")
		model.FilterByGenre(coll, ctx, c, genre)
	})
	router.POST("/TakeBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.RentBook(coll, ctx, c, requestBody.BookId)
	})
	router.POST("/ReturnBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.ReturnBook(coll, ctx, c, requestBody.BookId)
	})
	router.POST("/Admin/AddBook", func(c *gin.Context) {
		var newData model.BookData

		if err := c.ShouldBindJSON(&newData); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		model.AddData(coll, ctx, c, newData)
	})
	router.POST("/Admin/DeleteBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int32 `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.DeleteBook(coll, ctx, c, requestBody.BookId)
	})

	router.Run("localhost:9090")
}
