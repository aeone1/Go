package main

// import "github.com/gofiber/fiber/v2"

// func main() {
//     app := fiber.New()

//     app.Get("/", func(c *fiber.Ctx) error {
//         return c.SendString("Hello, World 👋!")
//     })

//     app.Listen(":8000")
// }

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	//c := make(chan string, 2) //you could specify the size of a channel
														//This allows the channel to buffer until read.
	
	go count("sheep", c)

	// for {
	// 	msg, open := <- c // will recieve one message //"open" tells us if the channel is still open
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	
	//OR

	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
    for i := 1; i <= 5; i++ {
        c <- thing
        time.Sleep(time.Microsecond * 500)
    }
		close(c) // to avoid deadlock. When the reciever is waiting for nothing
}