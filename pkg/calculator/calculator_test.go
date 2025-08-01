package calculator

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		name      string
		order     int
		packSizes []int
		expected  map[int]int
	}{
		{
			name:      "Order for 1 item",
			order:     1,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{250: 1},
		},
		{
			name:      "Order for 250 items",
			order:     250,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{250: 1},
		},
		{
			name:      "Order for 251 items",
			order:     251,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{500: 1},
		},
		{
			name:      "Order for 501 items",
			order:     501,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{500: 1, 250: 1},
		},
		{
			name:      "Order for 12001 items",
			order:     12001,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			name:      "Order for 0 items",
			order:     0,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{},
		},
		{
			name:      "Empty pack sizes",
			order:     100,
			packSizes: []int{},
			expected:  map[int]int{},
		},
		{
			name:      "Order for edge case",
			order:     500000,
			packSizes: []int{23, 31, 53},
			expected:  map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePacks(tt.order, tt.packSizes)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CalculatePacks() = %v, want %v", got, tt.expected)
			}
		})
	}
}
