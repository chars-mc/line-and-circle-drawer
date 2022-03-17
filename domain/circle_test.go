package domain

import (
	"reflect"
	"testing"
)

func TestCircle_GetCoordinates(t *testing.T) {
	tests := []struct {
		name string
		c    *Circle
		want Coordinates
	}{
		// TODO: add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expectd: %v, got: %v", tt.want, got)
			}
		})
	}
}
