package main

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}

		fmt.Print(" ")
	}
}
