package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!\n", name)
}
