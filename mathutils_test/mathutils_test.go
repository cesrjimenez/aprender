package mathutils_test

import (
	"testear/mathutil"
	"testing"
)

func TestSuma(t *testing.T) {
	got := mathutil.Suma(2, 3)
	expected := 5
	
	if got != expected {
		t.Errorf("Add(2, 3) = %d; se esperaba %d", got, expected)
	}
}
