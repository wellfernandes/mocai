package translations

import "sync"

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
