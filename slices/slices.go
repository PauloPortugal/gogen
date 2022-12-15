package slices

// Filter returns a slice consisting of the elements that match the given predicate.
func Filter[T any](slice []T, p func(T) bool) []T {
	var res []T
	for _, e := range slice {
		if p(e) {
			res = append(res, e)
		}
	}
	return res
}

// Map returns a slice consisting of the results of applying the given function to the elements of the input slice.
func Map[T any, R any](slice []T, f func(T) R) []R {
	res := make([]R, len(slice))
	for i, e := range slice {
		res[i] = f(e)
	}
	return res
}

// Count returns the count of elements in this stream.
func Count[T any](slice []T) int {
	return len(slice)
}

// Contains returns whether elements of this slice match the provided predicate.
func Contains[T any](slice []T, p func(T) bool) bool {
	for _, e := range slice {
		if p(e) {
			return true
		}
	}
	return false
}