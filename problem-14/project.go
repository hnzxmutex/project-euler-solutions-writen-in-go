package main

import (
	"fmt"
)

const max = 1000000

func calCollatz(n int, result *[]int) int {
	if n == 1 {
		(*result)[n] = 1
		return 1
	}
	if n >= max {
		if n%2 == 0 {
			return 1 + calCollatz(n/2, result)
		} else {
			return 1 + calCollatz(3*n+1, result)
		}
	}
	if (*result)[n] != 0 {
		return (*result)[n]
	}
	if n%2 == 0 {
		(*result)[n] = 1 + calCollatz(n/2, result)
	} else {
		(*result)[n] = 1 + calCollatz(3*n+1, result)
	}
	return (*result)[n]
}

func printLink(i int) {
	if i == 1 {
		fmt.Printf("1\n")
		return
	}
	fmt.Printf("%d->", i)
	if i%2 == 0 {
		printLink(i / 2)
	} else {
		printLink(i*3 + 1)
	}
}

func main() {
	result := make([]int, (max + 1))
	num, count := 0, 0
	for i := 1; i <= max; i++ {
		tmp := calCollatz(i, &result)
		if tmp > count {
			count = tmp
			num = i
		}
	}
	printLink(869889)
	fmt.Println(calCollatz(13, &result))
	fmt.Println(count, num)
}
