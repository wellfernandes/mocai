package mocai

import (
	"os"
	"runtime"
)

// SupportsANSI checks if the terminal supports ANSI escape codes.
func SupportsANSI() bool {
	// Check if the output is a terminal
	fileInfo, err := os.Stdout.Stat()
	if err != nil || (fileInfo.Mode()&os.ModeCharDevice) == 0 {
		return false // Not a terminal
	}

	// Check for Windows (cmd.exe or PowerShell without ANSI support)
	if runtime.GOOS == "windows" && os.Getenv("ANSICON") == "" && os.Getenv("ConEmuANSI") == "" {
		return false
	}

	return true
}

// FormatCPF formats a CPF number.
func FormatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}
