package divideexample

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
		wantErr  bool
	}{
		{10, 2, 5, false},
		{5, 0, 0, true},
		{9, 3, 3, false},
	}

	for _, test := range tests {
		result, err := Divide(test.a, test.b)
		if (err != nil) != test.wantErr {
			t.Errorf("Divide(%v, %v) error = %v, wantErr %v", test.a, test.b, err, test.wantErr)
		}
		if !test.wantErr && result != test.expected {
			t.Errorf("Divide(%v, %v) = %v; want %v", test.a, test.b, result, test.expected)
		}
	}
}
