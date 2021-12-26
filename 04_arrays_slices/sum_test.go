package main
import "testing"

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