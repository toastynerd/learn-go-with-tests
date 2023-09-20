package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d but got %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
}
