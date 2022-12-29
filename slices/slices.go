package slices

// All returns true if all elements match the given predicate
func All[T any](slice []T, p func(T) bool) bool {
	for _, e := range slice {
		if !p(e) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element matches the given predicate
func Any[T any](slice []T, p func(T) bool) bool {
	for _, e := range slice {
		if p(e) {
			return true
		}
	}
	return false
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

// Distinct returns a stream consisting of the distinct elements.
func Distinct[T comparable](slice []T) []T {
	ret := make([]T, 0)
	exists := make(map[T]bool)

	for _, e := range slice {
		_, ok := exists[e]
		if ok {
			continue
		}
		exists[e] = true
		ret = append(ret, e)
	}
	return ret
}

// DistinctBy returns a stream consisting of the distinct elements that match the provided predicate.
func DistinctBy[T any, R comparable](slice []T, p func(T) R) []T {
	ret := make([]T, 0)
	exists := make(map[R]bool)

	for _, e := range slice {
		r := p(e)
		_, ok := exists[r]
		if ok {
			continue
		}
		exists[r] = true
		ret = append(ret, e)
	}
	return ret
}

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
func Map[T any, R any](slice []T, p func(T) R) []R {
	res := make([]R, len(slice))
	for i, e := range slice {
		res[i] = p(e)
	}
	return res
}

// NoneMatch returns whether no elements match the provided predicate.
func NoneMatch[T any](slice []T, p func(T) bool) bool {
	for _, e := range slice {
		if p(e) {
			return false
		}
	}
	return true
}
