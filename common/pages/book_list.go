package pages

import "io"

// var booksList = template.New("bookList")

func initBookList() {
	// booksList = template.Must(template.ParseFiles("./common/pages/booklist.gohtml", "./common/pages/header.gohtml", "./common/pages/wrapper.gohtml"))
	// pages["booklist"] = template.Must(template.ParseFiles("./assets/pages/booklist.gohtml", "./assets/pages/wrapper.gohtml"))
	initGenericPage("booklist.gohtml")
}

func PageBookList(w io.Writer, data any) {
	executeGenericPage("booklist.gohtml", w, data)
}
