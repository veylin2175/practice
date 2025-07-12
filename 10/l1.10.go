package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	group := groupTemps(temps)

	for k, v := range group {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func groupTemps(temps []float64) map[int][]float64 {
	var gKey int
	group := make(map[int][]float64)

	for _, v := range temps {
		if v < 0 {
			gKey = int(math.Floor(v/10))*10 + 10
		} else {
			gKey = int(math.Floor(v/10)) * 10
		}

		group[gKey] = append(group[gKey], v)
	}

	return group
}
