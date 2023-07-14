package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateColl(coll *mongo.Collection, client *mongo.Client, ctx context.Context) {
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

}
