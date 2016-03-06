package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Adam", 21})
	fmt.Println(person{name: "Tom", age: 31})

	fmt.Println(&person{name: "Anne", age: 41})

	s := person{name: "Sean", age: 51}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 53
	fmt.Println(sp.age)
}
