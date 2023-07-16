package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteBook(coll *mongo.Collection, ctx context.Context, context *gin.Context, reqBookId int32) {
	filter := bson.M{"bookId": reqBookId}

	var book BookData
	err := coll.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Delete Failed": "Book Doesnot Exist"})
		return
	}
	context.IndentedJSON(http.StatusAccepted, book)

	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Delete Failed": "Book Doesnot Exist"})
		return
	}
}
