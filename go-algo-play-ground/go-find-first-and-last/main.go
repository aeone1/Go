package main

import (
	"fmt"
	"time"
)

/*
Given an array of integers "nums" sorted in ascending order,
find the starting and ending position of a given "target" value.

Your algorithm's runtime complexity must be in order of O(log n).

If the target is not found in the array, return [-1, -1].
*/

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func main(){
	defer timeTrack(time.Now(), "SearchRange")

	fmt.Printf("SearchRange([]int{5,7,7,8,8,9}, 8): %v\n", SearchRange([]int{5,7,7,8,8,9}, 8)) //[3,4]
	fmt.Printf("SearchRange([]int{5,7,7,8,8,9}, 6): %v\n", SearchRange([]int{5,7,7,8,8,9}, 6)) //[-1,-1]

}

func SearchRange(nums []int, target int ) []int {
	result := make([]int, 2)
	result[0] = FindStartingIndex(nums, target)
	result[1] = FindEndingIndex(nums, target)
	return result
}

func FindStartingIndex(nums []int, target int ) int {
	index := -1
	start := 0
	end := len(nums)-1

	for start <= end {
		midpoint := start + (end-start) / 2 // avoids overflow(?)
		if nums[midpoint] >= target {
			end = midpoint-1 
		} else {
			start = midpoint+1
		}
		if nums[midpoint] == target {
			index = midpoint
		}
	}

	return index
}

func FindEndingIndex(nums []int, target int ) int {
	index := -1
	start := 0
	end := len(nums)-1

	for start <= end {
		midpoint := start + (end-start)/2
		if nums[midpoint] <= target {
			start = midpoint+1
		} else {
			end = midpoint-1
		}
		if nums[midpoint] == target {
			index = midpoint
		}
	}

	return index
}
