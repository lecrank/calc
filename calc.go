package main

import (
	"fmt"
	"log"

	"github.com/lecrank/calc/calculator"
	"github.com/lecrank/calc/parser"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Get math expression
	expression, err := parser.ParseExp()
	checkError(err)

	result, err := calculator.GetResult(expression)
	checkError(err)
	fmt.Println(result)
}
