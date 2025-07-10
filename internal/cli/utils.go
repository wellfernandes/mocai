package cli

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
