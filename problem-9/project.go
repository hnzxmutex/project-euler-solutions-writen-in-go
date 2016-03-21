package main

import (
	"fmt"
	"math"
)

func main() {
	var result int
totalFor:
	for a := 1; a < 500; a++ {
		for b := a; b <= 500; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			if a*a+b*b == c*c &&
				(c > b && b > a) &&
				a+b+c == 1000 {
				fmt.Println(a, b, c, a*b*c)
				result = c
				break totalFor
			}
		}
	}
	fmt.Println(result)
}
