package translations

import (
	"fmt"
	"sync"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
)

var (
	registry    = make(map[string]map[string]string) // Registry stores translations for each language
	currentLang = "pt"                               // Default language
	mu          sync.RWMutex                         // Mutex to protect the registry
)

// Register adds translations for a specific language.
func Register(lang string, messages map[string]string) {
	if lang == "" {
		return
	}
	if messages == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if registry[lang] == nil {
		registry[lang] = make(map[string]string)
	}
	for key, value := range messages {
		registry[lang][key] = value
	}
}

// SetLanguage sets the current language for translations.
func SetLanguage(lang string) error {
	mu.RLock()
	_, exists := registry[lang]
	mu.RUnlock()
	if !exists {
		return fmt.Errorf(constants.ERROR_UNSUPPORTED_LANGUAGE+" %s", lang)
	}
	mu.Lock()
	currentLang = lang
	mu.Unlock()
	return nil
}

// GetLanguage returns the current language.
func GetLanguage() string {
	mu.RLock()
	defer mu.RUnlock()
	return currentLang
}

// Get retrieves a translation for a specific key in a given language.
func Get(lang, key string) string {
	if val, ok := registry[lang][key]; ok {
		return val
	}
	return key // Return the key itself if the translation is not found
}

// Translate retrieves a translation for a specific key in the current language.
func Translate(key string) string {
	return Get(currentLang, key)
}
