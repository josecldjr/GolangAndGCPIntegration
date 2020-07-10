package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	valuesA, valuesB := [3]int{4, 5, 12}, [3]int{52, 32, 11}

	for i := range valuesA {
		testSingleSum(t, valuesA[i], valuesB[i])
	}
}

func testSingleSum(t *testing.T, a int, b int) {

	result := sum(a, b)
	expectedResult := a + b

	if result != expectedResult {
		t.Errorf("The sum (%v + %v = %v) is not correct", a, b, result)
	} else {
		t.Logf("The sum (%v + %v = %v) is correct", a, b, result)
	}
}
