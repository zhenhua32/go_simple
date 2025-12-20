package csvformat

import (
	"bytes"
	"encoding/csv"
	"io"
)

type Book struct {
	Author string
	Title  string
}
type Books []Book

func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()
	return nil
}

func WriteCSVOutput() error {
	b := Books{
		Book{Author: "Herman Melville", Title: "Moby Dick"},
		Book{Author: "F. Scott Fitzgerald", Title: "The Great Gatsby"},
		Book{Author: "Jane Austen", Title: "Pride and Prejudice"},
	}
	return b.ToCSV(io.Discard)
}

func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{Author: "Herman Melville", Title: "Moby Dick"},
		Book{Author: "F. Scott Fitzgerald", Title: "The Great Gatsby"},
		Book{Author: "Jane Austen", Title: "Pride and Prejudice"},
	}
	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
