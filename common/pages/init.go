package pages

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

const pathToPages = "./assets/pages/"
const wrapperPage = "wrapper.gohtml"
const debug = true

func init() {
	initBookList()
	initCreateBook()
}

var pages = map[string]*template.Template{}

func initGenericPage(path string) {
	pages[path] = template.Must(
		template.New(path).Parse(fmt.Sprintf(readFile(pathToPages+wrapperPage), readFile(pathToPages+path), path)),
	)
}

func readFile(path string) string {
	res, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func executeGenericPage(path string, w io.Writer, data any) {
	if debug {
		initGenericPage(path)
	}
	err := pages[path].ExecuteTemplate(w, wrapperPage, page{path, data})
	if err != nil {
		log.Println(err)
	}
}
