package main

import "fmt"

const englishHelloPrix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrix + name
}

func main() {
	fmt.Print(Hello("Jake"))
}
