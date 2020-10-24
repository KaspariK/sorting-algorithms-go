package sort

import (
"reflect"
"testing"
)

// Randomly generated sets of 10 integers between 1 and 10,000
//
// 4231, 4958, 5894, 4048, 1328, 7331, 7091, 704, 6775, 8477
// 589, 1432, 1078, 5331, 1830, 786, 9439, 427, 468, 8176
// 2918, 3848, 3682, 6249, 8004, 1807, 2683, 8590, 791, 8382
// 4509, 4635, 9482, 3837, 9901, 8833, 8763, 5203, 1152, 5580

// TODO: Benchmark test
func TestBubbleSort(t *testing.T) {
	tests := map[string]struct{
		items []int
		want []int
	}{
		"Sorted slice": {items: []int{1562, 2244, 2410, 3157, 4356, 5802, 6688, 6926, 9187, 9855}, want: []int{1562, 2244, 2410, 3157, 4356, 5802, 6688, 6926, 9187, 9855}},
		"Unsorted slice": {items: []int{7628, 4451, 6040, 7123, 1483, 2340, 7931, 4080, 5482, 9492}, want: []int{1483, 2340, 4080, 4451, 5482, 6040, 7123, 7628, 7931, 9492}},
	}

	for name, tc := range tests {
		got := BubbleSort(tc.items)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

// TODO: Benchmark test
func TestInsertionSort(t *testing.T) {
	tests := map[string]struct{
		items []int
		want []int
	}{
		"Sorted slice": {items: []int{629, 7237, 1927, 9568, 3210, 6187, 74, 8393, 2510, 4347}, want: []int{74, 629, 1927, 2510, 3210, 4347, 6187, 7237, 8393, 9568}},
		"Unsorted slice": {items: []int{7003, 9058, 876, 9670, 3876, 7756, 4222, 8220, 9147, 7895}, want: []int{876, 3876, 4222, 7003, 7756, 7895, 8220, 9058, 9147, 9670}},
	}

	for name, tc := range tests {
		got := InsertionSort(tc.items)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

// TODO: Benchmark test
func TestSelectionSort(t *testing.T) {
	tests := map[string]struct{
		items []int
		want []int
	}{
		"Sorted slice": {items: []int{110, 756, 1071, 1894, 4019, 5613, 5739, 7812, 9590, 9504}, want: []int{110, 756, 1071, 1894, 4019, 5613, 5739, 7812, 9504, 9590}},
		"Unsorted slice": {items: []int{7812, 5582, 3142, 7010, 2894, 6752, 9468, 5795, 9553, 6040}, want: []int{2894, 3142, 5582, 5795, 6040, 6752, 7010, 7812, 9468, 9553}},
	}

	for name, tc := range tests {
		got := SelectionSort(tc.items)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
