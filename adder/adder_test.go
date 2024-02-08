package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	assertNum(t, sum, expected)
}

func TestAddToInt(t *testing.T) {
	num := Number{value: 3}
	num.AddTo(4)

	sum := num.value
	expected := 7

	assertNum(t, sum, expected)
}

func assertNum(t *testing.T, sum, expected int) {
	if sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}
