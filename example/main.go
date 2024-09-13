package main

import (
	"fmt"

	"github.com/owlsome-official/toBahtText"
)

func main() {

	// EXAMPLE
	input := "1021"
	expected := "หนึ่งพันยี่สิบเอ็ด"
	actual := toBahtText.From(input)
	fmt.Printf("Input = %v | Expected = %v | Actual = %v\n", input, expected, actual)
	fmt.Printf("IS EQUAL = %v\n", expected == actual)

}
