package services

import (
	"errors"

	"github.com/Nahom-Derese/Learning_Go/Task-3/library_management/models"
)

type LibraryCatalog interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type LibraryService struct {
	Books   []models.Book
	Members []models.Member
}

func (service LibraryService) AddBook(book models.Book) {

	service.Books = append(service.Books, book)
}

func (service LibraryService) RemoveBook(bookID int) {
	for i, book := range service.Books {
		if book.ID == bookID {
			service.Books = append(service.Books[:i], service.Books[i+1:]...)
			return
		}
	}
}

func (service LibraryService) BorrowBook(bookID int, memberID int) error {

	for i, book := range service.Books {
		if book.ID == bookID {
			service.Books[i].Status = "Borrowed"
			for _, member := range service.Members {
				if member.ID == memberID {
					member.BorrowedBooks = append(member.BorrowedBooks, book)
					return nil
				}
			}
		}
	}
	return errors.New("Book or Member not found")
}

func (service LibraryService) ReturnBook(bookID int, memberID int) error {

	for i, book := range service.Books {
		if book.ID == bookID {
			service.Books[i].Status = "Available"
			for _, member := range service.Members {
				if member.ID == memberID {
					for k, borrowedBook := range member.BorrowedBooks {
						if borrowedBook.ID == bookID {
							member.BorrowedBooks = append(member.BorrowedBooks[:k], member.BorrowedBooks[k+1:]...)
							return nil
						}
					}
				}
			}
			return errors.New("Member not found")
		}
	}
	return errors.New("Book not found")
}

func (service LibraryService) ListAvailableBooks() []models.Book {

	var availableBooks []models.Book
	for _, book := range service.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (service LibraryService) ListBorrowedBooks(memberID int) []models.Book {

	for _, member := range service.Members {
		if member.ID == memberID {
			return member.BorrowedBooks
		}
	}
	return nil
}
