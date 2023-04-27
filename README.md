# BookAPI

This is a simple CRUD API for managing books built using Golang and Goroutines.

## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

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

 This will run all the tests in the project and provide the test results.

## Contributing
   
Contributions to the BookAPI project are welcome! If you find any issues or have suggestions for improvement, please create an issue or submit a pull request.

Before contributing, please review the CONTRIBUTING.md file for guidelines.
