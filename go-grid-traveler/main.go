package main

import (
	"fmt"
)

func main() {
	fmt.Printf("gridTraveler(1, 1) = %s", gridTraveler(1, 1))
	fmt.Printf("gridTraveler(1, 1) = %s", gridTraveler(2, 3))
	fmt.Printf("gridTraveler(1, 1) = %s", gridTraveler(3, 2))
	fmt.Printf("gridTraveler(1, 1) = %s", gridTraveler(3, 3))
	fmt.Printf("gridTraveler(1, 1) = %s", gridTraveler(18, 18))
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