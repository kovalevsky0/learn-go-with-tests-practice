package hello

import "fmt"

const greetingPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("Max"))
}
