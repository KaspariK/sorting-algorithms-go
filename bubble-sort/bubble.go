package sorting_algorithms_go

// BubbleSort sorts the slice by iterating through the slice comparing adjacent
// values and swapping the values when n-1 > n.
//
// Complexity:
//     Average: O(n2)
//     Worst:   O(n2)
//     Best:    O(n)
//     Space:   O(1)
func BubbleSort(n []int) []int {
	// iterate over items in slice
		// for each item check item to the right, swap if left > right

	for i := len(n) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if n[j-1] > n[j] {
				n[j], n[j-1] = n[j-1], n[j]
			}
		}
	}

	return n
}