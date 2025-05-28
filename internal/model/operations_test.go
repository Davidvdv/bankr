package model

import (
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	t.Run("group integers by even/odd", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		result := GroupBy(numbers, func(n int) bool { return n%2 == 0 })

		expected := map[bool][]int{
			true:  {2, 4, 6},
			false: {1, 3, 5},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("group strings by length", func(t *testing.T) {
		words := []string{"cat", "dog", "elephant", "bird", "mouse"}
		result := GroupBy(words, func(s string) int { return len(s) })

		expected := map[int][]string{
			3: {"cat", "dog"},
			8: {"elephant"},
			4: {"bird"},
			5: {"mouse"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int
		result := GroupBy(empty, func(n int) int { return n })

		expected := make(map[int][]int)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single item", func(t *testing.T) {
		items := []string{"hello"}
		result := GroupBy(items, func(s string) string { return s })

		expected := map[string][]string{
			"hello": {"hello"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("transform integers to strings", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Map(numbers, func(n int) string { return string(rune(n + '0')) })

		expected := []string{"1", "2", "3", "4", "5"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("square numbers", func(t *testing.T) {
		numbers := []int{2, 3, 4}
		result := Map(numbers, func(n int) int { return n * n })

		expected := []int{4, 9, 16}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int
		result := Map(empty, func(n int) int { return n * 2 })

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got %v", result)
		}
	})

	t.Run("string to length", func(t *testing.T) {
		words := []string{"hello", "world", "go"}
		result := Map(words, func(s string) int { return len(s) })

		expected := []int{5, 5, 2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
		result := Filter(numbers, func(n int) bool { return n%2 == 0 })

		expected := []int{2, 4, 6, 8}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("filter strings by length", func(t *testing.T) {
		words := []string{"cat", "elephant", "dog", "bird"}
		result := Filter(words, func(s string) bool { return len(s) <= 3 })

		expected := []string{"cat", "dog"}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("no matches", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7}
		result := Filter(numbers, func(n int) bool { return n%2 == 0 })

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got %v", result)
		}
	})

	t.Run("all matches", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8}
		result := Filter(numbers, func(n int) bool { return n%2 == 0 })

		if !reflect.DeepEqual(result, numbers) {
			t.Errorf("Expected %v, got %v", numbers, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int
		result := Filter(empty, func(n int) bool { return n > 0 })

		if len(result) != 0 {
			t.Errorf("Expected empty slice, got %v", result)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("sum integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("sum float64", func(t *testing.T) {
		numbers := []float64{1.5, 2.5, 3.0}
		result := Sum(numbers)
		expected := 7.0

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("sum float32", func(t *testing.T) {
		numbers := []float32{1.1, 2.2, 3.3}
		result := Sum(numbers)
		expected := float32(6.6)

		// Use approximate comparison for floats
		if result < expected-0.001 || result > expected+0.001 {
			t.Errorf("Expected approximately %v, got %v", expected, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var empty []int
		result := Sum(empty)
		expected := 0

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single element", func(t *testing.T) {
		numbers := []int{42}
		result := Sum(numbers)
		expected := 42

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		numbers := []int{-1, -2, -3, 4, 5}
		result := Sum(numbers)
		expected := 3

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("uint types", func(t *testing.T) {
		numbers := []uint{1, 2, 3, 4}
		result := Sum(numbers)
		expected := uint(10)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("int64 types", func(t *testing.T) {
		numbers := []int64{100, 200, 300}
		result := Sum(numbers)
		expected := int64(600)

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
