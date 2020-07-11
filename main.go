package main

import (
	"fmt"
)

func main() {

	var valueA = 5
	var valueB = 5

	sumResult := sum(valueA, valueB)
	fmt.Printf("O resultado da soma de %v e %v é de %v !", valueA, valueB, sumResult)
}
