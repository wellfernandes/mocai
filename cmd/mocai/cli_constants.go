package mocai

// ANSI Escape Codes
var (
	CyanBold   string
	BlueBold   string
	Italic     string
	ResetStyle string
	LineBreak  = "\n"
)

// Cli Components
var (
	HeaderMain string
	SubHeader  string
	Footer     string
)

func init() {
	// Checks if the terminal supports ANSI
	if SupportsANSI() {
		CyanBold = "\033[1;36m"
		BlueBold = "\033[1;34m"
		Italic = "\033[3m"
		ResetStyle = "\033[0m"
	} else {
		CyanBold = ""
		BlueBold = ""
		Italic = ""
		ResetStyle = ""
	}

	// Defines the cli components
	HeaderMain = CyanBold + "»»————————  " + BlueBold + "M O C A Í" + ResetStyle + "  " + CyanBold + "————————««" + ResetStyle
	SubHeader = Italic + "✨ Generating mocks, simplifying tests, accelerating development." + ResetStyle + LineBreak +
		Italic + "🐧 Open Source | Mock Data Generator" + ResetStyle + LineBreak
	Footer = CyanBold + "\n»»—————————————————————————————««" + ResetStyle
}
