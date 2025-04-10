package nar

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
