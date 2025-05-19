package structure_decision

import "testing"

// Faça um Programa que peça um valor e mostre na tela se o valor é positivo ou negativo.

func checkNumber(value int) string {
	var result string

	if value > 0 {
		result = "positivo"
	} else if value < 0 {
		result = "negativo"
	} else {
		result = "zero"
	}
	return result

}

func TestCheckNumberPositive(t *testing.T) {
	var actual = checkNumber(10)
	var expected = "positivo"

	if expected != actual {
		t.Errorf("Status are different")
	}
}

func TestCheckNumberNegative(t *testing.T) {
	var actual = checkNumber(-1)
	var expected = "negativo"

	if expected != actual {
		t.Errorf("Status are different")
	}
}

func TestCheckNumberZero(t *testing.T) {
	var actual = checkNumber(0)
	var expected = "zero"

	if expected != actual {
		t.Errorf("Status are different")
	}
}
