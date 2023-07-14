package views

// import (
// 	"context"
// 	"fmt"
// 	model "lib-mng-sys/models"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func FlowControl(coll *mongo.Collection, ctx context.Context) {
// 	fmt.Println("\nWelcome to Library Management System!")

// 	fmt.Println("\n\nPlease enter your name: ")
// 	var userName string
// 	fmt.Scan(&userName)
// 	shouldExit := false

// 	for !shouldExit {
// 		choice := UserChoice(userName)

// 		switch choice {
// 		case 1:
// 			model.RentBook(coll, ctx, userName)
// 		case 2:
// 			model.ReturnBook(coll, ctx, userName)
// 		case 3:
// 			model.BooksAvailable(coll, ctx)
// 		default:
// 			Leaving(userName)
// 			shouldExit = true
// 		}
// 	}
// }
