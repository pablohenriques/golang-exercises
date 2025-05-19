package structure_decision

import "testing"

// Faça um Programa que peça dois números e imprima o maior deles

func biggerNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestBiggerNumber(t *testing.T) {
	actual := biggerNumber(10, 20)
	expected := 20

	if expected != actual {
		t.Errorf("Numbers are different")
	}

}
