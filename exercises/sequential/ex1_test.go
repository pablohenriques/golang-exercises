package sequential

import "testing"

/*
	Faça um Programa que mostre a mensagem "Alo mundo" na tela.
*/

func message() string {
	return "Hello World"
}

func TestMessage(t *testing.T) {
	result := message()
	println(result)
	expected := "Hello World"

	if result != expected {
		t.Errorf("Different messages")
	}
}
