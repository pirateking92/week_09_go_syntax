package main

import (
	"testing"
)

func TestReversesWordsCorrectly(t *testing.T) {
	result, expected := reverseString("david"), "divad"
	if result != expected {
		t.Errorf("Result is %v when %v is expected", result, expected)
	}
}

func TestMakeUppersLowersForPalindrowm(t *testing.T) {
	result, expected := palindromeCheck("Poop"), true
	if result != expected {
		t.Errorf("Result is %v", result)
	}
}

func TestIgnoreNonLetterChars(t *testing.T) {
	result, expected := palindromeCheck("!!!???KA££Y£££AK!@£$!£$%"), true
	if result != expected {
		t.Errorf("Non letter characters are not being ignored")
	}
}
