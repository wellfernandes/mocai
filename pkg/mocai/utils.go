package mocai

import (
	"math/rand"
	"time"
)

// GenerateRandomNumber generates a random number between min and max.
func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// FormatCPF formats a CPF number.
func FormatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}
