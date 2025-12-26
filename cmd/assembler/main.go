package main

import (
	"flag"
	"fmt"
	"os"

	"emulator/internal/assembler"
)

func main() {
	in := flag.String("in", "", "input .asm file")
	out := flag.String("out", "out.bin", "output binary file")
	flag.Parse()

	if *in == "" {
		fmt.Println("usage: go run ./cmd/assembler -in program.asm -out program.bin")
		os.Exit(2)
	}

	b, err := os.ReadFile(*in)
	if err != nil {
		fmt.Println("read:", err)
		os.Exit(1)
	}

	mem, err := assembler.Assemble(string(b), assembler.Options{MemSize: 256})
	if err != nil {
		fmt.Println("assemble error:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(*out, mem, 0644); err != nil {
		fmt.Println("write:", err)
		os.Exit(1)
	}

	fmt.Println("ok:", *out, "(256 bytes)")
}
