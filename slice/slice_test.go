package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("4 items", func(t *testing.T) {
		nums := []int{3, 4, 5, 6}

		got := Sum(nums)
		expected := 18

		if got != expected {
			t.Errorf("Expected: %d, got %d", expected, got)
		}
	})

	t.Run("5 items", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := Sum(nums)
		expected := 15

		if got != expected {
			t.Errorf("Expected: %d, got %d", expected, got)
		}
	})
}

func TestSumElements(t *testing.T) {
	t.Run("One element array", func(t *testing.T) {
		num_one := []int{3}
		num_two := []int{4}
		got := SumElements(num_one, num_two)
		expected := []int{7}

		checkArrError(t, got, expected)
	})

	t.Run("Multiple element array same size", func(t *testing.T) {
		num_one := []int{3, 4}
		num_two := []int{4, 5}
		got := SumElements(num_one, num_two)
		expected := []int{7, 9}

		checkArrError(t, got, expected)
	})

	t.Run("Multiple element array first bigger", func(t *testing.T) {
		num_one := []int{2, 10, 9}
		num_two := []int{4, 5}
		got := SumElements(num_one, num_two)
		expected := []int{6, 15, 9}

		checkArrError(t, got, expected)
	})

	t.Run("Multiple element array second bigger", func(t *testing.T) {
		num_one := []int{2, 1}
		num_two := []int{3, 9, 8}
		got := SumElements(num_one, num_two)
		expected := []int{5, 10, 8}

		checkArrError(t, got, expected)
	})

	t.Run("Multiple element with negatives", func(t *testing.T) {
		num_one := []int{-2, 1, -2}
		num_two := []int{3, -9, -3}
		got := SumElements(num_one, num_two)
		expected := []int{1, -8, -5}

		checkArrError(t, got, expected)
	})

	t.Run("Empty", func(t *testing.T) {
		num_one := []int{}
		num_two := []int{}
		got := SumElements(num_one, num_two)
		expected := []int{}

		checkArrError(t, got, expected)
	})
}

func checkArrError(t *testing.T, got []int, expected []int) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %d, got %d", expected, got)
	}
}
