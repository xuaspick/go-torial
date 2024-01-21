package main

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
} //? When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

func swap(x, y string) (string, string) {
	return y, x
} //? A function can return any number of results.

func split(sum int) (x, y int) { //? Named return values
	x = sum * 4 / 9
	y = sum - x
	return //? A return statement without arguments returns the named return values. This is known as a "naked" return.
}

func main() {
	//? Add
	fmt.Println(add(42, 13))

	//? Swap
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	a, b = swap("tu", "vieja")
	fmt.Println(a, b)

	//? Split
	fmt.Println(split(17))
	fmt.Println(split(32))

}
