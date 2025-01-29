package translations

var registry = make(map[string]map[string]string) // Registry stores translations for each language
var currentLang = "pt"                            // Default language

// Register adds translations for a specific language.
func Register(lang string, messages map[string]string) {
	if registry[lang] == nil {
		registry[lang] = make(map[string]string)
	}
	for key, value := range messages {
		registry[lang][key] = value
	}
}

// SetLanguage sets the current language for translations.
func SetLanguage(lang string) {
	currentLang = lang
}

// GetLanguage returns the current language.
func GetLanguage() string {
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
