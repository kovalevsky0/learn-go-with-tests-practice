package hello

import "testing"

func TestHello(t *testing.T) {
	assert := func(t testing.TB, expectedValue string, resultValue string) {
		t.Helper()

		if resultValue != expectedValue {
			t.Errorf("expected `%s` but got `%s`", expectedValue, resultValue)
		}
	}

	t.Run("should return `Hello, <name>!`", func(t *testing.T) {
		result := Hello("Max")
		expectedValue := "Hello, Max!"

		assert(t, expectedValue, result)
	})

	t.Run("should return `Hello, World!`", func(t *testing.T) {
		result := Hello("")
		expectedValue := "Hello, World!"

		assert(t, expectedValue, result)
	})
}
