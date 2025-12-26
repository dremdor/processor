package assembler

import "strings"

func MnemonicToOpcode(m string) (uint8, bool) {
	switch strings.ToUpper(m) {
	case "LOAD":
		return 0x01, true
	case "STORE":
		return 0x02, true
	case "ADD":
		return 0x03, true
	case "SUB":
		return 0x04, true
	case "CMP":
		return 0x05, true
	case "AND":
		return 0x06, true
	case "OR":
		return 0x07, true
	case "JMP":
		return 0x08, true
	case "JZ":
		return 0x09, true
	case "JNZ":
		return 0x0A, true
	case "JG":
		return 0x0B, true
	case "JL":
		return 0x0C, true
	case "LOADI":
		return 0x0D, true
	case "STOREI":
		return 0x0E, true
	case "HLT":
		return 0xFF, true
	default:
		return 0, false
	}
}

func NeedsOperand(op uint8) bool {
	return op != 0xFF
}
