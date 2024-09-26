// Package gosort provides implementations of various sorting algorithms.
//
// This package offers generic sorting functions that work with any type
// that satisfies the constraints.Ordered interface from the golang.org/x/exp/constraints
// package. This allows for sorting of slices containing basic types like int,
// float64, and string, as well as custom types that implement the necessary
// comparison methods.
//
// Currently implemented algorithms:
//   - Bubble Sort
//
// Usage:
//
//	import "github.com/yourusername/gosort"
//
//	numbers := []int{64, 34, 25, 12, 22, 11, 90}
//	sortedNumbers := gosort.BubbleSort(numbers)
//
// Note: The sorting functions in this package modify the original slice in-place
// and also return the sorted slice for convenience.
package gosort
