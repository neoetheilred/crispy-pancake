package common

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

type Book struct {
	ID      int64
	Title   string
	Summary string
}

var books = map[int64]Book{}
var booksMu sync.RWMutex
var idCounter atomic.Int64

func StartBookApi() {
	http.HandleFunc("/api/books/all", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(books)
	})

	http.HandleFunc("/api/books/createRandom", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		newBook("Random title", "Random summary")
		w.WriteHeader(http.StatusOK)
	})
	port := ":8080"
	log.Printf("Running on port %s", port)
	http.ListenAndServe(port, nil)
}

func newBook(title, summary string) {
	booksMu.Lock()
	defer booksMu.Unlock()

	idCounter.Add(1)
	books[idCounter.Load()] = Book{
		ID:      idCounter.Load(),
		Title:   title,
		Summary: summary,
	}
}
