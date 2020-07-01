package main

import "testing"

func TestAdd1(t *testing.T) {
	result := add1(3)
	if result != 2 {
		t.Errorf("Wrong sum, got result: %d, but instead want: %d", result, 2)
	}
}
