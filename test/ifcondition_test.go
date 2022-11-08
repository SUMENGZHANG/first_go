package test

import "testing"

func TestIf(t *testing.T) {
	if a := 1; a < 10 {
		a++
		t.Log(a)
	}
}

func TestSwitchCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {

		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is beyond 3")
		}
	}
}

func TestSwitchAsIfElse(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("default")
		}
	}
}
