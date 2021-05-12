package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 2)

	if result != 4 {
		t.Errorf("Result of sum should be 4, but it was %d", result)
	}
}
