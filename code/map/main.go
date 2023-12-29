package main

import "fmt"

type Student struct {
	Id   string
	Name string
}

func main() {
	m := make(map[string]*Student)

	a := m["hello"]

	fmt.Printf("%+v", a)
}
