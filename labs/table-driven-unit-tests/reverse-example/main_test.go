package reverseexample

import (
	"testing"
)

func TestReverseHello(t *testing.T) {
	input := "hello"
	expected := "olleh"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

func TestReverseEmpty(t *testing.T) {
	input := ""
	expected := ""
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

func TestReverseSingleChar(t *testing.T) {
	input := "a"
	expected := "a"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

func TestReverseWithPunctuation(t *testing.T) {
	input := "Go!"
	expected := "!oG"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

func TestReversePalindrome(t *testing.T) {
	input := "12321"
	expected := "12321"
	result := Reverse(input)
	if result != expected {
		t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
	}
}

// All of the above tests can be condensed into a single function using a table-driven approach. This makes the code cleaner and easier to maintain.
func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go!", "!oG"},
		{"12321", "12321"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
