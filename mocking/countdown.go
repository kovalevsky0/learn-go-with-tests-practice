package mocking

import (
	"fmt"
	"io"
	"time"
)

type ConfigurableSleeper struct {
	Duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.Duration)
}

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}

	fmt.Fprint(writer, "Go!")
}
