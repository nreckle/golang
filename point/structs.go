package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex2{1, 2})

	v := Vertex2{3, 4}
	v.X = 15
	fmt.Println(v.X)
}