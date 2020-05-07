package main

import "fmt"

func fibonacci2() func() int {
	count, sum1, sum2 := 0, 1, 1

    return func() int {
    	if count == 0 {
    		count++
		} else if count == 1 {
    		count++
		} else {
    		sum1, sum2 = sum2, sum1 + sum2
			//temp := sum2
    		//sum2 = sum1 + sum2
    		//sum1 = temp
		}

		return sum2
	}
}

func main() {
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
