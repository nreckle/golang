package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool
var x, y = 1, 2

func main() {
	fmt.Println("Hello, Golang")
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(5, 6))
	fmt.Println(swap("Second", "First"))
	fmt.Println(split(7))
	fmt.Println(split2(911))
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(x, y)

	var cc, dd, ee = true, false, "no!"
	fmt.Println(cc, dd, ee)

	// 短变量声明
	var l, m = 1, 2
	k := 3
	c2, python2, java2 := true, false, "no!"
	fmt.Println(l, m, k, c2, python2, java2)
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
