package main

import (
	"context"
	"fmt"
	"lib-mng-sys/controllers"
	model "lib-mng-sys/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB_URI string = "mongodb://root:example@mongodb:27017"

// var DB_URI string = "mongodb://localhost:27017"
// var DB_URI string = "mongodb+srv://borahimasaireddy:himu2003@cluster0.daxzqzv.mongodb.net/?retryWrites=true&w=majority"

func main() {
	// # Establishing connection
	fmt.Println("Starting Connection")
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

	coll := client.Database("ABCLibrary").Collection("Books")
	model.CreateColl(coll, client, ctx)

	router := gin.Default()

	controllers.HttpCallRoutes(coll, ctx, router)
	router.Run("0.0.0.0:3000")
}
