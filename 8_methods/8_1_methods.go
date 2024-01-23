package main

import (
	"fmt"
	"math"
)

type Coordenada struct {
	X, Y float64
}

// ? struct Coordenada has a method called Abs
func (v Coordenada) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func introToMethods() {
	v := Coordenada{3, 4}
	fmt.Println(v.Abs())
}
