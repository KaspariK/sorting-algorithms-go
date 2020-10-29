package sort

// QuickSort is a divide and conquer algorithm that selects a pivot and
// partitions the slice into sub sets (> pivot or < pivot). Slice is sorted
// recursively. Starting with the lomuto partition scheme but will iterate until
// I feel like it's good enough
//
// Complexity:
//     Average: O(n log n)
//     Worst:   O(n2)
//     Best:    O(n)
//     Space:   O(n) or maybe O(log n)?
func QuickSort(a []int, lo int, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		QuickSort(a, lo, p - 1)
		QuickSort(a, p + 1, hi)
	}
}

func partition(a []int, lo int, hi int) int {
	pivot := a[hi]
	i := lo
	// j should be index of lo?
	for j := 0; j < hi; j++ {
		if a[j] < pivot {
			// swap a[i] with a[j]
			i += 1
		}
		// swap a[i] with a[hi]
	}
	return i
}
