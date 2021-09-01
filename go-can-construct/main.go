package main

import (
	"fmt"
	"strings"
	"time"
)

/*
Check that a string of characters can be generated
from an array of strings.
Return True or False
*/

func main() {
	defer timeTrack(time.Now(), "CanConstruct")

	// fmt.Printf("CanConstruct(\"abcdef\", []string{\"ab\", \"abc\", \"cd\", \"def\", \"abcd\"}): %v\n", CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"})) // true
	// fmt.Printf("CanConstruct(\"skateboard\", []string{\"bo\", \"rd\", \"ate\", \"t\", \"ska\", \"sk\", \"boar\"}): %v\n", CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"})) // false
	// fmt.Printf("CanConstruct(\"enterapotentpot\", []string{\"a\", \"p\", \"ent\", \"enter\", \"ot\", \"o\", \"t\"}): %v\n", CanConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // true
	// fmt.Printf("CanConstruct(eeeeeee): %v\n", CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
	// 	"e",
	// 	"ee",
	// 	"eee",
	// 	"eeee",
	// 	"eeeee",
	// 	"eeeeee",})) // false

	fmt.Printf("CanConstructMemo(\"abcdef\", []string{\"ab\", \"abc\", \"cd\", \"def\", \"abcd\"}): %v\n", CanConstructMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}, make(map[string]bool))) // true
	fmt.Printf("CanConstructMemo(\"skateboard\", []string{\"bo\", \"rd\", \"ate\", \"t\", \"ska\", \"sk\", \"boar\"}): %v\n", CanConstructMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, make(map[string]bool))) // false
	fmt.Printf("CanConstructMemo(\"enterapotentpot\", []string{\"a\", \"p\", \"ent\", \"enter\", \"ot\", \"o\", \"t\"}): %v\n", CanConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, make(map[string]bool))) // true
	fmt.Printf("CanConstructMemo(eeeeeee): %v\n", CanConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{
		"e",
		"ee",
		"eee",
		"eeee",
		"eeeee",
		"eeeeee",}, make(map[string]bool))) // false

}

func CanConstruct(target string, wordBank []string) bool {
	/*
	m = len(target)
	n = len(wordBank)

	time complexity O(n^m * m) //"* m" creating suffix
	space '' O(m^2) // m is for suffix
	*/

	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			//fmt.Printf("suffix: %v\n", suffix)
			if CanConstruct(suffix, wordBank) == true {
				return true
			}
		}
	}

	return false
}

func CanConstructMemo(target string, wordBank []string, memo map[string]bool) bool {
	/*
	m = len(target)
	n = len(wordBank)

	time complexity O(n * m^2) //"* m" creating suffix, "* m" creating memo
	space '' O(m^2) // m is for suffix
	*/

	if val, ok := memo[target]; ok {
		return val
	}

	if target == "" {
		return true
	}

	for _, word := range wordBank {
		if strings.Index(target, word) == 0 {
			suffix := target[len(word):]
			//fmt.Printf("suffix: %v\n", suffix)
			if CanConstructMemo(suffix, wordBank, memo) == true {
				memo[target] = true
				return memo[target]
			}
		}
	}
	memo[target] = false
	return memo[target]
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
