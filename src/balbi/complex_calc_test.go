package balbi

import "testing"

func TestComplexCalcShouldReturnZero(t *testing.T) {
	if ComplexCalc(0, 0) != 0 {
		t.Error("Falha, esperava o zero")
	}
}

func TestComplexCalcShouldReturnSquareRootOf9(t *testing.T) {
	if ComplexCalc(9, 1) != 3 {
		t.Error("Falha, esperava o valor 3")
	}
}

func TestComplexCalcShouldIterateMoreThanOneAndNotReturnSquareRootOf25(t *testing.T) {
	if ComplexCalc(25, 2) == 5 {
		t.Error("Falha, esperava um valor diferente de 5")
	}
}

func TestComplexCalcShouldReturnZeroWhenNegativeValueOrTimePassed(t *testing.T) {
	if ComplexCalc(-25, 2) != 0 {
		t.Error("Falha, esperava 0 (zero) quando um valor negativo é passado")
	}
	if ComplexCalc(25, -30) != 0 {
		t.Error("Falha, esperava 0 (zero) quando um valor negativo é passado")
	}
}
