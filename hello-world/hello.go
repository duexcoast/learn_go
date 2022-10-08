// declare a main package, packages are a way
// to group functions, and it's made up of all
// the files in the same directory.
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}