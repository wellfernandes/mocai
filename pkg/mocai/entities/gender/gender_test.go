package gender

import (
	"errors"
	"testing"
)

func TestGenerateGender(t *testing.T) {
	g, err := NewGender("ptbr", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if g.Identity == "" {
		t.Error("gender identity is empty")
	}
}

func TestGenerateGenderUnsupportedLang(t *testing.T) {
	_, err := NewGender("xyz", nil)
	if err == nil {
		t.Fatal("expected error for unsupported language, got nil")
	}

	if !errors.Is(err, ErrNoGenderData) {
		t.Errorf("expected ErrNoGenderData, got: %v", err)
	}
}

func TestGenerateMultipleGenders(t *testing.T) {
	for i := 0; i < 50; i++ {
		g, err := NewGender("ptbr", nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if g.Identity == "" {
			t.Errorf("gender identity is empty on iteration %d", i)
		}
	}
}
