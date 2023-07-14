package main

import (
	"context"
	"fmt"
	model "lib-mng-sys/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_URI string = "mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority"

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

	router.Run("localhost:9090")
}
