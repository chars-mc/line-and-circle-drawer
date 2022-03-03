package domain

import (
	"reflect"
	"testing"
)

func TestStraightLine_GetCoordinates(t *testing.T) {
	tests := []struct {
		name     string
		l        *StraightLine
		expected Coordinates
	}{
		{
			name: "A(2, 2) and B(7, 8) - slope less than 1",
			l:    NewStraightLine(Coordinate{2, 2}, Coordinate{7, 8}),
			expected: Coordinates{
				{2, 2}, {3, 3}, {4, 4}, {5, 6}, {6, 7},
			},
		},
		{
			name: "A(2, 2) and B(8, 2) - slope 0 (0 deg)",
			l:    NewStraightLine(Coordinate{2, 2}, Coordinate{8, 2}),
			expected: Coordinates{
				{2, 2}, {3, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2},
			},
		},
		{
			name: "A(2, 2) and B(9, 9) - slope 1 (45 deg)",
			l:    NewStraightLine(Coordinate{2, 2}, Coordinate{9, 9}),
			expected: Coordinates{
				{2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8},
			},
		},
		// TODO: add undefined slope case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GetCoordinates(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
