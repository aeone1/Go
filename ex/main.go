package main

import "fmt"

type Set struct {
	items map[string]struct{}
}

func (s Set) Add(item string) {
	s.items[item] = struct{}{} // instead of bool
}

func (s Set) Remove(item string) {
	delete(s.items, item)
}

func main() {
	var s = Set{}
	s.items = map[string]struct{}{}
	s.Add("aaa")
	s.Add("bbb")
	s.Add("ccc")
	s.Add("aaa")
	s.Add("ddd")
	s.Add("bbb")
	s.Add("aaa")
	s.Add("eee")
	s.Remove("bbb")
	fmt.Println(s) // aaa ccc ddd eee
}
