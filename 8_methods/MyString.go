package main

import "fmt"

type MyString struct {
	S string
}

func (ms *MyString) M() {
	fmt.Println(ms.S)
}

func (ms *MyString) get() string {
	return ms.S
}
