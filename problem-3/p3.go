package main

import (
	"fmt"
	"math"
	"runtime"
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

func checkPrimeFactor(i uint64, c chan bool) {
	if i%200000 == 0 {
		fmt.Println(i)
	}
	if num%i == 0 && IsPrime(i) {
		flag = false
		result = i
	}
	c <- true
}

func getP4() uint64 {

	np := runtime.NumCPU()
	runtime.GOMAXPROCS(np)
	c := make(chan bool, np)
	for i := uint64(math.Sqrt(float64(num))); flag && i > 0; i-- {
		go checkPrimeFactor(i, c)
		<-c
	}
	return result
}
