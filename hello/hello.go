package main

import (
	"fmt"

	"github.com/duexcoast/greetings"
)

func main() {
	// get a greeting message and print it
	message := greetings.Hello("Gladyss")
	fmt.Println(message)
}
