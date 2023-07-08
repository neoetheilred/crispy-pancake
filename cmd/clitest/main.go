package main

import (
	"os"

	"github.com/neoetheilred/crispy-pancake/common/pages"
)

func main() {
	// pages.PageBookList(os.Stdout, []common.Book{{Title: "Book", Summary: "Something"}})

	pages.PageCreateBook(os.Stdout, struct{}{})
}
