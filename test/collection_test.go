package test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int

	arr1 := [...]int{1, 2, 3, 4}
	t.Log(arr[2], arr1[1])

	for _, e := range arr1 {
		t.Log(e)
	}

	t.Log(arr1[2:])
}
