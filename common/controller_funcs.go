package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/neoetheilred/crispy-pancake/common/pages"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func createRandomBook(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	newBook("Random title", "Random summary")
	w.WriteHeader(http.StatusOK)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	formData := map[string]string{}

	json.NewDecoder(r.Body).Decode(&formData)

	newBook(formData["title"], formData["summary"])
}

func updateBook(w http.ResponseWriter, r *http.Request) {
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
}

func showBooksPage(w http.ResponseWriter, r *http.Request) {
	// pages.PageWrapper(w, "Hello")
	pages.PageBookList(w, booksToArray())
}

func createBookPage(w http.ResponseWriter, r *http.Request) {
	pages.PageCreateBook(w, struct{}{})
}
