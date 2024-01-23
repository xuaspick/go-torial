package main

import "fmt"

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
