package shixiseng

import "testing"

func TestNumberMapper(t *testing.T) {
	testData := []struct {
		data   string
		answer int
	}{
		{"e535", 0},
		{"f693", 1},
		{"e5e7", 2},
		{"f19e", 3},
		{"f31b", 4},
		{"ebae", 5},
		{"f307", 6},
		{"e689", 7},
		{"e221", 8},
		{"f661", 9},
	}
	for _, test := range testData {
		data := NumberMapper(test.data)
		if data != test.answer {
			t.Errorf("Value should be %d, but got %d", test.answer, data)
		}
	}
}
