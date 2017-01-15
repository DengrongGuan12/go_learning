package mymath2

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return 10.0
}

// 大写的才能在外面调
func Adda(a string) {
	fmt.Println(test(a))
}
