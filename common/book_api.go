package common

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"sync/atomic"
)

type Book struct {
	ID      int64
	Title   string
	Summary string
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetSummary() string {
	return b.Summary
}

var books = map[int64]Book{}
var booksMu sync.RWMutex
var idCounter atomic.Int64

func StartBookApi() {
	http.HandleFunc("/api/books/all", getAllBooks)

	http.HandleFunc("/api/books/createRandom", createRandomBook)

	http.HandleFunc("/api/books/create", createBook)

	http.HandleFunc("/api/books/update", updateBook)

	http.HandleFunc("/home", showBooksPage)

	http.HandleFunc("/create", createBookPage)

	port := ":8080"
	log.Printf("Running on port %s", port)
	http.ListenAndServe(port, nil)
}

func booksToArray() []Book {
	booksMu.Lock()
	defer booksMu.Unlock()
	res := []Book{}
	for _, v := range books {
		res = append(res, v)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})

	return res
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
