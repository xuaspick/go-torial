package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func introToStructs() {
	fmt.Println("\n## Structs in Go ##")
	human := person{"Bob", 20}
	fmt.Println(human)
	human.name = "Alice"
	fmt.Println(human)

	pHuman := &human     //? pHuman is a pointer to human
	pHuman.name = "John" //? (*pHuman).name = "John" <- this is the same
	fmt.Println(human)
}

func structDefaults() {
	fmt.Println("\n## Structs Defaults in Go ##")
	var (
		human1 = person{"Bob", 20}
		human2 = person{name: "Alice"}
		human3 = person{age: 45}
		human4 = &person{"John", 35}
	)
	fmt.Println(human1, human2, human3, human4)
}

func introToArrays() {
	fmt.Println("\n## Arrays in Go ##")
	var a [2]string //? arrays have a fixed size
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4] //? slices are more common than arrays since they are more flexible
	fmt.Println(s)
}
