package main

import "fmt"

type I2 interface {
	M()
}

func main() {
	var i I2
	describeI2(i)
	i.M()
}

func describeI2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}