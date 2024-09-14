package person

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Info() {
	fmt.Printf("name: %s, age: %d\n", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}
