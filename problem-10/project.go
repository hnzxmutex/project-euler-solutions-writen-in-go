package main

import (
	"fmt"
	"math"
)

const max = 2000000

func main() {
	var result int
	primeArray := make([]bool, max)
	for i := range primeArray {
		primeArray[i] = true
	}
	primeArray[0] = false
	root := int(math.Floor(math.Sqrt(float64(max))))
	for i := 2; i <= root; i++ {
		if primeArray[i] {
			for j := i * i; j < max; j += i {
				if primeArray[j] {
					primeArray[j] = false
				}
			}
		}
	}
	primeArray[0] = false
	primeArray[1] = false
	// count := 0
	for i, isPrime := range primeArray {
		if isPrime {
			// count++
			// fmt.Println(count, i)
			// if count > 250 {
			// 	fmt.Println(count, result)
			// 	return
			// }
			result += i
		}
	}
	fmt.Println(result)
}
