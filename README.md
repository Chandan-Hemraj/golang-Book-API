# BookAPI

<p align="center">
  <img src="https://user-images.githubusercontent.com/87279692/234941393-08d8f2f9-dd05-4982-bbf4-d0bd014a312b.jpg" height="350px" alt="">
</p>

This is a simple CRUD API for managing books built using Golang and Goroutines.

## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)

## Description

The BookAPI is a Go application that provides endpoints for managing books. It uses Goroutines to handle concurrent requests efficiently. The API supports basic CRUD operations, including adding a book, retrieving all books, and deleting a book.

## Installation

To install and run the BookAPI locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/Chandan-Hemraj/golang-Book-API
   cd BookAPI
   go run main.go
  The BookAPI will be accessible at http://localhost:8080.
    
## Usage

The BookAPI provides the following endpoints:

   1. POST /books: Add a new book to the collection.
   2. GET /books: Retrieve all books from the collection.
   3. DELETE /books?id=<bookID>: Delete a book from the collection by ID.
   
You can use tools like cURL or Postman to interact with the API.

## Testing

The BookAPI includes unit tests to ensure its functionality. To run the tests, execute the following command:
   ```bash
   go test ./...
   
<<<<<<< HEAD
This will run all the tests in the project and provide the test results.
=======
This will run all the tests in the project and provide the test results.
>>>>>>> 56d218ce1739f6d9b2be092e5a82b6afb4dbb47f
