package main

import "testing"

func Test_addTwoNumbers(t *testing.T) {
	answer := addTwoNumbers(2, 2)
	if answer != 4 {
		t.Errorf("addTwoNumbers(2,2) = %d; want 4", answer)
	}
}
