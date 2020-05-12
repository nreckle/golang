package main

import (
	"fmt"
	"math"
)

type II interface {
	M()
}

type TT struct {
	S string
}

func (t *TT) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i II) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
    var i II

    i = &TT{"Hello, TT"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()
}
