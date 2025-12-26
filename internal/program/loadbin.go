package program

import (
	"fmt"
	"os"
)

func LoadBin256(path string) ([]uint8, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(b) != 256 {
		return nil, fmt.Errorf("expected 256 bytes, got %d", len(b))
	}
	return b, nil
}
