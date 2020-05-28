package op

import (
	"fmt"
)

// HandleOp handles opcodes
func HandleOp(buf []byte) int {
	opbytes := 1
	switch buf[0] {
	case 0x00:
		fmt.Printf("NOP\n")
	case 0x01:
		fmt.Printf("LXI\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x02:
		fmt.Printf("STAX\tB\n")
	case 0x03:
		fmt.Printf("INX\t\tB\n")
	case 0x04:
		fmt.Printf("INR\t\tB\n")
	case 0x05:
		fmt.Printf("DCR\t\tB\n")
	case 0x06:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x07:
		fmt.Printf("RLC\n")
	case 0x08:
		fmt.Printf("NOP\n")
	case 0x09:
		fmt.Printf("DAD\t\tB\n")
	case 0x0A:
		fmt.Printf("LDAX\tB\n")
	case 0x0B:
		fmt.Printf("DCX\t\tB\n")
	case 0x0C:
		fmt.Printf("INR\t\tC\n")
	case 0x0D:
		fmt.Printf("DCR\t\tC\n")
	case 0x0E:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x0F:
		fmt.Printf("RRC\n")
	case 0x10:
		fmt.Printf("NOP\n")
	case 0x11:
		fmt.Printf("LXI\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x12:
		fmt.Printf("STAX\tD\n")
	case 0x13:
		fmt.Printf("INX\t\tD\n")
	case 0x14:
		fmt.Printf("INR\t\tD\n")
	case 0x15:
		fmt.Printf("DCR\t\tD\n")
	case 0x16:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x17:
		fmt.Printf("RAL\n")
	case 0x18:
		fmt.Printf("NOP\n")
	case 0x19:
		fmt.Printf("DAD\t\tD\n")
	case 0x1A:
		fmt.Printf("LDAX\tD\n")
	case 0x1B:
		fmt.Printf("DCX\t\tD\n")
	case 0x1C:
		fmt.Printf("INR\t\tE\n")
	case 0x1D:
		fmt.Printf("DCR\t\tE\n")
	case 0x1E:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x1F:
		fmt.Printf("RAR\n")
	case 0x20:
		fmt.Printf("RIM\n")
	case 0x21:
		fmt.Printf("LXI\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x22:
		fmt.Printf("SHLD\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x23:
		fmt.Printf("INX\t\tH\n")
	case 0x24:
		fmt.Printf("INR\t\tH\n")
	case 0x25:
		fmt.Printf("DCR\t\tH\n")
	case 0x26:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x27:
		fmt.Printf("DAA\n")
	case 0x28:
		fmt.Printf("NOP\n")
	case 0x29:
		fmt.Printf("DAD\t\tH\n")
	case 0x2A:
		fmt.Printf("LHLD\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x2B:
		fmt.Printf("DCX\t\tH\n")
	case 0x2C:
		fmt.Printf("INR\t\tL\n")
	case 0x2D:
		fmt.Printf("DCR\t\tL\n")
	case 0x2E:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x2F:
		fmt.Printf("CMA\n")
	case 0x30:
		fmt.Printf("SIM\n")
	case 0x31:
		fmt.Printf("LXI\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x32:
		fmt.Printf("STA\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x33:
		fmt.Printf("INX\t\tSP\n")
	case 0x34:
		fmt.Printf("INR\t\tM\n")
	case 0x35:
		fmt.Printf("DCR\t\tM\n")
	case 0x36:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x37:
		fmt.Printf("STC\n")
	case 0x38:
		fmt.Printf("NOP\n")
	case 0x39:
		fmt.Printf("DAD\t\tSP\n")
	case 0x3A:
		fmt.Printf("LDA\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x3B:
		fmt.Printf("DCX\t\tSP\n")
	case 0x3C:
		fmt.Printf("INR\t\tA\n")
	case 0x3D:
		fmt.Printf("DCR\t\tA\n")
	case 0x3E:
		fmt.Printf("MVI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x3F:
		fmt.Printf("CMC\n")
	case 0x40:
		fmt.Printf("MOV\t\tB,B\n")
	case 0x41:
		fmt.Printf("MOV\t\tB,C\n")
	case 0x42:
		fmt.Printf("MOV\t\tB,D\n")
	case 0x43:
		fmt.Printf("MOV\t\tB,E\n")
	case 0x44:
		fmt.Printf("MOV\t\tB,H\n")
	case 0x45:
		fmt.Printf("MOV\t\tB,L\n")
	case 0x46:
		fmt.Printf("MOV\t\tB,M\n")
	case 0x47:
		fmt.Printf("MOV\t\tB,A\n")
	case 0x48:
		fmt.Printf("MOV\t\tC,B\n")
	case 0x49:
		fmt.Printf("MOV\t\tC,C\n")
	case 0x4A:
		fmt.Printf("MOV\t\tC,D\n")
	case 0x4B:
		fmt.Printf("MOV\t\tC,E\n")
	case 0x4C:
		fmt.Printf("MOV\t\tC,H\n")
	case 0x4D:
		fmt.Printf("MOV\t\tC,L\n")
	case 0x4E:
		fmt.Printf("MOV\t\tC,M\n")
	case 0x4F:
		fmt.Printf("MOV\t\tC,A\n")
	case 0x50:
		fmt.Printf("MOV\t\tD,B\n")
	case 0x51:
		fmt.Printf("MOV\t\tD,C\n")
	case 0x52:
		fmt.Printf("MOV\t\tD,D\n")
	case 0x53:
		fmt.Printf("MOV\t\tD,E\n")
	case 0x54:
		fmt.Printf("MOV\t\tD,H\n")
	case 0x55:
		fmt.Printf("MOV\t\tD,L\n")
	case 0x56:
		fmt.Printf("MOV\t\tD,M\n")
	case 0x57:
		fmt.Printf("MOV\t\tD,A\n")
	case 0x58:
		fmt.Printf("MOV\t\tE,B\n")
	case 0x59:
		fmt.Printf("MOV\t\tE,C\n")
	case 0x5A:
		fmt.Printf("MOV\t\tE,D\n")
	case 0x5B:
		fmt.Printf("MOV\t\tE,E\n")
	case 0x5C:
		fmt.Printf("MOV\t\tE,H\n")
	case 0x5D:
		fmt.Printf("MOV\t\tE,L\n")
	case 0x5E:
		fmt.Printf("MOV\t\tE,M\n")
	case 0x5F:
		fmt.Printf("MOV\t\tE,A\n")
	case 0x60:
		fmt.Printf("MOV\t\tH,B\n")
	case 0x61:
		fmt.Printf("MOV\t\tH,C\n")
	case 0x62:
		fmt.Printf("MOV\t\tH,D\n")
	case 0x63:
		fmt.Printf("MOV\t\tH,E\n")
	case 0x64:
		fmt.Printf("MOV\t\tH,H\n")
	case 0x65:
		fmt.Printf("MOV\t\tH,L\n")
	case 0x66:
		fmt.Printf("MOV\t\tH,M\n")
	case 0x67:
		fmt.Printf("MOV\t\tH,A\n")
	case 0x68:
		fmt.Printf("MOV\t\tL,B\n")
	case 0x69:
		fmt.Printf("MOV\t\tL,C\n")
	case 0x6A:
		fmt.Printf("MOV\t\tL,D\n")
	case 0x6B:
		fmt.Printf("MOV\t\tL,E\n")
	case 0x6C:
		fmt.Printf("MOV\t\tL,H\n")
	case 0x6D:
		fmt.Printf("MOV\t\tL,L\n")
	case 0x6E:
		fmt.Printf("MOV\t\tL,M\n")
	case 0x6F:
		fmt.Printf("MOV\t\tL,A\n")
	case 0x70:
		fmt.Printf("MOV\t\tM,B\n")
	case 0x71:
		fmt.Printf("MOV\t\tM,C\n")
	case 0x72:
		fmt.Printf("MOV\t\tM,D\n")
	case 0x73:
		fmt.Printf("MOV\t\tM,E\n")
	case 0x74:
		fmt.Printf("MOV\t\tM,H\n")
	case 0x75:
		fmt.Printf("MOV\t\tM,L\n")
	case 0x76:
		fmt.Printf("HLT\n")
	case 0x77:
		fmt.Printf("MOV\t\tM,A\n")
	case 0x78:
		fmt.Printf("MOV\t\tA,B\n")
	case 0x79:
		fmt.Printf("MOV\t\tA,C\n")
	case 0x7A:
		fmt.Printf("MOV\t\tA,D\n")
	case 0x7B:
		fmt.Printf("MOV\t\tA,E\n")
	case 0x7C:
		fmt.Printf("MOV\t\tA,H\n")
	case 0x7D:
		fmt.Printf("MOV\t\tA,L\n")
	case 0x7E:
		fmt.Printf("MOV\t\tA,M\n")
	case 0x7F:
		fmt.Printf("MOV\t\tA,A\n")
	case 0x80:
		fmt.Printf("ADD\t\tB\n")
	case 0x81:
		fmt.Printf("ADD\t\tC\n")
	case 0x82:
		fmt.Printf("ADD\t\tD\n")
	case 0x83:
		fmt.Printf("ADD\t\tE\n")
	case 0x84:
		fmt.Printf("ADD\t\tH\n")
	case 0x85:
		fmt.Printf("ADD\t\tL\n")
	case 0x86:
		fmt.Printf("ADD\t\tM\n")
	case 0x87:
		fmt.Printf("ADD\t\tA\n")
	case 0x88:
		fmt.Printf("ADC\t\tB\n")
	case 0x89:
		fmt.Printf("ADC\t\tC\n")
	case 0x8A:
		fmt.Printf("ADC\t\tD\n")
	case 0x8B:
		fmt.Printf("ADC\t\tE\n")
	case 0x8C:
		fmt.Printf("ADC\t\tH\n")
	case 0x8D:
		fmt.Printf("ADC\t\tL\n")
	case 0x8E:
		fmt.Printf("ADC\t\tM\n")
	case 0x8F:
		fmt.Printf("ADC\t\tA\n")
	case 0x90:
		fmt.Printf("SUB\t\tB\n")
	case 0x91:
		fmt.Printf("SUB\t\tC\n")
	case 0x92:
		fmt.Printf("SUB\t\tD\n")
	case 0x93:
		fmt.Printf("SUB\t\tE\n")
	case 0x94:
		fmt.Printf("SUB\t\tH\n")
	case 0x95:
		fmt.Printf("SUB\t\tL\n")
	case 0x96:
		fmt.Printf("SUB\t\tM\n")
	case 0x97:
		fmt.Printf("SUB\t\tA\n")
	case 0x98:
		fmt.Printf("SBB\t\tB\n")
	case 0x99:
		fmt.Printf("SBB\t\tC\n")
	case 0x9A:
		fmt.Printf("SBB\t\tD\n")
	case 0x9B:
		fmt.Printf("SBB\t\tE\n")
	case 0x9C:
		fmt.Printf("SBB\t\tH\n")
	case 0x9D:
		fmt.Printf("SBB\t\tL\n")
	case 0x9E:
		fmt.Printf("SBB\t\tM\n")
	case 0x9F:
		fmt.Printf("SBB\t\tA\n")
	case 0xA0:
		fmt.Printf("ANA\t\tB\n")
	case 0xA1:
		fmt.Printf("ANA\t\tC\n")
	case 0xA2:
		fmt.Printf("ANA\t\tD\n")
	case 0xA3:
		fmt.Printf("ANA\t\tE\n")
	case 0xA4:
		fmt.Printf("ANA\t\tH\n")
	case 0xA5:
		fmt.Printf("ANA\t\tL\n")
	case 0xA6:
		fmt.Printf("ANA\t\tM\n")
	case 0xA7:
		fmt.Printf("ANA\t\tA\n")
	case 0xA8:
		fmt.Printf("XRA\t\tB\n")
	case 0xA9:
		fmt.Printf("XRA\t\tC\n")
	case 0xAA:
		fmt.Printf("XRA\t\tD\n")
	case 0xAB:
		fmt.Printf("XRA\t\tE\n")
	case 0xAC:
		fmt.Printf("XRA\t\tH\n")
	case 0xAD:
		fmt.Printf("XRA\t\tL\n")
	case 0xAE:
		fmt.Printf("XRA\t\tM\n")
	case 0xAF:
		fmt.Printf("XRA\t\tA\n")
	case 0xB0:
		fmt.Printf("ORA\t\tB\n")
	case 0xB1:
		fmt.Printf("ORA\t\tC\n")
	case 0xB2:
		fmt.Printf("ORA\t\tD\n")
	case 0xB3:
		fmt.Printf("ORA\t\tE\n")
	case 0xB4:
		fmt.Printf("ORA\t\tH\n")
	case 0xB5:
		fmt.Printf("ORA\t\tL\n")
	case 0xB6:
		fmt.Printf("ORA\t\tM\n")
	case 0xB7:
		fmt.Printf("ORA\t\tA\n")
	case 0xB8:
		fmt.Printf("CMP\t\tB\n")
	case 0xB9:
		fmt.Printf("CMP\t\tC\n")
	case 0xBA:
		fmt.Printf("CMP\t\tD\n")
	case 0xBB:
		fmt.Printf("CMP\t\tE\n")
	case 0xBC:
		fmt.Printf("CMP\t\tH\n")
	case 0xBD:
		fmt.Printf("CMP\t\tL\n")
	case 0xBE:
		fmt.Printf("CMP\t\tM\n")
	case 0xBF:
		fmt.Printf("CMP\t\tA\n")
	case 0xC0:
		fmt.Printf("RNZ\n")
	case 0xC1:
		fmt.Printf("POP\t\tB\n")
	case 0xC2:
		fmt.Printf("JNZ\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC3:
		fmt.Printf("JMP\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC4:
		fmt.Printf("CNZ\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC5:
		fmt.Printf("PUSH\tB\n")
	case 0xC6:
		fmt.Printf("ADI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xC7:
		fmt.Printf("RST\t\t0\n")
	case 0xC8:
		fmt.Printf("RZ\n")
	case 0xC9:
		fmt.Printf("RET\n")
	case 0xCA:
		fmt.Printf("JZ\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCB:
		fmt.Printf("NOP\n")
	case 0xCC:
		fmt.Printf("CZ\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCD:
		fmt.Printf("CALL\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCE:
		fmt.Printf("ACI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xCF:
		fmt.Printf("RST\t\t1\n")
	case 0xD0:
		fmt.Printf("RNC\n")
	case 0xD1:
		fmt.Printf("POP\t\tD\n")
	case 0xD2:
		fmt.Printf("JNC\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xD3:
		fmt.Printf("OUT\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xD4:
		fmt.Printf("CNC\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xD5:
		fmt.Printf("PUSH\tD\n")
	case 0xD6:
		fmt.Printf("SUI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xD7:
		fmt.Printf("RST\t\t2\n")
	case 0xD8:
		fmt.Printf("RC\n")
	case 0xD9:
		fmt.Printf("NOP\n")
	case 0xDA:
		fmt.Printf("JC\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xDB:
		fmt.Printf("IN\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xDC:
		fmt.Printf("CC\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xDD:
		fmt.Printf("NOP\n")
	case 0xDE:
		fmt.Printf("SBI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xDF:
		fmt.Printf("RST\t\t3\n")
	case 0xE0:
		fmt.Printf("RPO\n")
	case 0xE1:
		fmt.Printf("POP\t\tH\n")
	case 0xE2:
		fmt.Printf("JPO\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xE3:
		fmt.Printf("XTHL\n")
	case 0xE4:
		fmt.Printf("CPO\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xE5:
		fmt.Printf("PUSH\tH\n")
	case 0xE6:
		fmt.Printf("ANI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xE7:
		fmt.Printf("RST\t\t4\n")
	case 0xE8:
		fmt.Printf("RPE\n")
	case 0xE9:
		fmt.Printf("PCHL\n")
	case 0xEA:
		fmt.Printf("JPE\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xEB:
		fmt.Printf("XCHG\n")
	case 0xEC:
		fmt.Printf("CPE\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xED:
		fmt.Printf("NOP\n")
	case 0xEE:
		fmt.Printf("XRI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xEF:
		fmt.Printf("RST\t\t5\n")
	case 0xF0:
		fmt.Printf("RP\n")
	case 0xF1:
		fmt.Printf("POP\t\tPSW\n")
	case 0xF2:
		fmt.Printf("JP\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xF3:
		fmt.Printf("DI\n")
	case 0xF4:
		fmt.Printf("CP\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xF5:
		fmt.Printf("PUSH\tPSW\n")
	case 0xF6:
		fmt.Printf("ORI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xF7:
		fmt.Printf("RST\t\t6\n")
	case 0xF8:
		fmt.Printf("RM\n")
	case 0xF9:
		fmt.Printf("SPHL\n")
	case 0xFA:
		fmt.Printf("JM\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xFB:
		fmt.Printf("EI\n")
	case 0xFC:
		fmt.Printf("CM\t\tB,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xFD:
		fmt.Printf("NOP\n")
	case 0xFE:
		fmt.Printf("CPI\t\tB,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xFF:
		fmt.Printf("RST\t\t7\n")
	default:
		fmt.Printf("Unknown OpCode: 0x%02X\n", buf[0])
	}

	return opbytes
}
