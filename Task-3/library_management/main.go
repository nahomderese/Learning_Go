package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Nahom-Derese/Learning_Go/Task-3/library_management/controllers"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// fmt.Print("\033[H\033[2J")
		fmt.Println("Library Management System")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			controllers.AddBook(reader)
		case "2":
			controllers.RemoveBook(reader)
		case "3":
			controllers.BorrowBook(reader)
		case "4":
			controllers.ReturnBook(reader)
		case "5":
			controllers.ListAvailableBooks()
		case "6":
			controllers.ListBorrowedBooks(reader)
		case "7":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
