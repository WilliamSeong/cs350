package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	println("Hello World!")
	fmt.Println("Hello World!, but using the imported fmt package hmm")
	fmt.Println("The time right now is ", time.Now())
	fmt.Println("This is an exported name", math.Pi)
	fmt.Println("By the way, this print is a Goroutine\n")

	fmt.Println("You can assign variable like 'var x int = 0' or you can do 'x := int 0'\n")

	y := 10
	fmt.Println(y, "\n")

	fmt.Println("Here's a loop")
	var ctr int = 0
	for ctr < 5 {
		fmt.Println(ctr)
		ctr++
	}

	fmt.Println("Here's another way to write a loop\n")
	for ctr := 0; ctr < 10; ctr++ {
		fmt.Println(ctr)
	}

	cities := []string{"Seoul", "New York", "Boston"}
	for index, value := range cities {
		fmt.Println(index, value)
	}

	for index, value := range cities {
		fmt.Println(index, value)
	}

	var salary int = 100
	fmt.Println(salary)

	salary1 := 100
	fmt.Println(salary1)

	if salary < 50 {
		fmt.Println("You are underpaid!")
	} else {
		fmt.Println("You are sufficiently paid!")
	}

	if x := 9; x < 10 {
		fmt.Println("foo")
	} else {
		fmt.Println("bar")
	}

	rating := 2
	switch rating {
	case 4:
		fmt.Println("Excellent")
	case 3:
		fmt.Println("Good")
	case 2:
		fmt.Println("Consistent")
	case 1:
		fmt.Println("Improved")
	default:
		fmt.Println("N/A")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	case t.Hour() < 22:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good night.")
	}

}
