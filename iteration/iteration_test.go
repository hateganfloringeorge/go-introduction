package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	result := Repeater("a", 5)
	expected := "aaaaa"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 100)
	}
}

func ExampleRepeater() {
	result := Repeater("y", 5)
	fmt.Println(result)
	// Output: yyyyy
}
