package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// struct to represent a book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// map to store books
var books = make(map[string]Book)

// channels to send response and error messages
var respChan = make(chan []byte)
var errChan = make(chan []byte)

// main function
func main() {
	http.HandleFunc("/books", handleBooks)

	log.Println("Server started on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

// handleBooks function to handle all requests
func handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		go getAllBooks(respChan, errChan)
		sendResponse(w)
	case http.MethodPost:
		go addBook(r, respChan, errChan)
		sendResponse(w)
	case http.MethodDelete:
		go deleteBook(r, respChan, errChan)
		sendResponse(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// addBook function to add a book to the map
func addBook(r *http.Request, respChan, errChan chan<- []byte) {
	log.Println("Got a request to add a book")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("Error decoding request:", err)
		errChan <- []byte(err.Error())
		return
	}

	books[book.ID] = book
	response, err := json.Marshal(books[book.ID])
	if err != nil {
		log.Println("Error marshaling response:", err)
		errChan <- []byte(err.Error())
		return
	}

	log.Println("Book added successfully")
	respChan <- response
}

// getAllBooks function to get all books from the map
func getAllBooks(respChan, errChan chan<- []byte) {
	log.Println("Got a request to get all books")
	response, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling response:", err)
		errChan <- []byte(err.Error())
		return
	}

	log.Println("Fetched all the books successfully")
	respChan <- response
}

// deleteBook function to delete a book from the map
func deleteBook(r *http.Request, respChan, errChan chan<- []byte) {
	log.Println("Got a request to delete a book")
	id := r.URL.Query().Get("id")

	delete(books, id)

	response, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling response:", err)
		errChan <- []byte(err.Error())
		return
	}

	log.Println("Book deleted successfully")
	respChan <- response
}

// sendResponse function to send response to the client
func sendResponse(w http.ResponseWriter) {
	select {
	case resp := <-respChan:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	case err := <-errChan:
		http.Error(w, string(err), http.StatusInternalServerError)
	}
}
