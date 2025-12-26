package main

import (
	"flag"
	"fmt"

	"emulator/internal/cpu"
	"emulator/internal/program"
)

func main() {
	binPath := flag.String("bin", "", "path to 256-byte binary program")
	trace := flag.Bool("trace", false, "trace execution (regs + mem dump after each step)")
	maxSteps := flag.Int("maxSteps", 100000, "maximum steps before stop")
	flag.Parse()

	var c cpu.CPU
	c.Reset()

	fmt.Print(c.DumpMem(0x00, 0xFF))

	if *binPath == "" {
		program.LoadMax8Demo(&c)
	} else {
		mem, err := program.LoadBin256(*binPath)
		if err != nil {
			panic(err)
		}
		copy(c.Mem[:], mem)
		c.PC = 0x00
	}

	if *trace {
		fmt.Println("START:", c.DumpRegs())
		fmt.Print(c.DumpMem(0x70, 0x8F))
	}

	steps := 0
	for steps < *maxSteps && !c.Halted {
		c.Step()
		steps++

		if *trace {
			fmt.Printf("STEP %d: %s\n", steps, c.DumpRegs())
			fmt.Print(c.DumpMem(0x70, 0x8F))
		}
	}

	fmt.Println("Stopped.")
	fmt.Println("Steps:", steps)
	fmt.Println("Final:", c.DumpRegs())
	fmt.Printf("MAX (mem[0x72]) = %d\n", c.Mem[0x72])

	fmt.Print(c.DumpMem(0x00, 0xFF))
	if steps >= *maxSteps && !c.Halted {
		fmt.Printf("Warning: maxSteps reached (%d). Possible infinite loop.\n", *maxSteps)
	}
}
