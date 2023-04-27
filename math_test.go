package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(1, 1)

	if total != 2 {
		t.Errorf("\nInvalid sum.\nExpected: %d\nGot: %d", 2, total)
	}
}
