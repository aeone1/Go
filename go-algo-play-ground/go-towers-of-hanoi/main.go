package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "\nTowers of Hanoi")

	fmt.Print("TowersOfHanoi(3,1,3): \n")
	TowersOfHanoi(3,1,3)
}

func TowersOfHanoi(n int, start int, end int) {
	if n == 1 {
		PrintMove(start, end)
	} else {
		other := 6 - (start + end)
		TowersOfHanoi(n - 1, start, other)
		PrintMove(start, end)
		TowersOfHanoi(n - 1, other, end)
	}
}

func PrintMove(start int, end int) {
	fmt.Printf("%v --> %v\n", start, end)
}