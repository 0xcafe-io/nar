package nar

import (
	"iter"
)

// Map applies a function to each element of a slice and returns a new slice with the results.
// Example: Map([]int{1, 2, 3}, func(x int) int { return x * x })
// returns []int{1, 4, 9}
func Map[S1 ~[]E1, E1 any, E2 any](s S1, f func(E1) E2) []E2 {
	if s == nil {
		return nil
	}
	result := make([]E2, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Filter returns a new slice containing only the elements of the original slice for which the f returns true
// Order is preserved.
// Example: Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
// returns []int{2, 4}
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	var result S
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Find returns the first element of the slice that satisfies the provided function f and a bool indicating whether such an element was found.
// Example: Find([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 }) returns (2, true)
// Example2: Find([]int{1, 3, 5}, func(x int) bool { return x%2 == 0 }) returns (0, false)
func Find[S ~[]E, E any](s S, f func(E) bool) (E, bool) {
	for _, v := range s {
		if f(v) {
			return v, true
		}
	}
	var zero E
	return zero, false
}

// IndexBy transforms a slice into a map by applying a key selector function to each element of the slice.
// The keys in the resulting map are obtained from the selector, and the values are the corresponding slice elements.
//
// Example: IndexBy([]string{"apple", "banana"}, func(s string) int { return len(s) })
// returns map[int]string{5: "apple", 6: "banana"}
//
// Example2: IndexBy([]string{"apple", "apricot"}, func(s string) byte { return s[0] })
// returns map[byte]string{'a': "apricot"} (the last value overrides the previous one for the same key)
func IndexBy[S ~[]E, E any, K comparable](s S, f func(E) K) map[K]E {
	result := make(map[K]E, len(s))
	for _, v := range s {
		result[f(v)] = v
	}
	return result
}

// GroupBy groups the elements of a slice into a map according to a key selector function.
// The keys in the resulting map are obtained from the selector, and the values are slices of elements that correspond to each key.
// Example: GroupBy([]string{"apple", "apricot", "banana"}, func(s string) byte { return s[0] })
// returns map[byte][]string{'a': {"apple", "apricot"}, 'b': {"banana"}}
func GroupBy[S ~[]E, E any, K comparable](s S, f func(E) K) map[K]S {
	result := make(map[K]S)
	for _, v := range s {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Zip combines two slices into a sequence of pairs and returns an iterator over these pairs.
// Iteration stops at the end of the shorter slice.
// Example: for x1, x2 := range Zip([]int{1, 2, 3}, []string{"a", "b"})
// yields (1, "a"), (2, "b") and then stops.
func Zip[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s1 S1, s2 S2) iter.Seq2[E1, E2] {
	return func(yield func(E1, E2) bool) {
		minLen := len(s1)
		if len(s2) < minLen {
			minLen = len(s2)
		}
		for i := 0; i < minLen; i++ {
			if !yield(s1[i], s2[i]) {
				return
			}
		}
	}
}

// ZipLongest combines two slices into a sequence of pairs and returns an iterator over these pairs.
// Iteration continues until the longer slice ends, yielding zero values for the shorter slice.
// Example: for x1, x2 := range ZipLongest([]int{1, 2, 3}, []string{"a", "b"})
// yields (1, "a"), (2, "b"), (3, "") and then stops.
func ZipLongest[S1 ~[]E1, S2 ~[]E2, E1 any, E2 any](s1 S1, s2 S2) iter.Seq2[E1, E2] {
	return func(yield func(E1, E2) bool) {
		maxLen := len(s1)
		if len(s2) > maxLen {
			maxLen = len(s2)
		}
		for i := 0; i < maxLen; i++ {
			var v1 E1
			if i < len(s1) {
				v1 = s1[i]
			}
			var v2 E2
			if i < len(s2) {
				v2 = s2[i]
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}
