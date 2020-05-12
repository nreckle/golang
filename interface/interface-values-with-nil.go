package main

import "fmt"

type I interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	var t *T2
	i = t
	describe(i)
	i.M()

	i = &T2{"hello, interface value with nil"}
	describe(i)
	i.M()
}