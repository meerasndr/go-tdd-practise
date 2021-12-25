package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T){
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Errorf("Expected %q Got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat(){
	repeatedstring := Repeat("a", 5)
	fmt.Println(repeatedstring)
	//Output: aaaaa
}