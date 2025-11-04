package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0

	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return original == reversed
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(1214))
}
