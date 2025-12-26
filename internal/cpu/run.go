package cpu

import "fmt"

func (c *CPU) Run(maxSteps int) error {
	for i := 0; i < maxSteps; i++ {
		if c.Halted {
			return nil
		}
		c.Step()
	}
	return fmt.Errorf("maxSteps reached (%d) - possible infinite loop", maxSteps)
}
