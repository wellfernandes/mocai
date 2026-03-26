package mocai

import (
	"testing"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func TestNewMockerDefaults(t *testing.T) {
	m := NewMocker()
	if m.Language() != "ptbr" {
		t.Errorf("expected default language 'ptbr', got %q", m.Language())
	}
}

func TestNewMockerWithOptions(t *testing.T) {
	m := NewMocker(
		WithLanguage("en_us"),
		WithFormatted(true),
	)
	if m.Language() != "en_us" {
		t.Errorf("expected language 'en_us', got %q", m.Language())
	}
}

func TestMockerImplementsMockGenerator(t *testing.T) {
	var _ MockGenerator = NewMocker()
}

func TestNewMockerWithRandSource(t *testing.T) {
	rnd := translations.NewSafeRandSource(&fixedRand{val: 5})
	m := NewMocker(WithRandSource(rnd))
	if m == nil {
		t.Fatal("expected non-nil Mocker")
	}
}

func TestNewPerson(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	p, err := m.NewPerson()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if p.FirstNameMale == "" || p.LastName == "" {
		t.Errorf("person has empty fields: %+v", p)
	}
}

func TestNewAddress(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	addr, err := m.NewAddress()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if addr.Street == "" || addr.City == "" || addr.State == "" {
		t.Errorf("address has empty fields: %+v", addr)
	}
}

func TestNewCompany(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	comp, err := m.NewCompany()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if comp.BrazilianCompany.Name == "" || comp.BrazilianCompany.CNPJ == "" {
		t.Errorf("company has empty fields: %+v", comp)
	}
}

func TestNewCPF(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	cpfVal, err := m.NewCPF()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cpfVal.Number == "" {
		t.Error("CPF number is empty")
	}
}

func TestNewCertificate(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	cert, err := m.NewCertificate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cert.Brazil == nil {
		t.Error("certificate Brazil is nil")
	}
}

func TestNewNationalID(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	nid, err := m.NewNationalID()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if nid.BrazilianRG.Number == "" {
		t.Error("RG number is empty")
	}
}

func TestNewVoteRegistration(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	vr, err := m.NewVoteRegistration()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if vr.BrazilianVoteRegistration.Number == "" {
		t.Error("vote registration number is empty")
	}
}

func TestNewGender(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	g, err := m.NewGender()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if g.Identity == "" {
		t.Error("gender identity is empty")
	}
}

func TestNewPhone(t *testing.T) {
	m := NewMocker(WithLanguage("ptbr"))
	ph, err := m.NewPhone()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ph.AreaCode == "" || ph.Number == "" {
		t.Errorf("phone has empty fields: %+v", ph)
	}
}

// fixedRand is a test helper that always returns the same value.
type fixedRand struct {
	val int
}

func (f *fixedRand) Intn(n int) int {
	if f.val >= n {
		return 0
	}

	return f.val
}
