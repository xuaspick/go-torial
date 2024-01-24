package main

import "math"

type Coordenada struct {
	X, Y float64
}

func (c *Coordenada) Scale(f float64) { //? struct Coordenada has a method called Scale
	c.X = c.X * f
	c.Y = c.Y * f
}

func (c Coordenada) ScaleWithoutPointer(f float64) { //? struct Coordenada has a method called Scale
	c.X = c.X * f
	c.Y = c.Y * f
}

// ? struct Coordenada has a method called Abs
func (c *Coordenada) Abs() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}
