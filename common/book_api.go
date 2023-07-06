package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	http.HandleFunc("/api/books/edit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		id, title, summary := r.Form["id"], r.Form["title"], r.Form["summary"]
		if len(id) == 0 || len(title) == 0 || len(summary) == 0 {
			json.NewEncoder(w).Encode([]string{"error: id, title and summary must be set"})
		}
		n, err := strconv.ParseInt(id[0], 10, 64)
		if err != nil {
			json.NewEncoder(w).Encode([]string{"id must be a valid int64"})
		}
		if err := edit(n, Book{n, title[0], summary[0]}); err != nil {
			json.NewEncoder(w).Encode([]string{fmt.Sprintf("error: %s", err.Error())})
		}

		json.NewEncoder(w).Encode([]string{"ok"})
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

func edit(id int64, book Book) error {
	booksMu.Lock()
	defer booksMu.Unlock()

	if _, ok := books[id]; !ok {
		return fmt.Errorf("Book with id %d does not exist", id)
	}
	books[id] = book
	return nil
}
