package cpu

func (c *CPU) fetch() (op, arg uint8) {
	pc := c.PC
	op = c.Mem[pc]
	arg = c.Mem[uint8(pc+1)]
	return
}

func (c *CPU) Step() {
	if c.Halted {
		return
	}

	op, arg := c.fetch()
	c.IRop, c.IRarg = op, arg

	nextPC := uint8(c.PC + 2)

	switch op {
	case OP_LOAD:
		c.ACC = int16(c.Mem[arg])

	case OP_STORE:
		c.Mem[arg] = uint8(c.ACC)

	case OP_ADD:
		c.ACC += int16(c.Mem[arg])

	case OP_SUB:
		c.ACC -= int16(c.Mem[arg])

	case OP_CMP:
		m := int16(c.Mem[arg])
		c.Flags.Z = (c.ACC == m)
		c.Flags.G = (c.ACC > m)
		c.Flags.L = (c.ACC < m)

	case OP_AND:
		c.ACC = int16(uint8(c.ACC) & c.Mem[arg])

	case OP_OR:
		c.ACC = int16(uint8(c.ACC) | c.Mem[arg])

	case OP_JMP:
		nextPC = arg

	case OP_JZ:
		if c.Flags.Z {
			nextPC = arg
		}

	case OP_JNZ:
		if !c.Flags.Z {
			nextPC = arg
		}

	case OP_JG:
		if c.Flags.G {
			nextPC = arg
		}

	case OP_JL:
		if c.Flags.L {
			nextPC = arg
		}

	case OP_LOADI:
		ptr := c.Mem[arg]
		c.ACC = int16(c.Mem[ptr])

	case OP_STOREI:
		ptr := c.Mem[arg]
		c.Mem[ptr] = uint8(c.ACC)

	case OP_HLT:
		c.Halted = true

	default:
		c.Halted = true
	}

	c.PC = nextPC
}
