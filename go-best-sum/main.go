package main
/*
Takes a target sum and an array of numbers as arguments

The func should return an array containing the shortest 
combination of numbers that add up to exactly the targetsum.

If there is any tie for the shortest combination, you may 
return any one of the shortest.
*/

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func BestSum(targetSum int, numbers []int) []int {
	/*
	N/B Taking the biggest number in numbers doesn't gaurantee
	fnding the shortest part.
	m = target sum
	n = len(numbers)

	time: O(n^m * m) //"* m" for array(slice) coppying.
	space: O(m^2)
	*/
	if targetSum == 0 {
		return []int{0}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int

	for _, num := range numbers {
		remainder := targetSum - num
		remainderCombination := BestSum(remainder, numbers)

		if remainderCombination != nil {
			combination := append(remainderCombination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
		
	}
	return shortestCombination
}

func BestSumMem(targetSum int, numbers []int, memo map[int][]int) []int {
	/*
	Memoized
	m = target sum
	n = len(numbers)

	time: O(n * m^2) //* m for array(slice) coppying.
	space: O(m^2) // extra m from memo
	*/
	
	// fmt.Printf("Memo: %v\n", memo)
	if val, ok := memo[targetSum]; ok {
		return val
	}

	if targetSum == 0 {
		return []int{0}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int
	// fmt.Printf("shortestCombination: %v\n", shortestCombination)

	for _, num := range numbers {
		remainder := targetSum - num
		remainderCombination := BestSumMem(remainder, numbers, memo)

		if remainderCombination != nil {
			combination := append(remainderCombination, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
		
	}
	memo[targetSum] = shortestCombination
	// fmt.Printf("shortestCombination: %v\n", shortestCombination)
	return memo[targetSum]
}

func main() {
	defer timeTrack(time.Now(), "BestSum")

	// fmt.Printf("BestSum(7, []int{5, 3, 4, 7}) = %v\n", BestSum(7, []int{5, 3, 4, 7})) //[7]
	// fmt.Printf("BestSum(8, []int{2, 3, 5}) = %v\n", BestSum(8, []int{2, 3, 5})) //[3, 5]
	// fmt.Printf("BestSum(8, []int{1, 4, 5}) = %v\n", BestSum(8, []int{1, 4, 5})) //[4, 4]
	// fmt.Printf("BestSum(100, []int{1, 2, 5, 25}) = %v\n", BestSum(100, []int{1, 2, 5, 25})) //[25, 25, 25, 25]

	// fmt.Printf("BestSumMem(7, []int{5, 3, 4, 7}) = %v\n", BestSumMem(7, []int{5, 3, 4, 7}, make(map[int][]int))) //[7]
	// fmt.Printf("BestSumMem(8, []int{2, 3, 5}) = %v\n", BestSumMem(8, []int{2, 3, 5}, make(map[int][]int))) //[3, 5]
	// fmt.Printf("BestSumMem(8, []int{1, 4, 5}) = %v\n", BestSumMem(8, []int{1, 4, 5}, make(map[int][]int))) //[4, 4]
	fmt.Printf("BestSumMem(100, []int{1, 2, 5, 25}) = %v\n", BestSumMem(100, []int{1, 2, 5, 25}, make(map[int][]int))) //[25, 25, 25, 25]

}
