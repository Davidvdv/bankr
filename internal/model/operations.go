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

func Sum[T any](items []T, getValue func(T) float64) float64 {
	total := 0.0
	for _, item := range items {
		total += getValue(item)
	}
	return total
}