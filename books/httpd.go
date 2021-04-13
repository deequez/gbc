package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db sync.Map // ISBN -> Book
	// db map[string]Book // Can't use - not goroutine safe
)

const (
	bookPrefix = "/book/"
	jsonCtype  = "application/json"
)

func main() {
	if err := healthCheck(); err != nil {
		log.Fatal(err)
	}

	// Routing
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/book", addHandler)
	http.HandleFunc("/book", getHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// GET /book/<isbn>
func getHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write code to return book, http.StatusNotFound if book not found
	//  - Use r.URL.Path to get the ISBN (trim the /book prefix)
	//  - Return http.StatusNotFound if book not found
	//  - Error is this is not a GET request
	isbn := r.URL.Path[len(bookPrefix):]
	log.Printf("isbn: %s\n", isbn)

	// Step 1: Validate & Unmarshal
	if r.Method != "GET" {
		http.Error(w, "this is not a GET request", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	// i, ok := db.Load(isbn)
	// if !ok {
	// 	http.Error(w, fmt.Sprintf("%q not found", isbn), http.StatusNotFound)
	// 	return
	// }

	// Step 2: Work
	book, err := getBook(isbn)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	// Step 3: Marshal response back to user
	resp := map[string]interface{}{
		"title":  book.Title,
		"author": book.Author,
		"isbn":   book.ISBN,
	}
	json.NewEncoder(w).Encode(resp)
}

func getBook(isbn string) (Book, error) {
	return Book{
		Title:  "The Colour of Magic",
		Author: "Terry Pratchett",
		ISBN:   "0062225677",
	}, nil
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Validate & Unmarshal
	if r.Method != "POST" {
		http.Error(w, "only POST", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		// TODO: Check that err.Error() does not contain sensitive information
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Validate book (e.g. empty ISBN)

	// Step 2: Work
	numBooks := addBook(book)

	// Step 3: Marshal the response back to user
	resp := map[string]interface{}{
		"count": numBooks,
		"isbn":  book.ISBN,
	}
	/*
		var resp struct {
			count int
			isbn string
		}
		resp.count = numBooks
		resp.isbn = book.ISBN
	*/
	w.Header().Set("Content-Type", jsonCtype)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		// Can't reset the HTTP status code here
		log.Printf("error encoding: %s", err)
	}
	// Other option for marshall
}

func addBook(book Book) int {
	db.Store(book.ISBN, book)

	var count int
	db.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

// var ok bool

func healthCheck() error {
	/*
		Example of failing health check
		ok = !ok
		if !ok {
			return fmt.Errorf("oops")
		}
	*/
	return nil // FIXME: Check connection to database
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := healthCheck(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "OK!\n")
}

type Book struct {
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
	Title  string `json:"title"`
}
