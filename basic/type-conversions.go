package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe2          = false
	MaxInt2 uint64 = 1<<64 - 1
	z2             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe2, ToBe2)
	fmt.Printf("Type: %T Value: %v\n", MaxInt2, MaxInt2)
	fmt.Printf("Type: %T Value: %v\n", z2, z2)

	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}
