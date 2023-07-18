package controllers

import (
	"context"
	model "lib-mng-sys/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HttpCallRoutes(coll *mongo.Collection, ctx context.Context, router *gin.Engine) {
	router.GET("/FindBooks", func(c *gin.Context) {
		model.BooksAvailable(coll, ctx, c)
	})
	router.GET("/FindBooks/:genre", func(c *gin.Context) {
		genre := c.Param("genre")
		model.FilterByGenre(coll, ctx, c, genre)
	})
	router.POST("/TakeBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.RentBook(coll, ctx, c, requestBody.BookId)
	})
	router.POST("/ReturnBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.ReturnBook(coll, ctx, c, requestBody.BookId)
	})
	router.POST("/Admin/AddBook", func(c *gin.Context) {
		var newData model.BookData

		if err := c.ShouldBindJSON(&newData); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		model.AddData(coll, ctx, c, newData)
	})

	router.POST("/Admin/DeleteBook", func(c *gin.Context) {
		var requestBody struct {
			BookId int32 `json:"bookId" bson:"bookId"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		model.DeleteBook(coll, ctx, c, requestBody.BookId)
	})
}
