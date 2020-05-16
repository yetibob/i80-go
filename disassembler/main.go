package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read rom into buffer
	romFile := "./roms/invaders.h"
	rom, err := ioutil.ReadFile(romFile)
	if err != nil {
		panic(err)
	}

	pc := 0

	for {
		opbytes := 1
		if pc == len(rom) {
			break
		}
		switch code := &rom[pc]; *code {
		case 0x00:
			fmt.Printf("NOP\n")
		case 0x01:
			fmt.Printf("LXI    B,$%02X%02X\n", rom[pc+2], rom[pc+1])
			opbytes = 3
		case 0x02:
			fmt.Printf("STAX   B\n")
		case 0x03:
			fmt.Printf("INX    B\n")
		case 0x04:
			fmt.Printf("INR    B\n")
		case 0x05:
			fmt.Printf("DCR    B\n")
		case 0x06:
			fmt.Printf("MVI    B,$%02X%02X\n", rom[pc+2], rom[pc+1])
			opbytes = 3
		case 0x07:
			fmt.Printf("RLC\n")
		case 0x08:
			fmt.Printf("NOP\n")
		case 0x3e:
			fmt.Printf("MVI    A,$%02X\n", rom[pc+1])
			opbytes = 2
		case 0xc3:
			fmt.Printf("JMP    $%02X%02X\n", rom[pc+2], rom[pc+1])
			opbytes = 3
		}
		pc += opbytes
	}
	// use byte at current position in buffer to determine opcode

	// print out the name of the opcode using the bytes of the opcode as data, if applicable

	// advance the current position the number of bytes required by the instruction

	// if not at end of buffer, return to step 2
}
