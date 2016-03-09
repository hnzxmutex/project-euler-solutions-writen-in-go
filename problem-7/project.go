package main

import (
	"fmt"
	"math"
)

func IsPrime(n uint64) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n&1 == 0 {
		return false
	}
	s := uint64(math.Sqrt(float64(n)))
	for i := uint64(3); i <= s; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	i, num := uint64(1), 0
	fmt.Println(num, i)
	for num < 10001 {
		i++
		if IsPrime(i) {
			num++
		}
	}
	fmt.Println(num, i)
}
