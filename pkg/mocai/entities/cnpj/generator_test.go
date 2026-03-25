package cnpj

import (
	"testing"
)

func TestGenerateCNPJ(t *testing.T) {
	cnpjVal, err := GenerateCNPJ(false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cnpjVal) != 14 {
		t.Errorf("expected 14 digits, got %d: %s", len(cnpjVal), cnpjVal)
	}
}

func TestGenerateCNPJFormatted(t *testing.T) {
	cnpjVal, err := GenerateCNPJ(true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cnpjVal) != 18 {
		t.Errorf("expected 18 chars formatted, got %d: %s", len(cnpjVal), cnpjVal)
	}
}

func TestGeneratedCNPJIsValid(t *testing.T) {
	for i := 0; i < 100; i++ {
		cnpjVal, err := GenerateCNPJ(false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		if !ValidateCNPJ(cnpjVal) {
			t.Errorf("generated CNPJ is invalid: %s", cnpjVal)
		}
	}
}

func TestGeneratedCNPJFormattedIsValid(t *testing.T) {
	for i := 0; i < 100; i++ {
		cnpjVal, err := GenerateCNPJ(true, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}
		if !ValidateCNPJ(cnpjVal) {
			t.Errorf("generated formatted CNPJ is invalid: %s", cnpjVal)
		}
	}
}

func TestCNPJDeterministic(t *testing.T) {
	rnd1 := &fixedRand{sequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 1}, idx: 0}
	rnd2 := &fixedRand{sequence: []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 1}, idx: 0}

	cnpj1, err := GenerateCNPJ(false, rnd1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cnpj2, err := GenerateCNPJ(false, rnd2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cnpj1 != cnpj2 {
		t.Errorf("expected deterministic output, got %s and %s", cnpj1, cnpj2)
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
