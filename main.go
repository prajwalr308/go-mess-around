package main

import (
	"fmt"
	"log"
	"mess-around/hello"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := hello.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println("Hello, World!")
}
