package main

import (
	"fmt"
)

func introToPointers() {
	fmt.Println("\n## Pointers in Go ##")
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
