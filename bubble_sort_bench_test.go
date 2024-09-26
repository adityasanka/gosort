package gosort

import (
	"math/rand"
	"sort"
	"testing"
)

// generateSlice creates a slice of integers with the given size
func generateSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(1000) // Random integers between 0 and 999
	}
	return slice
}

// BenchmarkBubbleSortSmall benchmarks BubbleSort with a small input size
func BenchmarkBubbleSortSmall(b *testing.B) {
	input := generateSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

// BenchmarkBubbleSortMedium benchmarks BubbleSort with a medium input size
func BenchmarkBubbleSortMedium(b *testing.B) {
	input := generateSlice(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

// BenchmarkBubbleSortLarge benchmarks BubbleSort with a large input size
func BenchmarkBubbleSortLarge(b *testing.B) {
	input := generateSlice(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

// BenchmarkBubbleSortSorted benchmarks BubbleSort with an already sorted slice
func BenchmarkBubbleSortSorted(b *testing.B) {
	input := generateSlice(1000)
	sort.Ints(input)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

// BenchmarkBubbleSortReverse benchmarks BubbleSort with a reverse-sorted slice
func BenchmarkBubbleSortReverse(b *testing.B) {
	input := generateSlice(1000)
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}
