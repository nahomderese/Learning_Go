# Library Management System Documentation

## Overview

This is a console-based Library Management System implemented in Go. It allows users to manage books and members in a library. Users can add, remove, borrow, and return books, as well as list available and borrowed books.

## Components

### Models

- **Book**: Represents a book in the library.
- **Member**: Represents a member of the library.

### Services

- **LibraryService**: Implements the core functionalities of the library, such as adding, removing, borrowing, and returning books.

### Controllers

- **LibraryController**: Provides functions to interact with the library service through console input.

### Main

- **main.go**: Entry point of the application. Provides a menu for users to interact with the library system.

## Usage

### Running the Application

To run the application, execute the following command in your terminal:

## Functionality

1. Add a new book

   - Prompts for book ID, title, and author.
   - Adds the book to the library.

2. Remove an existing book

   - Prompts for book ID.
   - Removes the book from the library.

3. Borrow a book

   - Prompts for book ID and member ID.
   - Marks the book as borrowed by the member.

4. Return a book

   - Prompts for book ID and member ID.
   - Marks the book as returned by the member.

5. List all available books

   - Displays all books that are currently available in the library.

6. List all borrowed books by a member

   - Prompts for member ID.
   - Displays all books borrowed by the member.

7. Exit

   - Exits the application.
