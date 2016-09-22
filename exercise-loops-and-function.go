package main

import "fmt"
import "math"

// func Sqrt(x float64) float64 {
// 	z := 1.0
// 	for i := 0; i < 10; i++ {
// 		z = z - (z*z-x)/(2*z)
// 	}
// 	return z
// }

func Sqrt(x float64) (float64, int) {
	z := 1.0
	zprev := 0.0
	i := 0
	for ; math.Abs(zprev-z) > 0.0000001; i++ {
		zprev = z
		z = z - (z*z-x)/(2*z)
	}
	return z, i
}

func main() {
	fmt.Println(0.0 == 1.0)
	fmt.Println(Sqrt(6))
}
