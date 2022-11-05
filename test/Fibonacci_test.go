package test

import (
	"testing"
)

func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int =1
	var (
		a = 1
		b = 1
	)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tep := a
		a = b
		b = tep + a
	}
}
