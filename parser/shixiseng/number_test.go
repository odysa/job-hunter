package shixiseng

import "testing"

func TestNumberMapper(t *testing.T) {
	testData := []struct {
		data   string
		answer int
	}{
		{"f0e5", 0},
		{"e0d8", 1},
		{"e5d1", 2},
		{"f313", 3},
		{"e609", 4},
		{"eb5b", 5},
		{"f2e1", 6},
		{"eb83", 7},
		{"f4c4", 8},
		{"eb46", 9},
	}
	for _, test := range testData {
		data := NumberMapper(test.data)
		if data != test.answer {
			t.Errorf("Value should be %d, but got %d", test.answer, data)
		}
	}
}
