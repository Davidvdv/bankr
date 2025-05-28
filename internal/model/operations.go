package model

func GroupBy[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)
	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	transformed := make([]R, len(items))
	for i, item := range items {
		transformed[i] = transform(item)
	}
	return transformed
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Sum[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](items []T) T {
	var total T
	for _, item := range items {
		total += item
	}
	return total
}
