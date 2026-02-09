package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %v, want %v", got, 5)
	}
	if got := Add(-1.5, 1.5); got != 0 {
		t.Fatalf("Add(-1.5,1.5) = %v, want %v", got, 0)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(5, 3); got != 2 {
		t.Fatalf("Sub(5,3) = %v, want %v", got, 2)
	}
	if got := Sub(0, 1.5); got != -1.5 {
		t.Fatalf("Sub(0,1.5) = %v, want %v", got, -1.5)
	}
}

func TestMul(t *testing.T) {
	if got := Mul(2, 3); got != 6 {
		t.Fatalf("Mul(2,3) = %v, want %v", got, 6)
	}
	if got := Mul(-2, 2.5); got != -5 {
		t.Fatalf("Mul(-2,2.5) = %v, want %v", got, -5)
	}
}

func TestDiv(t *testing.T) {
	if got, err := Div(10, 2); err != nil || got != 5 {
		t.Fatalf("Div(10,2) = %v, %v; want 5, nil", got, err)
	}
	if _, err := Div(1, 0); err == nil {
		t.Fatalf("Div(1,0) expected error, got nil")
	}
}

