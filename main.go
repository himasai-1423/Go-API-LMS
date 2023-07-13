package main

import (
	"context"
	"fmt"
	"lib-mng-sys/views"

	"go.mongodb.org/mongo-driver/bson"
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

	// Check if the collection already exists
	names, err := client.Database("ABCLibrary").ListCollectionNames(ctx, bson.M{"name": "Books"})
	if err != nil {
		fmt.Println("Failed to check collection existence:", err)
		return
	}

	collectionExists := len(names) > 0
	if !collectionExists {
		// Create the collection
		err = client.Database("ABCLibrary").CreateCollection(ctx, "Books")
		if err != nil {
			fmt.Println("Failed to create collection:", err)
			return
		}

		// Create the indexes
		indexes := []mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "bookID", Value: 1}},
				Options: options.Index().SetName("Idx").SetUnique(true),
			},
		}
		_, err = coll.Indexes().CreateMany(ctx, indexes)
		if err != nil {
			fmt.Println("Failed to create indexes:", err)
			return
		}

		fmt.Println("Collection created with indexes")
	}

	views.FlowControl(coll, ctx)
}
