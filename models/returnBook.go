package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReturnBook(coll *mongo.Collection, ctx context.Context, context *gin.Context, reqBookId int) {
	filter := bson.M{"bookId": reqBookId}

	var book BookData
	if err := coll.FindOne(ctx, filter).Decode(&book); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book Not Found"})
		return
	}

	book.Quantity = book.Quantity + 1
	update := bson.M{"$set": bson.M{"quantity": book.Quantity}}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update quantity"})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}
