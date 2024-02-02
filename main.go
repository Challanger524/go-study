package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!", "Golang goes to the Clip 😄")

	fmt.Print("QUOTE: ")
	fmt.Println(quote.Go())
}
