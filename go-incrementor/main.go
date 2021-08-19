package main

import "fmt"

type Incrementor struct {
	count int
}

func (self Incrementor) getCurrentValue() int {
	return self.count
}
func (self *Incrementor) increment() {
	// Points to the "counter" variable
	// So it increments this particular variable(instance) sent.
	self.count++
}

func main() {
	counter := Incrementor{}
	counter.increment()
	counter.increment()

	fmt.Printf("current value %d\n", counter.getCurrentValue())
}
