package book

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Year: %d", b.Title, b.Author, b.Year)
}
