package views

import "fmt"

func UserChoice(userName string) int {

	fmt.Printf("\n Hello! %v,", userName)
	fmt.Println("\nPlease choose any of the option below:-")
	fmt.Println("1. Rent a book ")
	fmt.Println("2. Return a book")
	fmt.Println("3. Read a book ")
	//add new book
	fmt.Println("Any other key to exit")

	var choice int
	fmt.Scan(&choice)
	return choice
}

func Leaving(userName string) {
	fmt.Printf("Thank you %v for visiting\n", userName)
}
