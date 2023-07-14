package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RentBook(coll *mongo.Collection, ctx context.Context, userName string) {

	// var reqGenre string
	// fmt.Println("Select any genre you want:- ")
	// fmt.Scan(&reqGenre)

	// checkGenre(coll, ctx, reqGenre)

	var updateID int32
	fmt.Print("Enter the ID of book you want to rent:- ")
	fmt.Scan(&updateID)

	takeBook(coll, ctx, updateID)

	fmt.Printf("Thank you %v for visiting!!\n", userName)
}

func takeBook(coll *mongo.Collection, ctx context.Context, reqID int32) {
	filter := bson.M{"bookId": reqID}
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

	// fmt.Println("\nHere are the details of the book you are renting:- ")
	// PrintBookDetails(book)

	quantity, _ := book["quantity"].(int32)

	if quantity == 0 {
		fmt.Println("We are sorry to inform, The book is out of stock")
	} else {
		quantity--
		fmt.Printf("New quantity: %d\n", quantity)

		update := bson.M{"$set": bson.M{"quantity": quantity}}
		_, err := coll.UpdateOne(ctx, filter, update)
		if err != nil {
			panic(err)
		}
	}
}
