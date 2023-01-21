package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

var sleepLabel = "sleep"
var writeLabel = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepLabel)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeLabel)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints the result string", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeperWriter := &SpyCountdownOperations{}

		Countdown(&buffer, spySleeperWriter)

		result := buffer.String()
		expected := `3
2
1
Go!`

		if expected != result {
			t.Errorf("Expected %q but received %q", expected, result)
		}
	})
	t.Run("prints and sleeps in specific order", func(t *testing.T) {
		spySleeperWriter := &SpyCountdownOperations{}

		Countdown(spySleeperWriter, spySleeperWriter)

		expectedCalls := []string{
			writeLabel,
			sleepLabel,
			writeLabel,
			sleepLabel,
			writeLabel,
			sleepLabel,
			writeLabel,
		}

		if !reflect.DeepEqual(expectedCalls, spySleeperWriter.Calls) {
			t.Errorf("Expected calls %v but received %v", expectedCalls, spySleeperWriter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if sleepTime != spyTime.durationSlept {
		t.Errorf("Expected %d but received %d", sleepTime, spyTime.durationSlept)
	}
}
