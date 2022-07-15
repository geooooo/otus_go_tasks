package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	var (
		packedString string
		expected     string
	)

	packedString = "a1"
	expected = "a"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a12"
	expected = "aaaaaaaaaaaa"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a4bc2d5e"
	expected = "aaaabccddddde"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "abcd"
	expected = "abcd"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a2bcd"
	expected = "aabcd"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = ""
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a0"
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a1b0c2"
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "45"
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a-1"
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "a1b-1"
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = "п2г3"
	expected = "ппггг"
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `qwe\4\5`
	expected = `qwe45`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `qwe\45`
	expected = `qwe44444`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `qwe\\5`
	expected = `qwe\\\\\`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\0`
	expected = `0`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\1\02`
	expected = `100`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\1\`
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\e`
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `ab2\e`
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `ab2\ec`
	expected = ""
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\\\\`
	expected = `\\`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\\2\02a2`
	expected = `\\00aa`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `abc3\\1\\2`
	expected = `abccc\\\`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\0\0a\03`
	expected = `00a000`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\02\\2\02`
	expected = `00\\00`
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}

	packedString = `\0-1`
	expected = ``
	if result := UnpackString(packedString); result != expected {
		t.Fatalf("UnpackString(%s)=%s should be %s\n", packedString, result, expected)
	}
}
