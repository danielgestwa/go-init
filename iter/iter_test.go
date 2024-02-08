package main

import "testing"

func TestIteration(t *testing.T) {
	count := Loop(10)
	expected := 10

	if count != expected {
		t.Errorf("Expected: %d Got: %d", count, expected)
	}
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(5)
	}
}
