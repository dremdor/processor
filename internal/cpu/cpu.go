package cpu

type Flags struct {
	Z bool
	G bool
	L bool
}

type CPU struct {
	Mem [256]uint8

	ACC int16
	PC  uint8

	IRop  uint8
	IRarg uint8

	Flags  Flags
	Halted bool
}

func (c *CPU) Reset() {
	c.ACC = 0
	c.PC = 0
	c.IRop = 0
	c.IRarg = 0
	c.Flags = Flags{}
	c.Halted = false
}
