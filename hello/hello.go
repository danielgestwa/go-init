package main

import (
	"fmt"
)

const greatings = "Hello "
const mark = "!"
const defaultGreatings = "Hello world!"

func Hello(name string) string {
	if name == "" {
		return defaultGreatings
	}
	return greatings + name + mark
}

func main() {
	fmt.Println(Hello("Andy"))
}
