package main

import "fmt"

func Palindrome(str string) {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			fmt.Printf("%s - not palindrome\n", str)
			return
		}
	}
	fmt.Printf("%s - palindrome\n", str)
}
