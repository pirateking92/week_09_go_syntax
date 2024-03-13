package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// result, counter := 3.0, 7.0

	// if math.Sqrt(counter) == result {
	// 	fmt.Println("Counter is Nine!")
	// } else if result := 49.0; math.Pow(counter, 3.0) == result {
	// 	fmt.Println("Here, we're changing the value of result to 49, then getting the power of counter (7) to 2,\nwhich is 7*2, then checking that its equal to result, which it is, which gives us this message!")
	// } else {
	// 	fmt.Println("if the if and else if arent reached, then we should be here")
	// }
	var number int
	fmt.Print("Enter a number to factoralise!:")
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal("Invalid input. Please enter a valid non-negative integer")
	}
	if number < 0 {
		log.Fatal("Invalid input. Please enter a non-negative integer")
	}

	// calculate factorial using iterative approach
	factorial_iter := iterative_factorial(number)
	fmt.Println("Factorial of", number, "(Iterative):", factorial_iter)
}

func next() {
	const greaterThan = 2 > 1
	const lessThan = 2 < 1
	const biggerEqual = 2 >= 2
	const lessEqual = 2 <= 2
	const equal = 2 == 2
	const notEqual = 2 != 2
	const stringEqual = "Hello" == "Hello"
	const stringNotEqual = "Hello" != "Hello"
	const equalSimilarType = 2 == 2.0
	// const equalDifferentType = 2 == "2"
	// This is not allowed by the compiler
	const opposite = !true
	fmt.Println("The results:", greaterThan, lessThan, biggerEqual, lessEqual, equal, notEqual, stringEqual, stringNotEqual, equalSimilarType, opposite)
}

func dr() {
	person := "Doctor"
	switch person {
	case "Doctor":
		fmt.Println("Doctor is in consult room")
	case "Nurse":
		fmt.Println("Nurse is preparing the injection")
	case "Patient":
		fmt.Println("Patient is in the waiting room")
	default:
		fmt.Println("This is a Hospital")
	}
}

func random() {
	randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	iAmNumber := randomSeed.Intn(10)

	if iAmNumber == 1 {
		fmt.Println("One!")
	} else if iAmNumber == 2 || iAmNumber == 3 {
		fmt.Println("Could be 2 or 3")
	} else if iAmNumber == 4 {
		fmt.Println("four")
	} else if iAmNumber > 5 {
		fmt.Println("Too much randomness")
	}

	fmt.Println(iAmNumber)
}

func fizzbuzz() {
	randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	seconds := randomSeed.Intn(60)

	if seconds%3 == 0 && seconds%5 == 0 {
		fmt.Println("Fizzbuzz")
	} else if seconds%5 == 0 {
		fmt.Println("Buzz")
	} else if seconds%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(seconds)
	}
	fmt.Println(seconds)
}

// challenge

// Write a program to find the factorial of 17.

// The factorial of a number is that number multiplied by every number below it until 1.
// So, 10 factorial is 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1, evaluating to 3628800.

// When correct, your program should produce the number 355687428096000.

func iterative_factorial(n int) int {
	factorial := 1
	// multiply numbers from 2 to n to calculate factorial
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}
