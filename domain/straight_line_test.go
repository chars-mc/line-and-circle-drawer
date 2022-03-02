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
		// TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GetCoordinates(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
