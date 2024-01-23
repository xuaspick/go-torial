package main

import (
	"fmt"
	"strings"
)

type coordenadas struct {
	Lat, Lon float64
}

var m map[string]coordenadas

func introToMaps() {
	fmt.Printf("\n## Maps in Go ##\n")
	m = make(map[string]coordenadas)
	m["Bell Labs"] = coordenadas{
		40.68433, -74.39967,
	}
	fmt.Printf("The coordinates of Bell Labs are: %v\n", m["Bell Labs"])

	m["Google"] = coordenadas{
		37.42202, -122.08408,
	}
	fmt.Printf("The latitude of Google is: %v\n", m["Google"].Lat)
}

func mutatingMaps() {
	fmt.Printf("\n## Mutating Maps in Go ##\n")
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"] //? ok is a bool that indicates if the key was present in the map, v is the default value of the type
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	counterMap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counterMap[word]++
	}
	return counterMap
}
