package main

import (
	"fmt"

	"github.com/samirspatel/go-examples/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Samir"))
}
