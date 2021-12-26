package main

import (
	"reflect"
	"testing"
)

func AssertSumDeepEqual(got , want []int, t *testing.T){
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSumArr(t *testing.T){
	numbers := [5]int{1, 2, 3, 4, 5}
	got := SumArr(numbers)
	want := 15

	if got != want{
		t.Errorf("Got %d Want %d, given %v", got, want, numbers)
	}
}

func TestSumSlice(t *testing.T){
	numbers := []int{1, 2, 3}
	got := SumSlice(numbers)
	want := 6
	if got != want {
		t.Errorf("got %d want %d, for given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T){
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T){
	t.Run("make the sums of some slices", func(t * testing.T){
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		AssertSumDeepEqual(got, want, t)

	})
	t.Run("safely sum empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		AssertSumDeepEqual(got, want, t)
	})
}