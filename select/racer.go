package _select

import (
	"errors"
	"net/http"
	"time"
)

var ErrConnectionTimeout = errors.New("Connection is timeout")
var defaultTimeout = 10 * time.Second

func Racer(firstUrl, secondUrl string) (winner string, error error) {
	return ConfigurableRacer(firstUrl, secondUrl, defaultTimeout)
}

func ConfigurableRacer(firstUrl, secondUrl string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout):
		return "", ErrConnectionTimeout
	}
}

func ping(url string) chan struct{} {
	// the type of channel (struct) does not matter
	// (i've tried to change 'struct' to 'int' and it works the same)
	// the reason why choose struct because it's the smallest data type from the memory perspective
	channel := make(chan struct{})
	go func() {
		http.Get(url)
		close(channel)
	}()
	return channel
}
