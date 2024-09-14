package main

import (
	"fmt"
	"part1/book"
	"part1/person"
	"part1/shapes"
)

func main() {
	person := person.Person{Name: "John Doe", Age: 30}
	person.Info()

	person.Birthday()
	person.Info()

	circle := shapes.Circle{Radius: 5}
	rectangle := shapes.Rectangle{Height: 10, Width: 20}

	fmt.Printf("Circle Area: %f\n", circle.Area())
	fmt.Printf("Rectangle Area: %f\n", rectangle.Area())

	shps := []shapes.Shape{&circle, &rectangle}

	shapes.ShapeSlice(shps)

	book := book.Book{Title: "1984", Author: "George Orwell", Year: 1949}
	fmt.Println(book.String())
}
