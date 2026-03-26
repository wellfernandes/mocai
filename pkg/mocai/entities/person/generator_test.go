package person

import (
	"errors"
	"testing"
)

func TestGeneratePerson(t *testing.T) {
	p, err := NewPerson("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.FirstNameMale == "" {
		t.Error("FirstNameMale is empty")
	}

	if p.FirstNameFemale == "" {
		t.Error("FirstNameFemale is empty")
	}

	if p.LastName == "" {
		t.Error("LastName is empty")
	}

	if p.Gender == nil {
		t.Error("Gender is nil")
	}

	if p.CPF == nil {
		t.Error("CPF is nil")
	}

	if p.Age < 18 || p.Age > 87 {
		t.Errorf("Age out of expected range [18, 87]: %d", p.Age)
	}
}

func TestGeneratePersonFormatted(t *testing.T) {
	p, err := NewPerson("ptbr", true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(p.CPF.Number) != 14 {
		t.Errorf("expected formatted CPF with 14 chars, got %d: %s", len(p.CPF.Number), p.CPF.Number)
	}
}

func TestGeneratePersonUnsupportedLang(t *testing.T) {
	_, err := NewPerson("xyz", false, nil)
	if err == nil {
		t.Fatal("expected error for unsupported language, got nil")
	}

	if !errors.Is(err, ErrNoFirstNames) {
		t.Errorf("expected ErrNoFirstNames, got: %v", err)
	}
}

func TestGenerateMultiplePersons(t *testing.T) {
	for i := 0; i < 50; i++ {
		p, err := NewPerson("ptbr", false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if p.FirstNameMale == "" || p.LastName == "" {
			t.Errorf("person has empty fields on iteration %d: %+v", i, p)
		}
	}
}
