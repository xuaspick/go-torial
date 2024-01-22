package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func introToSlices() {
	//? Slices work as a pointer to an existing array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] //? slice with elements 0 and 1
	b := names[1:3] //? slice with elements 1 and 2

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func slicesLiterals() {
	fmt.Print("\n## Slices Literals ##\n")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{i: 3},
		{b: true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
	type leType = struct {
		i int
		b bool
	}
	s = append(s, leType{17, false})
	fmt.Println(s)
	//? Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
}

func slicingSlices() {
	fmt.Print("\n## Slicing Slices ##\n")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[1:4]
	printSlice(s)
	s = s[:2]
	printSlice(s)
	s = s[1:]
	printSlice(s)
	// s = s[:9]  //! panic: runtime error: slice bounds out of range [:9] with capacity 5
	// printSlice(s)
}

func makeForSlices() {
	fmt.Print("\n## Make for Slices ##\n")
	a := make([]int, 5) //? len(a)=5
	printSlice(a)
	b := make([]int, 0, 5) //? len(b)=0, cap(b)=5
	printSlice(b)
	c := b[:2] //? len(c)=2, cap(c)=5
	printSlice(c)
	d := c[2:5] //? len(d)=3, cap(d)=3
	printSlice(d)
}

func slicesOfSlices() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	var s []int
	printSlice(s)
	s = append(s, 0) //? append works on nil slices.
	printSlice(s)
	s = append(s, 1, 2, 3, 4, 5) //? Can append more than one element at a time.
	printSlice(s)
}

func rangeForSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow { //? range returns two values: the index and the value.
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func rangeForSlice2() {
	//? You can skip the index or value by assigning to _.
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //? == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
