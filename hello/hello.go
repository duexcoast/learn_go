package main

import (
	"fmt"
	"log"

	"github.com/duexcoast/greetings"
)

func main() {
	// set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, sourcefile, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Garth", "Garret", "Gladyss"}

	// Request a greetings message
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message
	// to the console.

	fmt.Println(message)
}
