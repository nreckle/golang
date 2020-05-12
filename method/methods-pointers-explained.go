package main

import (
	"fmt"
	"math"
)

type Vertexxx struct {
	X, Y float64
}

func Abs(v Vertexxx) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertexxx, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertexxx{3, 4}
	Scale(&v, 10)
	fmt.Print(Abs(v))
}
