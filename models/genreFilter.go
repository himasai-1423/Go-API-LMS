package model

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func FilterByGenre(coll *mongo.Collection, ctx context.Context, context *gin.Context, genre string) {

}

// func checkGenre(coll *mongo.Collection, ctx context.Context, reqGenre string) {
// 	fmt.Printf("\nHere is the list of available books in %v category:- \n", reqGenre)
// 	cursor, err := coll.Find(ctx, bson.M{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var book bson.M
// 		if err = cursor.Decode(&book); err != nil {
// 			panic(err)
// 		}

// 		if genreRegex, ok := book["genre"].(primitive.Regex); ok {
// 			if strings.Contains(strings.ToLower(genreRegex.Pattern), strings.ToLower(reqGenre)) {
// 				// PrintBookDetails(book)
// 			}
// 		}
// 	}
// }
