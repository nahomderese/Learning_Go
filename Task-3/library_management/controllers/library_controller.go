package controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/Nahom-Derese/Learning_Go/Task-3/library_management/models"
	"github.com/Nahom-Derese/Learning_Go/Task-3/library_management/services"
)

var library services.LibraryCatalog = services.LibraryService{
	Books:   []models.Book{},
	Members: []models.Member{},
}

func AddBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
	library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func RemoveBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	library.RemoveBook(id)
	fmt.Println("Book removed successfully!")
}

func BorrowBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func ReturnBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error returning book:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func ListAvailableBooks() {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func ListBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	books := library.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
