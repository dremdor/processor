package cpu

import (
	"fmt"
	"strings"
)

func (c *CPU) DumpRegs() string {
	return fmt.Sprintf("PC=%02X ACC=%d IR=(%02X %02X) Z=%v G=%v L=%v HALT=%v",
		c.PC, c.ACC, c.IRop, c.IRarg, c.Flags.Z, c.Flags.G, c.Flags.L, c.Halted)
}

func (c *CPU) DumpMem(from, to uint8) string {
	var b strings.Builder
	for a := int(from); a <= int(to); a++ {
		addr := uint8(a)
		if (a-int(from))%16 == 0 {
			b.WriteString(fmt.Sprintf("\n0x%02X: ", addr))
		}
		b.WriteString(fmt.Sprintf("%02X ", c.Mem[addr]))
	}
	return b.String() + "\n"
}
