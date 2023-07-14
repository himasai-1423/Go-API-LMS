package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReturnBook(coll *mongo.Collection, ctx context.Context, userName string) {
	fmt.Println("Enter the ID of book you want to return:- ")
	var returnID int32
	fmt.Scan(&returnID)

	filter := bson.M{"bookId": returnID}
	var book bson.M
	err := coll.FindOne(ctx, filter).Decode(&book)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Book not found.")
		} else {
			panic(err)
		}
		return
	}

	fmt.Println("\nHere are the details of the book you are returning:- ")
	PrintBookDetails(book)

	quantity, _ := book["quantity"].(int32)

	quantity++
	fmt.Printf("New quantity: %d\n", quantity)

	update := bson.M{"$set": bson.M{"quantity": quantity}}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Thank you %v for returning %v book!!\n", userName, book["name"])

}
