package main

import (
	"fmt"
	"testing"
)

// This is the classic, clean, with meaningful names for each scenario
func TestReverseNameEmpty(t *testing.T) {
	result, expected := ReverseName(""), ""
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestReverseNameOneWord(t *testing.T) {
	result, expected := ReverseName("Horace"), "Horace"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestReverseNameTwoWords(t *testing.T) {
	result, expected := ReverseName("Horace Walpole"), "Walpole Horace"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestReverseNameThreeWords(t *testing.T) {
	result, expected := ReverseName("Felipe Juan Froilan"), "Froilan Juan Felipe"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

// An alternative design pattern in Go is creating a struct table
// with the data for each scenario, and then iterate through it.
// Handle with care!

func TestReverseNameUsingStructTable(t *testing.T) {

	var testData = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Horace", "Horace"},
		{"Horace Walpole", "Walpole Horace"},
		{"Felipe Juan Froilan", "Froilan Juan Felipe"},
		{"John Ronald Reuel Tolkien", "Tolkien Reuel Ronald John"},
		{"Mambo N 5", "5 N Mambo"},
	}

	for _, testElement := range testData {
		testName := fmt.Sprintf("Test for %v", testElement.input)
		t.Run(testName, func(t *testing.T) {
			result := ReverseName(testElement.input)
			if result != testElement.expected {
				t.Errorf("Result is %v when %v is expected", result, testElement.expected)
			}
		})
	}
}
