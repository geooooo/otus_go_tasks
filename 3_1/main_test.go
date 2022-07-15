package main

import (
	"testing"
)

func TestItoa(t *testing.T) {
	var (
		number   int
		expected string
	)

	number = 0
	expected = "0"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = 1
	expected = "1"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = 12
	expected = "12"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = 123
	expected = "123"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = -1
	expected = "-1"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = -12
	expected = "-12"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}

	number = -123
	expected = "-123"
	if result := Itoa(number); result != expected {
		t.Fatalf("Itoa(%d) should be \"%s\"\n", number, expected)
	}
}
