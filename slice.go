package compose

// Map is similar to Array.Map in javascript
func Map[T any, R any](input []T, filter func(T) R) []R {
	res := make([]R, 0, len(input))
	for _, item := range input {
		res = append(res, filter(item))
	}
	return res
}

// Filter filters a slice according to the predicate provided.
func Filter[T any](input []T, predicate func(item T) bool) []T {
	res := make([]T, 0, len(input))
	for _, item := range input {
		if predicate(item) {
			res = append(res, item)
		}
	}
	return res
}

// Reduce aggregates the result by running through all values in the input.
func Reduce[T any, R any](input []T, aggregator func(aggregate R, item T) R, first R) R {
	for _, item := range input {
		first = aggregator(first, item)
	}
	return first
}

// Unique returns the slice of unique values provided input slice.
func Unique[T comparable](input []T) []T {
	res := make([]T, 0, len(input))
	uniqMap := make(map[T]struct{}, len(input))

	for _, item := range input {
		if _, ok := uniqMap[item]; ok {
			continue
		}
		uniqMap[item] = struct{}{}
		res = append(res, item)
	}
	return res
}

// Clone returns a clone of the input.
func Clone[T any](input []T) []T { return Filter(input, func(T) bool { return true }) }

// Find returns the first element in the input that satisfies the predicate.
func Find[T any](input []T, predicate func(T) bool) (T, bool) {
	res := Filter(input, predicate)
	if len(res) == 0 {
		var r T
		return r, false
	}
	return res[0], true
}
