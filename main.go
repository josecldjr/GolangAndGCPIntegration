package main

import (
	"fmt"
)

func main() {

	var valueA = 23
	var valueB = 62

	sumResult := sum(valueA, valueB)
	fmt.Printf("O resultado da soma de %v e %v é de %v", valueA, valueB, sumResult)
}
