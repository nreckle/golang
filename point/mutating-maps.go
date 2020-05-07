package main

import "fmt"

func main() {
	mvp := make(map[string]int)

	mvp["Answer"] = 42
	fmt.Println("The value:", mvp["Answer"])

	mvp["Answer"] = 48
	fmt.Println("The value:", mvp["Answer"])

	delete(mvp, "Answer")
	fmt.Println("The value:", mvp["Answer"])

	v := 0
	ok := false
	v, ok = mvp["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}