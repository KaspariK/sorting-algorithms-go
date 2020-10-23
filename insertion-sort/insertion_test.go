package insertion

import (
	"reflect"
	"testing"
)

// TODO: Benchmark test
func TestInsertionSort(t *testing.T) {
	tests := map[string]struct{
		items []int
		want []int
	}{
		"Sorted slice": {items: []int{1,4,57,894,6549,6550,6550}, want: []int{1, 4, 23, 53, 98, 543}},
		"Unsorted slice": {items: []int{6550,1,4,894,6550,6549,57}, want: []int{1, 4, 23, 53, 98, 543}},
	}

	for name, tc := range tests {
		got := InsertionSort(tc.items)

		if reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
