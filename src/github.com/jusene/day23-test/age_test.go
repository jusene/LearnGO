package main

import (
	"fmt"
	"testing"
)

func TestAge(t *testing.T) {
	var (
		input    = -100
		expected = 0
	)

	actual := Age(input)
	if actual != expected {
		t.Errorf("Age(%d) = %d 预期 %d", input, actual, expected)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("HELLO")
	}
}
