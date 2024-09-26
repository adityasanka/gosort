package gosort

import "golang.org/x/exp/constraints"

// BubbleSort performs the bubble sort algorithm on a copy of the given slice.
// It takes a slice of any ordered type T and returns a new sorted slice.
//
// The function does not modify the original slice. Instead, it creates and returns
// a new sorted slice, preserving the original input.
//
// Time complexity: O(n^2), where n is the number of elements in the slice.
// Space complexity: O(n), as it creates a copy of the input slice.
//
// Parameters:
//   - list: A slice of type T, where T satisfies the constraints.Ordered interface
//
// Returns:
//   - A new sorted slice of type []T
//
// Example:
//
//	numbers := []int{64, 34, 25, 12, 22, 11, 90}
//	sortedNumbers := BubbleSort(numbers)
//	fmt.Println(sortedNumbers) // Output: [11 12 22 25 34 64 90]
//	fmt.Println(numbers)       // Output: [64 34 25 12 22 11 90] (original slice unchanged)
//
//	words := []string{"banana", "apple", "cherry", "date"}
//	sortedWords := BubbleSort(words)
//	fmt.Println(sortedWords)   // Output: [apple banana cherry date]
//	fmt.Println(words)         // Output: [banana apple cherry date] (original slice unchanged)
func BubbleSort[T constraints.Ordered](list []T) []T {
	// Create a copy of the input slice
	sortedList := make([]T, len(list))
	copy(sortedList, list)

	n := len(sortedList)

	for k := 0; k < n-1; k++ {
		for i := 0; i < n-k-1; i++ {
			if sortedList[i] > sortedList[i+1] {
				sortedList[i], sortedList[i+1] = sortedList[i+1], sortedList[i]
			}
		}
	}

	return sortedList
}
