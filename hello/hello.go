package main

import (
	"fmt"
	"log"

	"github.com/mayyamark/go-demo/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request for a greeting message.
	// message, err := greetings.Hello(name) // This would return error:		greetings: passed an empty name
	message, err := greetings.Hello("Mayya")

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	// And the quote as well.
	fmt.Println(quote.Go())
}
