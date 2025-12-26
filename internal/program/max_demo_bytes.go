package program

import "emulator/internal/cpu"

const (
	ARR0 = 0x80
	ARR1 = 0x81
	ARR2 = 0x82
	ARR3 = 0x83
	ARR4 = 0x84
	ARR5 = 0x85
	ARR6 = 0x86
	ARR7 = 0x87

	MAX = 0x72
)

func LoadMax8Demo(c *cpu.CPU) {
	c.Mem[ARR0] = 5
	c.Mem[ARR1] = 120
	c.Mem[ARR2] = 7
	c.Mem[ARR3] = 42
	c.Mem[ARR4] = 250
	c.Mem[ARR5] = 9
	c.Mem[ARR6] = 13
	c.Mem[ARR7] = 1

	p := []uint8{
		cpu.OP_LOAD, ARR0,
		cpu.OP_STORE, MAX,

		cpu.OP_LOAD, ARR1,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x0C,
		cpu.OP_JMP, 0x10,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x10,

		cpu.OP_LOAD, ARR2,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x18,
		cpu.OP_JMP, 0x1C,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x1C,

		cpu.OP_LOAD, ARR3,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x24,
		cpu.OP_JMP, 0x28,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x28,

		cpu.OP_LOAD, ARR4,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x30,
		cpu.OP_JMP, 0x34,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x34,

		cpu.OP_LOAD, ARR5,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x3C,
		cpu.OP_JMP, 0x40,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x40,

		cpu.OP_LOAD, ARR6,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x48,
		cpu.OP_JMP, 0x4C,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x4C,

		cpu.OP_LOAD, ARR7,
		cpu.OP_CMP, MAX,
		cpu.OP_JG, 0x54,
		cpu.OP_JMP, 0x58,
		cpu.OP_STORE, MAX,
		cpu.OP_JMP, 0x58,

		cpu.OP_HLT, 0x00,
	}

	copy(c.Mem[:], p)
	c.PC = 0x00
}
