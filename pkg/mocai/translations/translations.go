package translations

import (
	"sync"

	address_ptbr "github.com/brazzcore/mocai/pkg/mocai/entities/address/mocks/ptbr"
)

// GetUFMap returns the mapping of states for the specified language, if available.
func GetUFMap(lang string) map[string]string {
	if lang == "ptbr" {
		// returns a defensive copy to avoid exposing the global map
		copyMap := make(map[string]string, len(address_ptbr.UFs))
		for k, v := range address_ptbr.UFs {
			copyMap[k] = v
		}
		return copyMap
	}
	return nil
}

// registryList stores lists of translations by language and keyword.
// registrySingle stores single string translations by language and keyword.
// Both are protected by mu for concurrent access safety.
var (
	registryList   = make(map[string]map[string][]string)
	registrySingle = make(map[string]map[string]string)
	mu             sync.RWMutex
)

// RegisterList records lists of translations.
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

// Register records single string translations.
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

// Get returns a single string translation for a given language and key.
func Get(lang, key string) string {
	if lang == "" || key == "" {
		return key
	}
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := registrySingle[lang][key]; ok {
		return val
	}

	// fallback: try en_us if not found in requested language
	if lang != "en_us" {
		if val, ok := registrySingle["en_us"][key]; ok {
			return val
		}
	}
	return key
}

// GetList returns a list of translations for a given language and keyword.
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

// GetRandom returns a random value from a list of translations.
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

// RandSource is an interface for randomness sources, compatible with math/rand.Rand.
type RandSource interface {
	Intn(n int) int
}
