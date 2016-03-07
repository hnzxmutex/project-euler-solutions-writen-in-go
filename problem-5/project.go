package main

import (
	"fmt"
)

func isValid(num int) bool {
	flag := true
	for i := 1; i < 21; i++ {
		if 0 != num%i {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	for i := 1; ; i++ {
		if isValid(i) {
			fmt.Println(i)
			return
		}
	}
}
