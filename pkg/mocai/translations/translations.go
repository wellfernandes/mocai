package translations

import (
	"sync"

	address_ptbr "github.com/brazzcore/mocai/pkg/mocai/entities/address/mocks/ptbr"
)

// GetUFMap Returns the mapping of states for the specified language, if available
func GetUFMap(lang string) map[string]string {
	if lang == "ptbr" {
		return address_ptbr.UFs
	}
	return nil
}

// registryList stores lists of translations by language and keyword
var (
	registryList = make(map[string]map[string][]string)
	mu           sync.RWMutex
)

// registry for single string translations
var (
	registrySingle = make(map[string]map[string]string)
)

// RegisterList records lists of translations
func RegisterList(lang string, messages map[string][]string) {
	if lang == "" || messages == nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if registryList[lang] == nil {
		registryList[lang] = make(map[string][]string)
	}
	for key, values := range messages {
		registryList[lang][key] = values
	}
}

// Register records single string translations
func Register(lang string, messages map[string]string) {
	if lang == "" || messages == nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if registrySingle[lang] == nil {
		registrySingle[lang] = make(map[string]string)
	}
	for key, value := range messages {
		registrySingle[lang][key] = value
	}
}

// Get returns a single string translation for a given language and key
func Get(lang, key string) string {
	if lang == "" || key == "" {
		return key
	}
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := registrySingle[lang][key]; ok {
		return val
	}
	return key
}

// GetList returns a list of translations for a given language and keyword
func GetList(lang, key string) []string {
	if lang == "" || key == "" {
		return nil
	}
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := registryList[lang][key]; ok {
		return val
	}
	return nil
}

// GetRandom returns a random value from a list of translations
func GetRandom(lang, key string, rnd RandSource) string {
	values := GetList(lang, key)
	if len(values) == 0 {
		return key
	}
	if rnd == nil {
		return values[0]
	}
	idx := rnd.Intn(len(values))
	return values[idx]
}

// RandSource it is an interface for randomness sources >> compatible with rand.rand.
type RandSource interface {
	Intn(n int) int
}
