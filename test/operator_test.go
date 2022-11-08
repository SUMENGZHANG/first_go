package test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Execuable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	t.Log(a == b)

}

func TestConstant(t *testing.T) {
	a := 7 //0111

	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Execuable == Execuable)
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Execuable)
}
