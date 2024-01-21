package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", float32(rand.Intn(100))*rand.Float32()/math.Pi)
}
