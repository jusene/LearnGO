package main

import "testing"

func TestPrime(t *testing.T) {
	var primeTests = []struct {
		input    int
		expected bool
	}{
		{1, false},
		{2, true},
	}

	for _, tt := range primeTests {
		actual := isPrime(tt.input)
		if actual != tt.expected {
			t.Errorf("%d %v %v", tt.input, actual, tt.expected)
		}
	}
}
