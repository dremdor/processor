package assembler

import "fmt"

type Symbols struct {
	m map[string]int
}

func NewSymbols() *Symbols {
	return &Symbols{m: make(map[string]int)}
}

func (s *Symbols) Define(name string, addr int) error {
	if _, exists := s.m[name]; exists {
		return fmt.Errorf("label redefined: %s", name)
	}
	s.m[name] = addr
	return nil
}

func (s *Symbols) Lookup(name string) (int, bool) {
	v, ok := s.m[name]
	return v, ok
}
