package pages

import "io"

func initCreateBook() {
	initGenericPage("add_book.gohtml")
}

func PageCreateBook(w io.Writer, data any) {
	executeGenericPage("add_book.gohtml", w, data)
}
