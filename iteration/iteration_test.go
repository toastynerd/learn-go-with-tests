package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestAnotherRepeat(t *testing.T) {
	repeated := Repeat("b", 5)
	expected := "bbbbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestSpecificNumberRepeat(t *testing.T) {
	repeated := Repeat("b", 3)
	expected := "bbb"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
}
