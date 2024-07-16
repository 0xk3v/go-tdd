package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 3 numbers", func(t *testing.T) {
		numbers := []int{3, 4, 6}

		got := sum(numbers)
		want := 13
		if got != want {
			t.Errorf("Got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{3, 10, 7, 4, 6}

		got := sum(numbers)
		want := 30
		if got != want {
			t.Errorf("Got %d, want %d, given %v", got, want, numbers)
		}
	})
}
