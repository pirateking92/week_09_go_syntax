package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("Cannot divide by zero duh...")
	}

	return float64(a) / float64(b), nil
}

func addAndMultiplyIntegers(a int, b int) (int, int) {
	return a + b, a * b
}

func callingAddAndMultiply() {
	var a, b int
	fmt.Println("Enter 2 integers separated by a space ^^ :")
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println("I said you need to enter 2 integers separated by a space...")
		return
	}
	var sum, multiplication int
	sum, multiplication = addAndMultiplyIntegers(a, b)
	fmt.Printf("The result of adding %d and %d is: %d", a, b, sum)
	fmt.Printf("\nThe result of multipliying %d and %d is: %d", a, b, multiplication)
}

func greetNumber(number int) string {
	return "Nice to meet you " + strconv.Itoa(number)

}

func subtractFromHigh() int {
	var a, b int
	fmt.Println("\nEnter 2 numbers separated by a space:")
	fmt.Scanln(&a, &b)
	if a > b {
		a -= b
		return a
	} else if b > a {
		b -= a
		return b
	}
	return 0
}

func addFive(num int) int {
	return num + 5
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func superify(word string) (string, string) {

	fmt.Println("Enter 1 word you want to superify/batify:")
	fmt.Scanln(&word)

	reversedWord := reverseString(word)

	return "Super" + word, "Bat" + reversedWord
}

func main() {
	// var x, y int
	// fmt.Println("Enter 2 integers separated by a space!!!")
	// _, err := fmt.Scanln(&x, &y)
	// result, err := divide(x, y)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("Dividing %d by %d equals %f", x, y, result)
	fmt.Println(greetNumber(addFive(subtractFromHigh())))
	fmt.Println(superify(""))
}
