package main

import (
	"fmt"
	"time"
)


func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func ()  {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func ()  {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// for { //the func will wait for the slower(2 sec) func to complete even though the faster one is ready
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	//SELECT

	for { //outputs the first ready function
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}

}