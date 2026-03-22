package translations

import (
	"math/rand"
	"sync"
	"time"
)

// DefaultRandSource returns a standard thread-safe randomness source for global use
func DefaultRandSource() RandSource {
	return NewSafeRandSource(rand.New(rand.NewSource(time.Now().UnixNano())))
}

// SafeRandSource wraps a RandSource with a mutex for concurrent safety
type SafeRandSource struct {
	rnd RandSource
	mu  sync.Mutex
}

func NewSafeRandSource(rnd RandSource) *SafeRandSource {
	return &SafeRandSource{rnd: rnd}
}

func (s *SafeRandSource) Intn(n int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.rnd.Intn(n)
}
