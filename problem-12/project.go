package main

import (
	"fmt"
	"math"
)

func getDivisors(n int) int {
	var count int = 0
	root := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 1; i <= root; i++ {
		if n%i == 0 {
			count++
		}
	}
	count *= 2
	if root*root == n {
		count--
	}
	fmt.Printf("%d result = %d\n", n, count)
	return count
}

func main() {
	i := 1
	result := 0
	for {
		result += i
		i++
		if getDivisors(result) >= 500 {
			break
		}

	}
	fmt.Println(result)
}
