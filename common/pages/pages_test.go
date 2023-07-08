package pages

import (
	"bytes"
	"log"
	"testing"
)

func TestCreateBook() {

}

func TestBookList(t *testing.T) {
	testCase := []struct {
		Title   string
		Summary string
	}{{Title: "Title", Summary: "Summary"}}
	expected := `
<html>
    <head></head>

    <body>
        <header>
            <h1>Books</h1>
            <p><a href="/home">Home</a></p>
            <p><a href="/create">New book</a></p>
        </header>


    <div>
        <h3>Title</h3>
        <p>Summary</p>
    </div>


    </body>
</html>`
	var buf bytes.Buffer
	PageBookList(&buf, testCase)
	if buf.String() != expected {
		diff(expected, buf.String())
		t.Errorf("Expected: %s\nGot: %s", expected, buf.String())
		// t.Fail()
	}
}

func diff(s1 string, s2 string) {
	if len(s1) != len(s2) {
		log.Printf("strings lengths are different, %d vs %d\n", len(s1), len(s2))
	}

	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			log.Printf("at pos %d: %c vs %c\n", i, s1[i], s2[i])
		}
	}
}
