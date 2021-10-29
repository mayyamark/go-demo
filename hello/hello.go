package main

import (
	"fmt"

	"github.com/mayyamark/go-demo/greetings"
	"rsc.io/quote"
)

func main() {
	// Get a greeting message and print it.
	name := "Mayya"
	message := greetings.Hello(name)
	fmt.Println(message)

	fmt.Println(quote.Go())
}
