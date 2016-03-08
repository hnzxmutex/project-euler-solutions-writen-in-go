package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := int64(0), int64(0)
	for i := int64(1); i < 101; i++ {
		a += int64(math.Pow(float64(i), 2))
		b += i
	}
	fmt.Println(int64(math.Pow(float64(b), 2)) - a)
}
