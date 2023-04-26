package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUDOperations(t *testing.T) {
	// Initialize the books map
	books = make(map[string]Book)

	// Initialize the server and the test client
	server := httptest.NewServer(http.HandlerFunc(handleBooks))
	defer server.Close()

	client := server.Client()

	// Test case: Add a book
	testAddBook(t, server.URL, client)

	// Test case: Get all books
	testGetAllBooks(t, server.URL, client)

	// Test case: Delete a book
	testDeleteBook(t, server.URL, client)

	// Verify that the book is deleted
	testVerifyBookDeleted(t, server.URL, client)
}

func testAddBook(t *testing.T, serverURL string, client *http.Client) {
	book := Book{
		ID:     "1",
		Title:  "Golang for Dummies",
		Author: "John Doe",
	}
	payload, _ := json.Marshal(book)
	req, err := http.NewRequest(http.MethodPost, serverURL+"/books", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Add book failed. Expected status code %v, but got %v", http.StatusCreated, resp.StatusCode)
	}
}

func testGetAllBooks(t *testing.T, serverURL string, client *http.Client) {
	req, err := http.NewRequest(http.MethodGet, serverURL+"/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Get all books failed. Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
	}
	var booksResponse map[string]Book
	err = json.NewDecoder(resp.Body).Decode(&booksResponse)
	if err != nil {
		t.Fatal(err)
	}
	if len(booksResponse) != 1 {
		t.Errorf("Get all books failed. Expected 1 book, but got %v", len(booksResponse))
	}
	if booksResponse["1"].ID != "1" || booksResponse["1"].Title != "Golang for Dummies" || booksResponse["1"].Author != "John Doe" {
		t.Errorf("Get all books failed. Incorrect book data")
	}
}

func testDeleteBook(t *testing.T, serverURL string, client *http.Client) {
	req, err := http.NewRequest(http.MethodDelete, serverURL+"/books?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Delete book failed. Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
	}
}

func testVerifyBookDeleted(t *testing.T, serverURL string, client *http.Client) {
	req, err := http.NewRequest(http.MethodGet, serverURL+"/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Delete book failed. Expected status code %v, but got %v", http.StatusOK, resp.StatusCode)
	}
	var booksResponse2 map[string]Book
	err = json.NewDecoder(resp.Body).Decode(&booksResponse2)
	if err != nil {
		t.Fatal(err)
	}
	if len(booksResponse2) != 0 {
		t.Errorf("Delete book failed. Book still exists")
	}
}
