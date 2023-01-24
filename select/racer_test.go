package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("less than 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		// defer - it will be executed at the end of the function (when test will be finished)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		result, _ := Racer(slowUrl, fastUrl)
		expected := fastUrl

		if expected != result {
			t.Errorf("Expected %q but received %q", expected, result)
		}
	})
	t.Run("more than 10 seconds", func(t *testing.T) {
		firstServer := makeDelayedServer(11 * time.Millisecond)
		secondServer := makeDelayedServer(12 * time.Millisecond)

		// defer - it will be executed at the end of the function (when test will be finished)
		defer firstServer.Close()
		defer secondServer.Close()

		firstUrl := firstServer.URL
		secondUrl := secondServer.URL

		_, err := ConfigurableRacer(firstUrl, secondUrl, 15*time.Millisecond)

		if err == nil {
			t.Errorf("Expected an error but received nil")
		}
	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	serverHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(duration)
		writer.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(serverHandler)
}
