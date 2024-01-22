package main

import (
	"fmt"

	"rsc.io/quote" //? Task1 requires executing `go get rsc.io/quote` in the terminal.
)

// ? Task 1: Import the quote package and print a quote from it.
func task1() {
	fmt.Println(quote.Go())
}

// ? Task 2: use quotes from the quote package to print a quote from it.
func task2() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
