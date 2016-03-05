package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(str string) bool {
	arr := []byte(str)
	length := int(len(arr))
	if 1 == (length % 2) {
		return false
	}
	flag := true
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-i-1] {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	var result uint32
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			tmp := i * j
			if isPalindrome(strconv.Itoa(tmp)) {
				fmt.Println(i, j, tmp)
				tmp := uint32(tmp)
				if tmp > result {
					result = tmp
				}
			}
		}
	}
	fmt.Println(result)
}
