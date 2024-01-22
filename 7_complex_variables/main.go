package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("## Pointers in Go ##")
	// introToPointers()
	introToStructs()
	structDefaults()
}

func introToPointers() {
	señora := "juana"
	copiaSeñora := señora
	punteroSeñora := &señora

	fmt.Print("\nEstado inicial señora: juana\n")
	fmt.Printf("Direccion punteroSeñora: %v\n", punteroSeñora)
	fmt.Printf("señora: %v, copiaSeñora: %v, punteroSeñora: %v\n", señora, copiaSeñora, *punteroSeñora)

	fmt.Print("\nCambio señora a: tu vieja\n")
	señora = "tu vieja"
	fmt.Printf("señora: %v, copiaSeñora: %v, punteroSeñora: %v\n", señora, copiaSeñora, *punteroSeñora)

	fmt.Print("\nCambio puntero señora a: tu abuela\n")
	*punteroSeñora = "tu abuela"
	fmt.Printf("señora: %v, copiaSeñora: %v, punteroSeñora: %v\n", señora, copiaSeñora, *punteroSeñora)
}

func introToStructs() {
	human := person{"Bob", 20}
	fmt.Println(human)
	human.name = "Alice"
	fmt.Println(human)

	pHuman := &human     //? pHuman is a pointer to human
	pHuman.name = "John" //? (*pHuman).name = "John" <- this is the same
	fmt.Println(human)
}

func structDefaults() {
	var (
		human1 = person{"Bob", 20}
		human2 = person{name: "Alice"}
		human3 = person{age: 45}
		human4 = &person{"John", 35}
	)
	fmt.Println(human1, human2, human3, human4)
}
