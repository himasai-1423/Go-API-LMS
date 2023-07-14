package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddData(coll *mongo.Collection, ctx context.Context, context *gin.Context, newData BookData) {
	filter := bson.M{"bookId": newData.BookId}

	var book BookData
	err := coll.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		_, err := coll.InsertOne(ctx, newData)
		if err != nil {
			context.IndentedJSON(http.StatusFailedDependency, gin.H{"Error": "Unable to insert Data due to wrong input"})
			return
		}

		context.IndentedJSON(http.StatusAccepted, newData)
		return
	} else {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Book Id already exists"})
		return
	}
}
