package main

import (
	"fmt"
	"io"
	"math"
	"strings"
)

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func introToMethods() {
	c := Coordenada{3, 4}
	fmt.Println("The method in Coordenada returns: ", c.Abs())
}

func nonStructMethod() {
	f := MyFloat(-math.Sqrt2)                                                          //? f is a variable of type MyFloat
	fmt.Printf("This nonstruct number (%v) has a method to print as %v\n", f, f.Abs()) //? f.Abs() is a method call
}

func methodsPointers() {
	c := Coordenada{3, 4}
	fmt.Println("Initial state", c.Abs())
	c.ScaleWithoutPointer(10)
	fmt.Println("Scale 10 without pointer: ", c.Abs())
	c.Scale(10)
	fmt.Println("Scale 10 with pointer: ", c.Abs())
}

func introToInterfaces() {
	type Abser interface { //? interface Abser has a method called Abs
		Abs() float64
	}
	var a Abser //? Abser: anything able to call Abs()
	f := MyFloat(-math.Sqrt2)
	c := Coordenada{3, 4}

	a = f  //? a MyFloat implements Abser
	a = &c //? a *Vertex implements Abser

	//! In the following line, c is a Coordenada (not *Coordenada)
	//! and does NOT implement Abser.
	// a = c //! This line will throw an error
	//? To print, we need any item capable of calling Abs()
	fmt.Println("Abs implemented through interface", a.Abs())
}

func interfaceSatisfiedImplicitly() {
	var i I = &MyString{"hello"} //? T implements I
	i.M()
}

func interfaceValues() {
	fmt.Println("## Interface values ##")
	var i I
	describe(i) //? i is nil, so it has no value or type
	// i.M() //! This line will throw an error, because i is nil so no types nor methods

	i = &MyString{"hello"}
	describe(i)
	// describeEmptyInterface(i) //? this can be done because i is an interface
	i.M() //? i implements I with MyString struct

	i = MyFloat(math.Pi)
	describe(i)
	i.M() //? i implements I with MyFloat type
}

func emptyInterface() {
	fmt.Println("## Empty interface ##")
	var i interface{}
	describeEmptyInterface(i)

	i = 42
	describeEmptyInterface(i)

	i = "hello"
	describeEmptyInterface(i)
}

func typeAssertions() {
	fmt.Printf("\n## Type assertions ##\n")
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //? This can be done because we are getting the value of ok
	fmt.Println(f, ok)

	// f = i.(float64) //! panic, because we are not getting the value of ok
	// fmt.Println(f)
}

func typeSwitches() {
	fmt.Printf("\n## Type switches ##\n")
	interfaceTypeSwtich(21)
	interfaceTypeSwtich("hello")
	interfaceTypeSwtich(true)
}

func interfaceTypeSwtich(i interface{}) {
	switch v := i.(type) { //? v is the type of i
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)

	}
}

func introToStringers() {
	fmt.Printf("\n## Stringers ##\n")
	//? fmt package looks for a method called String in the struct
	//? if it finds it, it will print the result of that method when printing the struct
	a := Person{"Arthur Dent", 42} //? struct Person has a method called String
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

func introToErrors() {
	fmt.Printf("\n## Intro to errors ##\n")
	//? error is a built-in interface
	if err := runError(); err != nil {
		fmt.Println(err)
	}
}

func introToReaders() {
	fmt.Printf("\n## Intro to readers ##\n")
	//?
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	fmt.Println(r)
	fmt.Println(b)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
