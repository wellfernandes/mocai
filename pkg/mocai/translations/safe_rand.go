package translations

import (
	"math/rand/v2"
	"sync"
)

// globalRand uses the thread-safe, auto-seeded global rand/v2 source.
type globalRand struct{}

func (g *globalRand) Intn(n int) int {
	return rand.IntN(n)
}

// DefaultRandSource returns a thread-safe, auto-seeded randomness source.
// In Go 1.22+, the global math/rand/v2 functions are thread-safe and auto-seeded.
func DefaultRandSource() RandSource {
	return &globalRand{}
}

// SafeRandSource wraps a RandSource with a mutex for concurrent safety.
// Use this when injecting a custom *rand.Rand that is not thread-safe.
type SafeRandSource struct {
	rnd RandSource
	mu  sync.Mutex
}

// NewSafeRandSource creates a new SafeRandSource wrapping the given RandSource.
func NewSafeRandSource(rnd RandSource) *SafeRandSource {
	return &SafeRandSource{rnd: rnd}
}

// Intn returns a non-negative pseudo-random int in [0,n) in a thread-safe manner.
func (s *SafeRandSource) Intn(n int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.rnd.Intn(n)
}
