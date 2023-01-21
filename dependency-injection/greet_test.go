package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "John Smith")

	result := buffer.String()
	expected := "Hello, John Smith!"

	if expected != result {
		t.Errorf("Expected %q but received %q", expected, result)
	}
}
