package test

import (
	"fmt"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c = b
	t.Log(a, b, c)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	if s == "" {
		fmt.Println("*" + s)
	}
}
