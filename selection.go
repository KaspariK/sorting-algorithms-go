package sort

// SelectionSort regards the slice as two subsets: sorted(left) and
// unsorted(right). Iterate through the unsorted set to find the minimum value.
// Each time the minimum is determined, add it to the sorted set.
//
// Complexity:
//     Average: O(n2)
//     Worst:   O(n2)
//     Best:    O(n)
//     Space:   O(1)
func SelectionSort(n []int) []int {
	for i := 0; i < len(n) - 1; i++ {
		minIndex := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[minIndex] {
				minIndex = j
			}
		}
		n[i], n[minIndex] = n[minIndex], n[i]
	}

	return n
}