package sequential

import (
	"fmt"
	"testing"
)

/*
	Faça um Programa que peça um número e então mostre a mensagem O número informado foi [número].
*/

func numberMessage(number int) string {
	return fmt.Sprintf("The number %d", number)
}

func TestNumberMessage(t *testing.T) {
	result := numberMessage(10)
	expected := "The number 10"

	if result != expected {
		t.Errorf("Differents messages")
	}
}
