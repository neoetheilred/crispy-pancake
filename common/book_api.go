package common

import (
	"log"
	"net/http"

	"github.com/neoetheilred/crispy-pancake/common/storage"
)

type Book struct {
	ID      int64
	Title   string
	Summary string
}

func (b *Book) GetID() int64 {
	return b.ID
}

func (b *Book) SetID(id int64) {
	b.ID = id
}

type User struct {
	ID       int64
	Name     string
	Password string
}

func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetSummary() string {
	return b.Summary
}

// var books = map[int64]Book{}
// var booksMu sync.RWMutex
// var idCounter atomic.Int64
var books = storage.NewStorage[*Book]()

func StartBookApi() {
	http.HandleFunc("/api/books/all", getAllBooks)

	http.HandleFunc("/api/books/createRandom", createRandomBook)

	http.HandleFunc("/api/books/create", createBook)

	http.HandleFunc("/api/books/update", updateBook)

	http.HandleFunc("/api/books/delete", deleteBook)

	http.HandleFunc("/home", showBooksPage)

	http.HandleFunc("/create", createBookPage)

	port := ":8080"
	log.Printf("Running on port %s", port)
	http.ListenAndServe(port, nil)
}

// func bookExists(id int64) bool {
// 	booksMu.Lock()
// 	defer booksMu.Unlock()
// 	_, ok := books[id]
// 	return ok
// }

// func deleteBookById(id int64) {
// 	booksMu.Lock()
// 	defer booksMu.Unlock()
// 	if _, ok := books[id]; ok {
// 		delete(books, id)
// 	}
// }

// func booksToArray() []Book {
// 	booksMu.Lock()
// 	defer booksMu.Unlock()
// 	res := []Book{}
// 	for _, v := range books {
// 		res = append(res, v)
// 	}

// 	sort.Slice(res, func(i, j int) bool {
// 		return res[i].ID < res[j].ID
// 	})

// 	return res
// }

// func newBook(title, summary string) {
// 	booksMu.Lock()
// 	defer booksMu.Unlock()

// 	idCounter.Add(1)
// 	books[idCounter.Load()] = Book{
// 		ID:      idCounter.Load(),
// 		Title:   title,
// 		Summary: summary,
// 	}
// }

// func edit(id int64, book Book) error {
// 	booksMu.Lock()
// 	defer booksMu.Unlock()

// 	if _, ok := books[id]; !ok {
// 		return fmt.Errorf("Book with id %d does not exist", id)
// 	}
// 	books[id] = book
// 	return nil
// }
