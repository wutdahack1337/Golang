package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!!!")

	fmt.Println()
	// ========================================================================================================================================================

	fmt.Println(quote.Go())

	fmt.Println()
	// ========================================================================================================================================================

	// Set properties of the predefined Logger, including the log entry prefix and a flag to disable printing the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Minh", "Khang", "Bon", "Ben"}

	// Request greeting messages for the names.
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of messages to the console.
	fmt.Println(message)

	fmt.Println()
	// ========================================================================================================================================================

	fmt.Println(reverse.String("Hello"))
}
