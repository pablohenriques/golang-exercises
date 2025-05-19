package sequential

import "testing"

// Faça um programa que peça as 4 notas bimestrais e mostre a média.

func gradePointAverage(values [4]float32) float32 {
	return sumValues(values) / 4
}

func sumValues(arrayValues [4]float32) float32 {
	var total float32 = 0
	for _, value := range arrayValues {
		total += value
	}
	return total
}

func TestGradePointAverage(t *testing.T) {
	var notes = [4]float32{10, 10, 10, 10}
	var result = gradePointAverage(notes)
	var expected float32 = 10

	if result != expected {
		t.Errorf("Average is not equal")
	}

}
