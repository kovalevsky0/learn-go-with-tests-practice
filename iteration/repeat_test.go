package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("Failed! Expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 3)
	fmt.Println(repeated)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}
