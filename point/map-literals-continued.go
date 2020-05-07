package main

import "fmt"

type Vertexxx struct {
	Lat, Long float64
}

var mmm = map[string]Vertexxx{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(mmm)
}