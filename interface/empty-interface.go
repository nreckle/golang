package main

import "fmt"

func main() {
	var i interface{}
	describeEmptyInterface(i)

	i = 42
	describeEmptyInterface(i)

	i = "hello"
	describeEmptyInterface(i)
}

func describeEmptyInterface(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}