package main

import (
	"fmt"
	"log"

	"github.com/mayyamark/go-demo/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including:
	log.SetPrefix("greetings: ") // the log entry prefix,
	log.SetFlags(0)              // a flag to disable printing the time, source file, and line number.

	// Request for a greeting message.
	// message, err := greetings.Hello(name) // This would return error:		greetings: passed an empty name to Hello
	message, err := greetings.Hello("Mayya")
	messages, err2 := greetings.Hellos([]string{"Ivan", "Mayya", "Martin"})

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(message)
	fmt.Println(messages)
	fmt.Println(quote.Go())
}
