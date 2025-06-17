package nar

import "iter"

// Map applies a function to each element of a slice and returns a new slice with the results
func Map[T any, U any](s []T, f func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// Filter returns a new slice containing only the elements of the original slice for which the function returns true
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// IndexBy transforms a slice into a map by applying a key selector function to each element of the slice.
// The keys in the resulting map are obtained from the selector, and the values are the corresponding slice elements.
func IndexBy[T any, K comparable](s []T, f func(T) K) map[K]T {
	result := make(map[K]T)
	for _, v := range s {
		result[f(v)] = v
	}
	return result
}

// Zip combines two slices into a sequence of pairs.
// Iteration stops at the end of the shorter slice.
func Zip[T1 any, T2 any](s1 []T1, s2 []T2) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
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

// ZipLongest combines two slices into a sequence of pairs.
// Iteration continues until the longer slice ends, yielding zero values for the shorter slice.
func ZipLongest[T1 any, T2 any](s1 []T1, s2 []T2) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		maxLen := len(s1)
		if len(s2) > maxLen {
			maxLen = len(s2)
		}
		for i := 0; i < maxLen; i++ {
			var v1 T1
			if i < len(s1) {
				v1 = s1[i]
			}
			var v2 T2
			if i < len(s2) {
				v2 = s2[i]
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}
