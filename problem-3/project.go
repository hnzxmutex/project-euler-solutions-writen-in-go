package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

var num uint64 = 600851475143
var result uint64
var flag bool = true

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

//I just try it in serval goroutine,but it seem be low performance :(
func checkPrimeFactor(i uint64, c chan bool) {
	if num%i == 0 && IsPrime(i) {
		flag = false
		result = i
	}
	c <- true
}

func checkPrimeFactor2(i uint64) bool {
	return num%i == 0 && IsPrime(i)
}

func solution() {
	np := runtime.NumCPU()
	runtime.GOMAXPROCS(np)
	c := make(chan bool, np)
	for i := uint64(math.Sqrt(float64(num))); flag && i > 0; i-- {
		go checkPrimeFactor(i, c)
		<-c
	}
	fmt.Println(result)
}

func solution2() {
	for i := uint64(math.Sqrt(float64(num))); flag && i > 0; i-- {
		if checkPrimeFactor2(i) {
			result = i
			break
		}
	}
	fmt.Println(result)
}

func main() {
	t1 := time.Now().UnixNano()
	solution2()
	t2 := time.Now().UnixNano()
	t := float64(t2-t1) / 1000000000
	fmt.Println(t)
}
