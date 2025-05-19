package sequential

import "testing"

/**
Faça um programa que peça dois números e imprima a soma.
*/

func sumNumbers(first, second int) int {
	return first + second
}

func TestSumNumber(t *testing.T) {
	var first = 1
	var second = 1
	result := sumNumbers(first, second)
	expected := 2

	if result != expected {
		t.Errorf("The sum result is different")
	}
}
