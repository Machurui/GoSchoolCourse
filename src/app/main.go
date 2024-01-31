package main

import (
	"app/calculator"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(calculator.Sum(3, 5), "Version :", calculator.Version)

	fmt.Println(quote.Hello())
}
