package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := healthCheck(); err != nil {
		log.Fatal(err)
	}

	// Routing
	http.HandleFunc("/health", healthHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

var ok bool

func healthCheck() error {
	ok = !ok
	if !ok {
		return fmt.Errorf("oops")
	}
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
	ISNB   string
	Author string
	Title  string
}
