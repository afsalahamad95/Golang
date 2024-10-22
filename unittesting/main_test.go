package main

import "testing"

func divide(a, b int) int {
	return a / b
}
func TestDivide(t *testing.T) {
	expected := 5
	got := divide(10, 2)
	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	} else {
		t.Log("test passed")
	}
}
