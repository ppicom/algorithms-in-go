package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	keys := []int{5, 6, 47, 2, 9, 89}
	want := []int{2, 5, 6, 9, 47, 89}

	sorted := InsertionSort(keys)

	if !slicesAreEqual(keys, sorted) {
		t.Errorf("got %d want %d given, %v", sorted, want, keys)
	}
}

func slicesAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}