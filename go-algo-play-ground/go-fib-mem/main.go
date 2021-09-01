package main

import (
	"fmt"
)

func main() {
	//fmt.Println(fib(6))
	fmt.Println(fib(7))
	//fmt.Println(fib(50))
	
	fmt.Println(fibMem(7, make(map[int]int)))
	
}

func fib (n int) int { // exp time complexity
	if n <= 2 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}

func fibMem (n int, memo map[int]int ) int { // Memoization
	if m, ok := memo[n]; ok {
		//fmt.Printf("found val for %v = %v \n", n, m)
		return m
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fibMem(n - 1, memo) + fibMem(n - 2, memo)
	fmt.Println(memo)
	return memo[n]
}


