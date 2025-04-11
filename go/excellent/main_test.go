package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "偶数" {
		t.Errorf("expected: 偶数, actual: %s", result)
	}
}
