package main

import "fmt"

var c, python, java bool //? A var declaration can include initializers, one per variable.

func main() {
	var i int //? A var declaration can include initializers, one per variable.
	fmt.Println(i, c, python, java)
	ExecuteTask() //? ExecuteTask is a function declared in tasks.go
}
