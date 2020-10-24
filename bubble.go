package sort

// BubbleSort regards the slice as two subsets: sorted(right) and unsorted(left).
// Iterates through the slice comparing adjacent values and swapping the values
// when n-1 > n. Values 'bubble' up to the sorted set.
//
// Complexity:
//     Average: O(n2)
//     Worst:   O(n2)
//     Best:    O(n)
//     Space:   O(1)
func BubbleSort(n []int) []int {
	for i := len(n) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if n[j-1] > n[j] {
				n[j], n[j-1] = n[j-1], n[j]
			}
		}
	}

	return n
}
