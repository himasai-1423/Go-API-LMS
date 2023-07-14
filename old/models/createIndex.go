package model

// import (
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func ensureIndex(coll *mongo.Collection) error {
// 	ctx := context.TODO()
// 	indexModel := mongo.IndexModel{
// 		Keys:    bson.D{{Key: "bookID", Value: 1}},
// 		Options: options.Index().SetName("Idx").SetUnique(true),
// 	}

// 	cursor, err := coll.Indexes().List(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var res bson.M

// 		if err := cursor.Decode(&res); err != nil {
// 			return err
// 		}

// 		indexName, _ := res["name"].(string)
// 		if indexName == "Idx" {
// 			fmt.Println("Index already exists")
// 			return nil
// 		}
// 	}

// 	_, err = coll.Indexes().CreateOne(ctx, indexModel)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("Index created")
// 	return nil
// }
