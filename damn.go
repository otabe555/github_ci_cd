package main

import (
	"fmt"
	"time"
)

func init() {
	time.Sleep(3 * time.Second)
}
func main() {
	fmt.Print("main starts")
	messages := make(chan string, 3)

	messages <- "one"
	messages <- "two"
	messages <- "three"

	go func(m *chan string) {
		fmt.Println("Entering the goroutine...")
		for {
			fmt.Println(<-*m)
		}
	}(&messages)
	time.Sleep(1 * time.Second)
}
