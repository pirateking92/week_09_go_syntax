package main

import (
	"fmt"
	"strconv"
)

type drink struct {
	name   string
	volume float32
	price  int
}

// why would we do it this way:
func (myDrink drink) calculatePriceAtCheckout(quantity int) int {
	return myDrink.price * quantity
}

// as opposed to this way?:
func calculatePriceAtCheckout(myDrink drink, quantity int) int {
	return myDrink.price * quantity
}

type person struct {
	name     string
	age      int
	location string
}

func (onePerson person) introduce() string {
	return "\nI am " + onePerson.name + "! " + "I am " + strconv.Itoa(onePerson.age) + " years old and live in " + onePerson.location + "."
}

func (onePerson person) announce(otherPerson person) string {
	return "\nI, " + onePerson.name + ", would like to introduce " + otherPerson.name + " of " + otherPerson.location
}

func main() {
	coke := drink{name: "coke", volume: 0.33, price: 2}
	guinness := drink{name: "guinness", volume: 0.58, price: 4}
	juice := drink{name: "orange juice", volume: 0.33, price: 3}
	fmt.Printf("I bought two %v, and that cost me %d", juice.name, juice.calculatePriceAtCheckout(17)) // these change by adding the variable at the beginning rather than with the quantity
	fmt.Printf("\nI bought two %v, and that cost me %d", coke.name, coke.calculatePriceAtCheckout(2))
	fmt.Printf("\nI bought five %v, and that cost me %d", guinness.name, guinness.calculatePriceAtCheckout(5))

	people := []person{
		{name: "Matt Doyle", age: 31, location: "London"},
		{name: "Ted Theodore Logan", age: 40, location: "San Dimas"},
		{name: "Bill S. Preston Esquire", age: 40, location: "San Dimas"},
		{name: "Goku", age: 55, location: "Saiyan"},
		{name: "Monkey D. Luffy", age: 20, location: "Dawn Island"},
	}

	for i, p := range people {
		fmt.Print(p.introduce())
		if i < len(people)-1 {
			fmt.Print(p.announce(people[i+1]))
		}
	}
}
