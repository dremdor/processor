build:
	go build -o assembler ./cmd/assembler
	go build -o emulator ./cmd/emulator

asm-max8:
	go run ./cmd/assembler -in programs/max8.asm -out programs/max8.bin

asm-maxN:
	go run ./cmd/assembler -in programs/maxN.asm -out programs/maxN.bin

run-max8:
	go run ./cmd/emulator -bin programs/max8.bin

run-maxN:
	go run ./cmd/emulator -bin programs/maxN.bin

trace-max8:
	go run ./cmd/emulator -bin programs/max8.bin -trace

trace-maxN:
	go run ./cmd/emulator -bin programs/maxN.bin -trace

clean:
	rm -f assembler emulator
	rm -f programs/max8.bin programs/maxN.bin
