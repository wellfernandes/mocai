package adhar

import (
	"math/rand"
	"strconv"
	"time"
)

// Generate a random Aadhaar number
func GenerateAadhaarNumber() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	aadhaar := ""

	for i := 0; i < 12; i++ {
		aadhaar += strconv.Itoa(rng.Intn(10))
	}
	return aadhaar
}
