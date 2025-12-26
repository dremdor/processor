package assembler

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseValue(tok string, sym *Symbols, allowSymbols bool) (uint16, error) {
	tok = strings.TrimSpace(tok)
	if tok == "" {
		return 0, fmt.Errorf("empty value")
	}

	if strings.HasPrefix(tok, "0x") || strings.HasPrefix(tok, "0X") {
		v, err := strconv.ParseUint(tok[2:], 16, 16)
		if err != nil {
			return 0, fmt.Errorf("bad hex number %q", tok)
		}
		if v > 0xFFFF {
			return 0, fmt.Errorf("number too large %q", tok)
		}
		return uint16(v), nil
	}

	if isDecimal(tok) {
		v, err := strconv.ParseUint(tok, 10, 16)
		if err != nil {
			return 0, fmt.Errorf("bad decimal number %q", tok)
		}
		return uint16(v), nil
	}

	if sym == nil {
		return 0, fmt.Errorf("unknown symbol %q", tok)
	}
	if addr, ok := sym.Lookup(tok); ok {
		if addr < 0 || addr > 0xFFFF {
			return 0, fmt.Errorf("symbol out of range %q", tok)
		}
		return uint16(addr), nil
	}

	if allowSymbols {
		return 0, fmt.Errorf("undefined symbol %q", tok)
	}
	return 0, nil
}

func isDecimal(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
