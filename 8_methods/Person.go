package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { //? struct Person has a method called String
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
