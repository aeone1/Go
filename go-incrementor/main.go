package main

import "fmt"

type Incrementor struct {
	count int
}

func (self Incrementor) currentValue() int {
	return self.count
}
func (self *Incrementor) increment() {
	self.count++
}

func main() {
	counter := Incrementor{}
	counter.increment()
	counter.increment()

	fmt.Printf("current value %d\n", counter.currentValue())
}
