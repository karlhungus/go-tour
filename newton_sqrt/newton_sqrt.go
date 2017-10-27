package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < int(x); i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

func main() {
	for z := 1; z < 11; z++ {
		fmt.Println(z, " math", math.Sqrt(float64(z)), " mine ", Sqrt(float64(z)))
	}
}
