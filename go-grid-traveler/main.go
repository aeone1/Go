package main

import (
	"fmt"
	"strconv"
)

/*
How many ways can I get from the top left
to the bottom right. 
*/

func main() {
	fmt.Printf("gridTraveler(1, 1) = %d\n", gridTraveler(1, 1)) //1
	//fmt.Printf("gridTraveler(3, 2) = %d\n", gridTraveler(3, 2)) //3
	//fmt.Printf("gridTraveler(2, 3) = %d\n", gridTraveler(2, 3)) //3
	//fmt.Printf("gridTraveler(3, 3) = %d\n", gridTraveler(3, 3)) //6
	//fmt.Printf("gridTraveler(18, 18) = %d\n", gridTraveler(18, 18)) //2333606220

	//fmt.Printf("gridTravelerMem(1, 17) = %d\n", gridTravelerMem(1, 17, make(map[string]int))) //1
	//fmt.Printf("gridTravelerMem(3, 2) = %d\n", gridTravelerMem(3, 2, make(map[string]int))) //3
	//fmt.Printf("gridTravelerMem(2, 3) = %d\n", gridTravelerMem(2, 3, make(map[string]int))) //3
	//fmt.Printf("gridTravelerMem(3, 3) = %d\n", gridTravelerMem(3, 3, make(map[string]int))) //6
	fmt.Printf("gridTravelerMem(18, 18) = %d\n", gridTravelerMem(18, 18, make(map[string]int))) //2333606220
}

func gridTraveler(m int, n int) int {
	if m == 1 && n == 1  {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

func gridTravelerMem(m int, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	keyReverse := strconv.Itoa(n) + "," + strconv.Itoa(m)
	fmt.Printf("key: %s\n", key)
	fmt.Printf("keyReverse: %s\n", keyReverse)

	if val, ok := memo[key]; ok {
		fmt.Printf("found key - val: %s - %d\n", key, val)
		return val
	}

	if val, ok := memo[keyReverse]; ok {
		fmt.Printf("found keyReverse - val: %s - %d\n", keyReverse, val)
		return val
	}

	if m == 1 && n == 1  {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = gridTravelerMem(m-1, n, memo) + gridTravelerMem(m, n-1, memo)
	return memo[key]
}