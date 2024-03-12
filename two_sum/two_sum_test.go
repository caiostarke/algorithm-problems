package two_sum_test

import (
	"algorithm-problems/two_sum"
	"testing"
)

func Test_Two_Sum(t *testing.T) {
	expected := 9
	val1, val2 := two_sum.Two_sum(9, []int{2, 3, 5, 7, 9, 11})

	if val1+val2 != expected {
		t.Errorf("expecting %d, got %d", expected, val1+val2)
	}

	expected = -1
	val1, val2 = two_sum.Two_sum(1000, []int{2, 3, 5, 7, 9, 11})

	if val1 != expected && val2 != expected {
		t.Errorf("expecting %d, got %d and %d", expected, val1, val2)
	}
}

func Test_Two_Sum_HashTable(t *testing.T) {
	expected := 9
	val1, val2 := two_sum.Two_Sum_HashTable(9, []int{2, 3, 5, 7, 9, 11})

	if val1+val2 != expected {
		t.Errorf("expecting %d, got %d", expected, val1+val2)
	}

	expected = -1
	val1, val2 = two_sum.Two_Sum_HashTable(1000, []int{2, 3, 5, 7, 9, 11})

	if val1 != expected && val2 != expected {
		t.Errorf("expecting %d, got %d and %d", expected, val1, val2)
	}
}
