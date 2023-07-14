package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RentBook(coll *mongo.Collection, ctx context.Context, context *gin.Context, reqBookId int) {
	filter := bson.M{"bookId": reqBookId}

	var book BookData
	err := coll.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book Not Found"})
		return
	}

	if book.Quantity > 0 {
		book.Quantity--
		update := bson.M{"$set": bson.M{"quantity": book.Quantity}}

		_, err := coll.UpdateOne(ctx, filter, update)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update quantity"})
			return
		}
	} else {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "Book out of stock"})
		return
	}
	context.IndentedJSON(http.StatusAccepted, book)
}
