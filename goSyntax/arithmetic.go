package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("Apples:", 50-23)
	fmt.Println("Minutes in 2020 (secs):", 2020/60)
	// calculating the amount of time in a year
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	// calculate the duration between the start and end dates
	duration := end.Sub(start)
	// convert this to minutes
	minutes := int(duration.Minutes())
	// print it out
	fmt.Printf("Minutes in 2020 (year): %d\n", minutes)

	// convert an integer into a float and divide it by 13, then assign it an object
	a := float64(32452) / 13
	// perform math on this to get the integer to the highest next round number
	fmt.Printf("For 32,452 herring, you will need: %d tanks\n", int(math.Ceil(a)))
	// also perform modulo to get how many will be in the last tank
	fmt.Printf("There will be %d herring in the last tank", 32452%13)

	// calculate how long I've been on the makers course. Can use the same as the year function
	start_of_makers := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	makers_now := time.Now()

	time_at_makers := makers_now.Sub(start_of_makers)
	days := int(time_at_makers.Hours() / 24)
	fmt.Printf("\nTime at Makers as of now: %d", days)
}
