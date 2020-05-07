package main

import "fmt"

type Verital struct {
	X int
	Y int
}

func main() {
	p := &(Verital{123, 231})
	fmt.Println((*p).X)
	fmt.Println((*p).Y)

	p.X = 1e9
	fmt.Println(*p)
}
