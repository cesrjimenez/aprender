package mathutil

import "testing"

func TestSuma(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{name: "dos números positivos", a: 2, b: 3, want: 5},
		{name: "suma con cero", a: 0, b: 7, want: 7},
		{name: "números negativos", a: -1, b: -5, want: -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // este subtest se ejecuta en paralelo
			got := Suma(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Suma(%d, %d) = %d; se esperaba %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
