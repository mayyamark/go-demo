package greetings

import (
	"errors"
	"fmt"
	"math/rand" // This package generates a random number.
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("passed an empty name")
	}

	return fmt.Sprintf(randomFormat(), name), nil
}

// The init function seeds  the rand package with the current time.
// Go executes init functions automatically at program startup, after global variables have been initialized.
// Starts with lowercase letter => it's not exported.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// The randomFormat function returns one of a set of greeting messages.
// The returned message is selected at random.
// Starts with lowercase letter => it's not exported.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
