package cpf

import (
	"errors"
	"testing"
)

func TestGenerateCPF(t *testing.T) {
	cpfVal, err := NewCPF("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cpfVal.Number) != 11 {
		t.Errorf("expected 11 digits, got %d: %s", len(cpfVal.Number), cpfVal.Number)
	}
}

func TestGenerateCPFFormatted(t *testing.T) {
	cpfVal, err := NewCPF("ptbr", true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cpfVal.Number) != 14 {
		t.Errorf("expected 14 chars formatted, got %d: %s", len(cpfVal.Number), cpfVal.Number)
	}
}

func TestGeneratedCPFIsValid(t *testing.T) {
	for i := 0; i < 100; i++ {
		cpfVal, err := NewCPF("ptbr", false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if !ValidateCPF(cpfVal.Number) {
			t.Errorf("generated CPF is invalid: %s", cpfVal.Number)
		}
	}
}

func TestGeneratedCPFFormattedIsValid(t *testing.T) {
	for i := 0; i < 100; i++ {
		cpfVal, err := NewCPF("ptbr", true, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if !ValidateCPF(cpfVal.Number) {
			t.Errorf("generated formatted CPF is invalid: %s", cpfVal.Number)
		}
	}
}

func TestCPFDeterministic(t *testing.T) {
	rnd1 := &fixedRand{sequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, idx: 0}
	rnd2 := &fixedRand{sequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, idx: 0}

	cpf1, err := NewCPF("ptbr", false, rnd1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cpf2, err := NewCPF("ptbr", false, rnd2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cpf1.Number != cpf2.Number {
		t.Errorf("expected deterministic output, got %s and %s", cpf1.Number, cpf2.Number)
	}
}

func TestCPFErrorIsSentinel(t *testing.T) {
	err := ErrInvalidCPF
	if !errors.Is(err, ErrInvalidCPF) {
		t.Error("expected ErrInvalidCPF to match with errors.Is")
	}
}

// fixedRand is a test helper that returns values from a sequence.
type fixedRand struct {
	sequence []int
	idx      int
}

func (f *fixedRand) Intn(n int) int {
	if len(f.sequence) == 0 {
		return 0
	}

	val := f.sequence[f.idx%len(f.sequence)]
	f.idx++
	if val >= n {
		return val % n
	}
	return val
}
