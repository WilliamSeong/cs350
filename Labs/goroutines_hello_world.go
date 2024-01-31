package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world", time.Now())
}

func goodbye() {
	fmt.Println("Goodbye world", time.Now())
}

func main() {
	fmt.Println(time.Now())
	go hello()
	go goodbye()
	time.Sleep(1 * time.Second)
	// now we exit
}

// call main
// spin up hello in a goroutine
// spin up goodbye in a goroutine
// (sleep if not commented out)
// main exits
// if you didn't sleep then the program exits before the goroutines can run (most of the time)
