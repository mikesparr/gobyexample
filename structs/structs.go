package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) // infers values given signature match
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"}) // age 0
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age) // 50
	sp.age = 51
	fmt.Println(s.age) // 51
}