package main

import (
	"fmt"

	"github.com/mayyamark/go-demo/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Mayya")
	fmt.Println(message)
}
