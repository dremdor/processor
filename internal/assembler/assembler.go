package assembler

import (
	"fmt"
)

type Options struct {
	MemSize int
}

func Assemble(src string, opt Options) ([]uint8, error) {
	if opt.MemSize <= 0 {
		opt.MemSize = 256
	}

	lines, err := LexLines(src)
	if err != nil {
		return nil, err
	}

	sym := NewSymbols()
	pc := 0

	for _, ln := range lines {
		if ln.Label != "" {
			if err := sym.Define(ln.Label, pc); err != nil {
				return nil, err
			}
		}

		switch ln.Kind {
		case LineEmpty:
		case LineOrg:
			v, err := ParseValue(ln.Arg0, sym, false)
			if err != nil {
				return nil, fmt.Errorf("%d: .org: %w", ln.LineNo, err)
			}
			pc = int(v)
			if pc < 0 || pc >= opt.MemSize {
				return nil, fmt.Errorf("%d: .org out of memory range", ln.LineNo)
			}
		case LineDb:
			pc += len(ln.Args)
			if pc > opt.MemSize {
				return nil, fmt.Errorf("%d: .db writes past memory size", ln.LineNo)
			}
		case LineInstr:
			pc += 2
			if pc > opt.MemSize {
				return nil, fmt.Errorf("%d: program writes past memory size", ln.LineNo)
			}
		default:
			return nil, fmt.Errorf("%d: unknown line kind", ln.LineNo)
		}
	}

	mem := make([]uint8, opt.MemSize)
	pc = 0

	for _, ln := range lines {
		switch ln.Kind {
		case LineEmpty:

		case LineOrg:
			v, err := ParseValue(ln.Arg0, sym, true)
			if err != nil {
				return nil, fmt.Errorf("%d: .org: %w", ln.LineNo, err)
			}
			pc = int(v)

		case LineDb:
			for i, a := range ln.Args {
				v, err := ParseValue(a, sym, true)
				if err != nil {
					return nil, fmt.Errorf("%d: .db: %w", ln.LineNo, err)
				}
				mem[pc+i] = uint8(v)
			}
			pc += len(ln.Args)

		case LineInstr:
			op, ok := MnemonicToOpcode(ln.Mnemonic)
			if !ok {
				return nil, fmt.Errorf("%d: unknown mnemonic %q", ln.LineNo, ln.Mnemonic)
			}

			operand := uint8(0)
			if NeedsOperand(op) {
				v, err := ParseValue(ln.Arg0, sym, true)
				if err != nil {
					return nil, fmt.Errorf("%d: operand: %w", ln.LineNo, err)
				}
				operand = uint8(v)
			}

			mem[pc] = op
			mem[pc+1] = operand
			pc += 2
		}
	}

	return mem, nil
}
