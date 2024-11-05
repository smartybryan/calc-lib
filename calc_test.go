package calc_lib

import (
	"fmt"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 1},
		{a: 1, b: 1, want: 2},
		{a: 2, b: 3, want: 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			addition := &Addition{}
			got := addition.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: -1},
		{a: 1, b: 1, want: 0},
		{a: 2, b: 3, want: -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			operation := &Subtraction{}
			got := operation.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 0},
		{a: 1, b: 1, want: 1},
		{a: 2, b: 3, want: 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			operation := &Multiplication{}
			got := operation.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{a: 0, b: 5, want: 0},
		{a: 0, b: 1, want: 0},
		{a: 1, b: 1, want: 1},
		{a: 6, b: 3, want: 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			operation := &Division{}
			got := operation.Calculate(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_ByZeroPanics(t *testing.T) {
	division := &Division{}
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Division by zero did not panic")
		}
	}()
	division.Calculate(0, 0)
}
