package main

import (
	"fmt"
	"time"
)

/*
Return a slice of combinations that can 
add up to the given number
*/

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func HowSum(targetSum int, numbers []int) []int {
	/*
	m = target sum
	n = len(numbers)

	time: O(n^m * m) //* m for array(slice) coppying.
	space: O(m)
	*/
	if targetSum == 0 {
		return []int{0}
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := HowSum(remainder, numbers)

		if remainderResult != nil {
			return append(remainderResult, num) // this call will take m steps
		}
		
	}
	return nil
}

func HowSumMem(targetSum int, numbers []int, memo map[int][]int) []int {
	/*
	Memoized
	m = target sum
	n = len(numbers)

	time: O(n * m^2) //* m for array(slice) coppying.
	space: O(m^2)
	*/
	
	if val, ok := memo[targetSum]; ok {
		return val
	}

	if targetSum == 0 {
		return []int{0}
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remainderResult := HowSumMem(remainder, numbers, memo)

		if remainderResult != nil {
			memo[targetSum] = append(remainderResult, num)
			return memo[targetSum]
		}
		
	}
	memo[targetSum] = nil
	return memo[targetSum]
}

func main() {
	defer timeTrack(time.Now(), "HowSum")

	// fmt.Printf("HowSum(7, []int{2, 3}) = %v\n", HowSum(7, []int{2, 3})) //[3, 2, 2]
	// fmt.Printf("HowSum(7, []int{5, 3, 4, 7}) = %v\n", HowSum(7, []int{5, 3, 4, 7})) //[4, 3]
	// fmt.Printf("HowSum(7, []int{2, 4}) = %v\n", HowSum(7, []int{2, 4})) //nil
	// fmt.Printf("HowSum(8, []int{2, 3, 5}) = %v\n", HowSum(8, []int{2, 3, 5})) //[2, 2, 2, 2]
	// fmt.Printf("HowSum(8, []int{3, 5, 2}) = %v\n", HowSum(8, []int{3, 5, 2})) //[2, 3, 3]
	// fmt.Printf("HowSum(300, []int{7, 14}) = %v\n", HowSum(300, []int{7, 14})) //nil

	fmt.Printf("HowSumMem(7, []int{2, 3}) = %v\n", HowSumMem(7, []int{2, 3}, make(map[int][]int))) //[3, 2, 2]
	fmt.Printf("HowSumMem(7, []int{5, 3, 4, 7}) = %v\n", HowSumMem(7, []int{5, 3, 4, 7}, make(map[int][]int))) //[4, 3]
	fmt.Printf("HowSumMem(7, []int{2, 4}) = %v\n", HowSumMem(7, []int{2, 4}, make(map[int][]int))) //nil
	fmt.Printf("HowSumMem(8, []int{2, 3, 5}) = %v\n", HowSumMem(8, []int{2, 3, 5}, make(map[int][]int))) //[2, 2, 2, 2]
	fmt.Printf("HowSumMem(8, []int{3, 5, 2}) = %v\n", HowSumMem(8, []int{3, 5, 2}, make(map[int][]int))) //[2, 3, 3]
	fmt.Printf("HowSumMem(300, []int{7, 14}) = %v\n", HowSumMem(300, []int{7, 14}, make(map[int][]int))) //nil

}