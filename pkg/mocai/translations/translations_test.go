package translations

import (
	"testing"
)

func TestRegisterAndGet(t *testing.T) {
	Register("test_lang", map[string]string{
		"hello": "world",
	})
	val := Get("test_lang", "hello")
	if val != "world" {
		t.Errorf("expected 'world', got %q", val)
	}
}

func TestGetFallbackToEnUS(t *testing.T) {
	Register("en_us", map[string]string{
		"test_fallback_key": "fallback_value",
	})
	val := Get("nonexistent_lang", "test_fallback_key")
	if val != "fallback_value" {
		t.Errorf("expected 'fallback_value', got %q", val)
	}
}

func TestGetReturnsKeyWhenNotFound(t *testing.T) {
	val := Get("nonexistent_lang", "nonexistent_key")
	if val != "nonexistent_key" {
		t.Errorf("expected key returned as-is, got %q", val)
	}
}

func TestGetEmptyParams(t *testing.T) {
	val := Get("", "key")
	if val != "key" {
		t.Errorf("expected 'key', got %q", val)
	}
	val = Get("lang", "")
	if val != "" {
		t.Errorf("expected empty string, got %q", val)
	}
}

func TestRegisterListAndGetList(t *testing.T) {
	RegisterList("test_lang", map[string][]string{
		"colors": {"red", "green", "blue"},
	})
	val := GetList("test_lang", "colors")
	if len(val) != 3 {
		t.Errorf("expected 3 items, got %d", len(val))
	}
	if val[0] != "red" || val[1] != "green" || val[2] != "blue" {
		t.Errorf("unexpected values: %v", val)
	}
}

func TestGetListReturnsNilWhenNotFound(t *testing.T) {
	val := GetList("nonexistent_lang", "nonexistent_key")
	if val != nil {
		t.Errorf("expected nil, got %v", val)
	}
}

func TestGetListEmptyParams(t *testing.T) {
	val := GetList("", "key")
	if val != nil {
		t.Errorf("expected nil for empty lang, got %v", val)
	}
	val = GetList("lang", "")
	if val != nil {
		t.Errorf("expected nil for empty key, got %v", val)
	}
}

func TestGetRandomWithRandSource(t *testing.T) {
	RegisterList("test_rand_lang", map[string][]string{
		"fruits": {"apple", "banana", "cherry"},
	})
	rnd := &fixedRand{val: 1}
	val := GetRandom("test_rand_lang", "fruits", rnd)
	if val != "banana" {
		t.Errorf("expected 'banana' (index 1), got %q", val)
	}
}

func TestGetRandomWithNilRandSource(t *testing.T) {
	RegisterList("test_rand_lang2", map[string][]string{
		"animals": {"cat", "dog", "bird"},
	})
	val := GetRandom("test_rand_lang2", "animals", nil)

	if val != "cat" {
		t.Errorf("expected 'cat' (first element), got %q", val)
	}
}

func TestGetRandomReturnsKeyWhenListEmpty(t *testing.T) {
	val := GetRandom("nonexistent_lang", "nonexistent_key", nil)
	if val != "nonexistent_key" {
		t.Errorf("expected key returned as-is, got %q", val)
	}
}

func TestRegisterIgnoresEmptyLang(t *testing.T) {
	Register("", map[string]string{"key": "value"})
	RegisterList("", map[string][]string{"key": {"value"}})
}

func TestRegisterIgnoresNilMessages(t *testing.T) {
	Register("some_lang", nil)
	RegisterList("some_lang", nil)
}

func TestGetUFMapPtbr(t *testing.T) {
	ufMap := GetUFMap("ptbr")
	if ufMap == nil {
		t.Fatal("expected non-nil UF map for ptbr")
	}
	if len(ufMap) == 0 {
		t.Error("expected non-empty UF map for ptbr")
	}
}

func TestGetUFMapReturnsNilForUnsupportedLang(t *testing.T) {
	ufMap := GetUFMap("en_us")
	if ufMap != nil {
		t.Errorf("expected nil UF map for en_us, got %v", ufMap)
	}
}

func TestGetUFMapReturnsDefensiveCopy(t *testing.T) {
	ufMap1 := GetUFMap("ptbr")
	ufMap2 := GetUFMap("ptbr")

	for k := range ufMap1 {
		ufMap1[k] = "MODIFIED"
		break
	}

	for k, v := range ufMap2 {
		if v == "MODIFIED" {
			t.Errorf("GetUFMap returned a shared reference, key %q was modified", k)
		}
		break
	}
}

func TestDefaultRandSource(t *testing.T) {
	rnd := DefaultRandSource()
	if rnd == nil {
		t.Fatal("expected non-nil default rand source")
	}

	val := rnd.Intn(100)
	if val < 0 || val >= 100 {
		t.Errorf("expected value in [0, 100), got %d", val)
	}
}

func TestSafeRandSourceConcurrency(t *testing.T) {
	rnd := DefaultRandSource()
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				rnd.Intn(100)
			}
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestRegisterListPtbrDataExists(t *testing.T) {
	keys := []string{
		"person_first_name_male",
		"person_first_name_female",
		"person_last_name",
		"gender",
		"company_name",
		"address_street",
		"address_city",
		"address_state",
		"address_zip",
		"phone_area_code",
	}
	for _, key := range keys {
		val := GetList("ptbr", key)
		if len(val) == 0 {
			t.Errorf("expected non-empty list for ptbr/%s", key)
		}
	}
}

func TestRegisterPtbrConstantsExist(t *testing.T) {
	keys := []string{
		"brazilian_rg_state",
		"brazilian_rg_issuing_body",
		"invalid_cpf",
		"invalid_cnpj",
		"invalid_certificate",
	}
	for _, key := range keys {
		val := Get("ptbr", key)
		if val == key {
			t.Errorf("expected translated value for ptbr/%s, got key back", key)
		}
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
