package model

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FilterByGenre(coll *mongo.Collection, ctx context.Context, context *gin.Context, reqGenre string) {
	cursor, err := coll.Find(ctx, bson.M{})

	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	cnt := 0
	for cursor.Next(ctx) {
		var bookList BookData
		if err = cursor.Decode(&bookList); err != nil {
			panic(err)
		}
		for _, g := range bookList.Genre {
			if strings.EqualFold(g, reqGenre) {
				cnt++
				context.IndentedJSON(http.StatusOK, bookList)
				break
			}
		}
	}

	if cnt == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "The Requested Genre is not Found"})
	}
}
