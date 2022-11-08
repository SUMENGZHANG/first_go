package test

import "testing"

func TestLoop(t *testing.T) {
	n := 0
	for n < 5 {
		n++
		t.Log(n)
	}
}
