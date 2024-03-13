package main

import (
	"fmt"
)

func main() {
	for loopPosition := 0; loopPosition <= 7; loopPosition++ {
		fmt.Println("I am looping!!:", loopPosition)

	}
	rangerman()

}

func ranger() {
	nums := []int{1, 3, 5, 7, 8, 9, 11, 13, 14, 15, 16, 17} // This is a slice (dynamic array)
	for loopPosition, num := range nums {
		if num%2 == 0 {
			fmt.Println("I found the evil even number at position:", loopPosition, "\nand this is the number!:", num)
		} else {
			fmt.Println("I am looping happily with my oddy:", num)
		}
	}
}

func rangerman() {
	start := 2321
	end := 34325
	var my_list []int // defining an empty list so we can add stuff to it (i). declaring variable at the beginning
	for i := start; i <= end; i++ {
		my_list = append(my_list, i) // appending the empty list with i
	}
	sum := 0
	for num := range my_list {
		if num%2 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}

func fizzbuzz() {
	var nums []int
	for i := 0; i <= 100; i++ {
		nums = append(nums, i)
	}
	for num := range nums {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(num)
		}
	}
}

// some research has said that it will either
// be done recursively, or iteratively/
// Try iteratively first, as this is what we're
// learning here first. You've opened a webpage
// on it already.
func factorial() {

}
