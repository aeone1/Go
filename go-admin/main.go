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
	"sync"
	"time"
)

func main() {
 var wg sync.WaitGroup
 wg.Add(1)

 go func ()  {
    count("sheep")
    wg.Done() // reduces wg counter by 1
 }()

 wg.Wait() // will block till counter is at 0
}

func count(thing string) {
    for i := 1; i <= 5; i++ {
        fmt.Println(i, thing)
        time.Sleep(time.Microsecond * 500)
    }
}