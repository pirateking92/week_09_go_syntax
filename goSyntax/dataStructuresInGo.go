package main

import (
	"fmt"
	"strings"
)

func grumpyArray() {
	var grumpyArray [5]int
	grumpyArray[2] = 45
	fmt.Println("I hope I access a correct index to show this number:", grumpyArray[2])
	fmt.Println("Other elements are initialized with default value:", grumpyArray[3])

	// Arrays can be initialized in one go, like other variables in Go
	initializedArray := [5]string{"Guayou", "Valteri", "Alex", "Daniel", "Max"}
	for position, name := range initializedArray {
		fmt.Printf("\n%v will not do big one liners until they have mastered lesson %d", name, position)
	}
}

func character() {
	archerCharMap := make(map[string]int) // String is the key, int the value in this example
	archerCharMap["speed"] = 8
	archerCharMap["attack"] = 9
	archerCharMap["defence"] = 5
	archerCharMap["range"] = 15
	fmt.Println("My character stats:", archerCharMap)

	awesomeAttack := archerCharMap["attack"]

	// deleting one element
	delete(archerCharMap, "defence")
	fmt.Printf("Who needs defence with this attack %d! New stats: %v", awesomeAttack, archerCharMap)
}

func changeASlice() {
	myList := []string{"Cat", "Mouse", "Frog"}

	myList = append(myList, "Lynx")
	// Write some code to amend the list here.

	fmt.Println(myList)

	// Should print:
	// [Mouse Lynx Cat Frog]
}

func petList() {
	petList := []string{"Cat", "cat", "frog", "cat", "dog", "Dog"}
	counters := make(map[string]int)

	for i, item := range petList {
		petList[i] = strings.ToLower(item)
	}

	for _, item := range petList {
		counters[item]++
	}
	fmt.Println(counters)
	// Write some code to count the items in the list here.

	// Should print (in any order)
	// [cat:3 dog:2 frog: 1]
}

func slice() {
	shySlice := make([]string, 3)
	fmt.Println("So shy it is empty:", shySlice)
	shySlice[0] = "Hello,"
	shySlice[1] = "I am"
	shySlice[2] = "Shy"
	fmt.Println("Not so shy anymore:", shySlice)

	// Add an element to the slice
	shySlice = append(shySlice, "until you knew me better")
	// Update an element
	shySlice[1] = "I was"
	fmt.Println("I am the heart of the party now:", shySlice)
}

func main() {
	shySlice := []string{"Matt", "Katrina", "Dom", "Jason", "Aakash", "Sam", "Maria", "Majo", "Bobby", "Ben", "Bogdan", "Hayri"}

	cohoList := make(map[string]int)

	// making all the names lowercase
	for _, item := range shySlice {
		lowercaseName := strings.ToLower(item)
		// extracting the first letter of each name
		firstLetter := lowercaseName[:1]
		// go through these and group them together
		cohoList[firstLetter]++
	}
	// we iterate through the shySlice and for each name,
	// we convert it to lowercase using strings.ToLower.
	// Then, we extract the first letter of the lowercase
	// name using string slicing (lowercaseName[:1]).
	// Finally, we increment the count for that first letter in the cohoList map.
	// When you run this code, the output will be a map where the keys are the
	// first letters of the names, and the values are the counts of names starting with that letter.

	fmt.Println("My cohort yo!:", shySlice)
	fmt.Println(cohoList)
}
