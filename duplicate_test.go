package main

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"empty", []int{}, false},
		{"single", []int{42}, false},
		{"no duplicates", []int{1, 2, 3, 4, 5}, false},
		{"duplicate", []int{1, 2, 2, 3}, true},
		{"duplicate at start", []int{7, 7, 1, 2}, true},
		{"all same", []int{5, 5, 5, 5}, true},
		{"negatives", []int{-1, 0, 1, -1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicate(tt.arr); got != tt.want {
				t.Errorf("HasDuplicate(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
