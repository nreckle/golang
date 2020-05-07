package main

import "fmt"

type Vertexx struct {
	Lat, Long float64
}

var mm = map[string]Vertexx{
	"Bell Labs": Vertexx{
		40.68433, -74.39967,
	},
	"Google": Vertexx{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(mm)
}
