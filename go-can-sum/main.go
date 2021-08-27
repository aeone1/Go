package main

import (
	"fmt"
	"time"
)

/*
Return true or false if its possible to generate the target sum
using numbers from given array.
 - We can use the elements of the array as many times as we want
 - We can assume that all input numbers are nonnegative
*/

func main() {
	defer timeTrack(time.Now(), "canSum")
	//fmt.Printf("canSum(7, []int{5, 3, 4, 7}) = %v\n", canSum(7, []int{5, 3, 4, 7})) //true
	//fmt.Printf("canSum(7, []int{2, 4}) = %v\n", canSum(7, []int{2, 4})) //false
	//fmt.Printf("canSum(8, []int{2, 3, 5}) = %v\n", canSum(8, []int{2, 3, 5})) //true
	// fmt.Printf("canSum(300, []int{7, 14}) = %v\n", canSum(300, []int{7, 14})) //false

	// fmt.Printf("canSumMemo(7, []int{2, 4}) = %v\n", canSumMemo(7, []int{2, 4}, make(map[int]bool))) //false
	// fmt.Printf("canSumMemo(8, []int{2, 3, 5}) = %v\n", canSumMemo(8, []int{2, 3, 5}, make(map[int]bool))) //true
	fmt.Printf("canSumMemo(300, []int{7, 14}) = %v\n", canSumMemo(300, []int{7, 14}, make(map[int]bool))) //false
}

func canSum(targetSum int, numbers []int) bool { 
	// O(n^m) time complexity and O(m) space complexity
	// n - array length, m - target sum
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers { // index, value
		remainder := targetSum - num
		if canSum(remainder, numbers) {
			return true
		}
		//fmt.Print(num)
	}
	return false
}

func canSumMemo(targetSum int, numbers []int, memo map[int]bool) bool { 
	// O(n * m) time complexity and O(m) space complexity
	// n - array length, m - target sum

	fmt.Printf("target sum: %v  memo: %v\n", targetSum, memo)

	if val, ok := memo[targetSum]; ok {
		return val
	}

	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, num := range numbers { // index, value
		remainder := targetSum - num
		if canSumMemo(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
		//fmt.Print(num)
	}
	memo[targetSum] = false
	return false
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
