package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BooksAvailable(coll *mongo.Collection, ctx context.Context) {
	fmt.Println("Here is the list of available books:- ")
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bookList bson.M
		if err = cursor.Decode(&bookList); err != nil {
			panic(err)
		}
		PrintBookDetails(bookList)
	}
}

func PrintBookDetails(book bson.M) {
	bookId, _ := book["bookId"].(int32)
	name, _ := book["name"].(string)
	author, _ := book["author"].(string)
	quantity, _ := book["quantity"].(int32)

	fmt.Printf("Book Id: %v\n", bookId)
	fmt.Printf("Book Name: %v\n", name)
	fmt.Printf("Author: %v\n", author)
	fmt.Printf("Quantity: %v\n", quantity)

	if genreRegex, ok := book["genre"].(primitive.Regex); ok {
		fmt.Println("Genre:", genreRegex.Pattern)
	}

	fmt.Println()

}
