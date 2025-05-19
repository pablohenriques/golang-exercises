package sequential

import "testing"

// Faça um programa que converta metros para centímetros.

func centimetersToMeters(value int) float32 {
	result := float32(value / 100)
	return result
}

func TestCentimetersToMeters(t *testing.T) {
	actual := centimetersToMeters(100)
	expected := float32(1)
	if actual != expected {
		t.Errorf("Differents messages")
	}
}
