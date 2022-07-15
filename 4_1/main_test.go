package main

import (
	"testing"
)

func Test(t *testing.T) {
	var (
		slice       [][]int
		expected    []int
		result      []int
		lenResult   int
		lenExpected int
	)

	// Case 1
	slice = [][]int{{1, 2, 3}, {4, 5}, {6, 7}}
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	result = MakeFlatSlice(slice)

	lenResult = len(result)
	lenExpected = len(expected)
	if lenResult != lenExpected {
		t.Fatalf("len(result)=%d != len(expected)=%d\n", lenResult, lenExpected)
	}

	for i := 0; i < lenResult; i++ {
		if result[i] != expected[i] {
			t.Fatalf("result[%d]=%d != expected[%d]=%d\n", i, result[i], i, expected[i])
		}
	}

	// Case 2
	slice = [][]int{}
	expected = []int{}
	result = MakeFlatSlice(slice)

	lenResult = len(result)
	lenExpected = len(expected)
	if lenResult != lenExpected {
		t.Fatalf("len(result)=%d != len(expected)=%d\n", lenResult, lenExpected)
	}

	for i := 0; i < lenResult; i++ {
		if result[i] != expected[i] {
			t.Fatalf("result[%d]=%d != expected[%d]=%d\n", i, result[i], i, expected[i])
		}
	}

	// Case 3
	slice = [][]int{{1}, {2, 3}}
	expected = []int{1, 2, 3}
	result = MakeFlatSlice(slice)

	lenResult = len(result)
	lenExpected = len(expected)
	if lenResult != lenExpected {
		t.Fatalf("len(result)=%d != len(expected)=%d\n", lenResult, lenExpected)
	}

	for i := 0; i < lenResult; i++ {
		if result[i] != expected[i] {
			t.Fatalf("result[%d]=%d != expected[%d]=%d\n", i, result[i], i, expected[i])
		}
	}

	// Case 4
	slice = [][]int{{1, 2}}
	expected = []int{1, 2}
	result = MakeFlatSlice(slice)

	lenResult = len(result)
	lenExpected = len(expected)
	if lenResult != lenExpected {
		t.Fatalf("len(result)=%d != len(expected)=%d\n", lenResult, lenExpected)
	}

	for i := 0; i < lenResult; i++ {
		if result[i] != expected[i] {
			t.Fatalf("result[%d]=%d != expected[%d]=%d\n", i, result[i], i, expected[i])
		}
	}
}
