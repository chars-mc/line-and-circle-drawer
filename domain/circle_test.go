package domain

import (
	"reflect"
	"testing"
)

func TestCircle_getInitialCoordinates(t *testing.T) {
	tests := []struct {
		name string
		c    *Circle
		want Coordinates
	}{
		{
			name: "center (2, 3) and radio 10",
			c:    NewCircle(Coordinate{2, 3}, 10),
			want: Coordinates{
				Coordinate{0, 10},
				Coordinate{1, 10},
				Coordinate{2, 10},
				Coordinate{3, 10},
				Coordinate{4, 9},
				Coordinate{5, 9},
				Coordinate{6, 8},
				Coordinate{7, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getInitialCoordinates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expectd: %v, got: %v", tt.want, got)
			}
		})
	}
}
