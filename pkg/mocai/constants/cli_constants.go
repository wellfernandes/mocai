package constants

// ANSI Escape Codes
const (
	CyanBold   = "\033[1;36m"
	BlueBold   = "\033[1;34m"
	Italic     = "\033[3m"
	ResetStyle = "\033[0m"
	LineBreak  = "\n"
)

// Cli Components
const (
	HeaderMain = CyanBold + "»»————————  " + BlueBold + "M O C A Í" + ResetStyle + "  " + CyanBold + "————————««" + ResetStyle
	SubHeader  = Italic + "✨ Generating mocks, simplifying tests, accelerating development." + ResetStyle + LineBreak +
		Italic + "🐧 Open Source | Mock Data Generator" + ResetStyle + LineBreak

	Footer = CyanBold + "\n»»—————————————————————————————««" + ResetStyle
)
