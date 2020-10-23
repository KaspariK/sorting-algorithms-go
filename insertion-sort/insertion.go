package insertion

// InsertionSort sorts the slice by iterating through the slice comparing adjacent
// values and swapping the values when n-1 > n. The slice gets treated as two sub
// sets, the left slice is sorted and starts with length 1, as we iterate through
// the unsorted set we add items to the sorted set.
//
// Complexity:
//     Average: O(n2)
//     Worst:   O(n2)
//     Best:    O(n)
//     Space:   O(1)
func InsertionSort(n []int) []int {
	// start from 1 since 0 is already "sorted"
	for i := 1; i <= len(n)-1; i++ {
		for j:= i; j > 0; j-- {
			if n[j-1] > n[j] {
				n[j], n[j-1] = n[j-1], n[j]
			}
		}
	}

	return n
}
