package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(1, 1)
	expected := 2

	if total != expected {
		t.Errorf("\nInvalid sum.\nExpected: %d\nGot: %d", expected, total)
	}
}

func TestSub(t *testing.T) {
	result := Sub(1, 1)
	expected := 0

	if result != expected {
		t.Errorf("\nInvalid subtraction.\nExpected: %d\nGot: %d", expected, result)
	}
}
