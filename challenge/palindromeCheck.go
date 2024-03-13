package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromeCheck(word string) bool {
	fmt.Println("Enter 1 word for palicheck:")
	fmt.Scanln(&word)
	nonLetterRegex := regexp.MustCompile(`[^a-zA-Z]`)
	cleanedWord := nonLetterRegex.ReplaceAllString(word, "")
	lowerWord := strings.ToLower(cleanedWord)
	reversedWord := reverseString(lowerWord)

	return reversedWord == lowerWord
}

func main() {
	if palindromeCheck("") {
		fmt.Println("PALINDROOOOOME")
	} else {
		fmt.Println("NO PALINDRO")
	}
}
