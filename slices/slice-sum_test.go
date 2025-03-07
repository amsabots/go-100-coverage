package slices

import (
	"reflect"
	"testing"
)



func TestSliceSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}
		got := SliceSum(numbers)
		expected := 15
		if got != expected {
			t.Errorf("got %d expected %d give %v", got, expected, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1,2,3}
		got := SliceSum(numbers)
		expected := 6
		if got != expected {
			t.Errorf("got %d expected %d give %v", got, expected, numbers)
		}
	})
}

func BenchmarkSSliceSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{1,2,3,4,5,6,7,8,8,32,353,232,42,3,45,4}
		 SliceSum(numbers)
	}
}


func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{2,3})
	expected := []int{3, 5}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected: %v", got, expected)
	}
}


func TestSumAllTails(t *testing.T) {
	checkArraySums := func (t testing.TB, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, expected: %v", got, want)
		}
	}
	t.Run("make sum of non empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{2,3}, []int{4,5,6})
		expected := []int{0, 3, 11}
		checkArraySums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2,3})
		expected := []int{0, 3}
		checkArraySums(t, got, expected)
	})
	
}