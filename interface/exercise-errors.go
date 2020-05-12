package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("errorcode : %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -10, ErrNegativeSqrt(-19)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}