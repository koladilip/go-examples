package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func incAge(p person) person {
	p.age++
	return p
}

func main() {
	var p = person{name: "Dilip", age: 32}
	fmt.Println("before", p)
	fmt.Println("retuned", incAge(p))
	fmt.Println("after", p)
}
