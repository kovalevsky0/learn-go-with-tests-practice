package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func mockWebsiteChecker(url string) bool {
	if url == "https://duckduckgo.com" {
		return true
	}
	return false
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://duckduckgo.com",
		"https://reddit.com",
	}

	t.Run("normal speed", func(t *testing.T) {
		result := CheckWebsites(mockWebsiteChecker, websites)
		expected := map[string]bool{
			"https://duckduckgo.com": true,
			"https://google.com":     false,
			"https://reddit.com":     false,
		}

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v but received %v", expected, result)
		}
	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "someurl"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
