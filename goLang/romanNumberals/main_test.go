package main

import (
	"testing"
)

func TestEight(t *testing.T) {
	result := intToRomanNumeral(8)
	if result != "VIII" {
		t.Errorf("Expect %v but got %v", "VIII", result)
	}
}

func TestNine(t *testing.T) {
	result := intToRomanNumeral(9)
	answer := "IX"
	if result != answer {
		t.Errorf("Expect %v but got %v", answer, result)
	}
}

func TestForty(t *testing.T) {
	result := intToRomanNumeral(40)
	if result != "XL" {
		t.Errorf("Expect %v but got %v", "XL", result)
	}
}
func TestYear(t *testing.T) {
	result := intToRomanNumeral(3749)
	answer := "MMMDCCXLIX"
	if result != answer {
		t.Errorf("Expect %v but got %v", answer, result)
	}
}
func Test1994(t *testing.T) {
	result := intToRomanNumeral(1994)
	answer := "MCMXCIV"
	if result != answer {
		t.Errorf("Expect %v but got %v", answer, result)
	}
}
