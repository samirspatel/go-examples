package main

import (
	"fmt"
	"log"

	"github.com/samirspatel/go-examples/greetings"
	"github.com/samirspatel/go-examples/modulo"
	"rsc.io/quote"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	m2, err := modulo.Hello("asdasd")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(m2)
	
	message, err := greetings.Hello("samir")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote.Go())
	fmt.Println(message)

	// pass in a slice of names
	names := []string{"max", "sluggy", "chungus"}

	//request messages
	messages, err := modulo.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// if no error returned, print messages
	fmt.Println(messages)
}
