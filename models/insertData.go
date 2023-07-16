package model

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddData(coll *mongo.Collection, ctx context.Context, c *gin.Context, newData BookData) {
	filter := bson.M{"bookId": newData.BookId}
	count, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check document existence"})
		return
	}

	if count > 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book ID already exists"})
		return
	}

	_, err = coll.InsertOne(ctx, newData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newData)
}
