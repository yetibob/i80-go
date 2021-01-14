package vm

import (
	"fmt"
)

func handleOp(buf []byte) int {
	opbytes := 1
	switch buf[0] {
	case 0x00:
		fmt.Printf("NOP\n")
	case 0x01:
		fmt.Printf("LXI     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x02:
		fmt.Printf("STAX    B\n")
	case 0x03:
		fmt.Printf("INX     B\n")
	case 0x04:
		fmt.Printf("INR     B\n")
	case 0x05:
		fmt.Printf("DCR     B\n")
	case 0x06:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x07:
		fmt.Printf("RLC\n")
	case 0x08:
		fmt.Printf("NOP\n")
	case 0x09:
		fmt.Printf("DAD     B\n")
	case 0x0A:
		fmt.Printf("LDAX    B\n")
	case 0x0B:
		fmt.Printf("DCX     B\n")
	case 0x0C:
		fmt.Printf("INR     C\n")
	case 0x0D:
		fmt.Printf("DCR     C\n")
	case 0x0E:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x0F:
		fmt.Printf("RRC\n")
	case 0x10:
		fmt.Printf("NOP\n")
	case 0x11:
		fmt.Printf("LXI     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x12:
		fmt.Printf("STAX    D\n")
	case 0x13:
		fmt.Printf("INX     D\n")
	case 0x14:
		fmt.Printf("INR     D\n")
	case 0x15:
		fmt.Printf("DCR     D\n")
	case 0x16:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x17:
		fmt.Printf("RAL\n")
	case 0x18:
		fmt.Printf("NOP\n")
	case 0x19:
		fmt.Printf("DAD     D\n")
	case 0x1A:
		fmt.Printf("LDAX    D\n")
	case 0x1B:
		fmt.Printf("DCX     D\n")
	case 0x1C:
		fmt.Printf("INR     E\n")
	case 0x1D:
		fmt.Printf("DCR     E\n")
	case 0x1E:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x1F:
		fmt.Printf("RAR\n")
	case 0x20:
		fmt.Printf("RIM\n")
	case 0x21:
		fmt.Printf("LXI     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x22:
		fmt.Printf("SHLD    B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x23:
		fmt.Printf("INX     H\n")
	case 0x24:
		fmt.Printf("INR     H\n")
	case 0x25:
		fmt.Printf("DCR     H\n")
	case 0x26:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x27:
		fmt.Printf("DAA\n")
	case 0x28:
		fmt.Printf("NOP\n")
	case 0x29:
		fmt.Printf("DAD     H\n")
	case 0x2A:
		fmt.Printf("LHLD    B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x2B:
		fmt.Printf("DCX     H\n")
	case 0x2C:
		fmt.Printf("INR     L\n")
	case 0x2D:
		fmt.Printf("DCR     L\n")
	case 0x2E:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x2F:
		fmt.Printf("CMA\n")
	case 0x30:
		fmt.Printf("SIM\n")
	case 0x31:
		fmt.Printf("LXI     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x32:
		fmt.Printf("STA     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x33:
		fmt.Printf("INX     SP\n")
	case 0x34:
		fmt.Printf("INR     M\n")
	case 0x35:
		fmt.Printf("DCR     M\n")
	case 0x36:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x37:
		fmt.Printf("STC\n")
	case 0x38:
		fmt.Printf("NOP\n")
	case 0x39:
		fmt.Printf("DAD     SP\n")
	case 0x3A:
		fmt.Printf("LDA     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0x3B:
		fmt.Printf("DCX     SP\n")
	case 0x3C:
		fmt.Printf("INR     A\n")
	case 0x3D:
		fmt.Printf("DCR     A\n")
	case 0x3E:
		fmt.Printf("MVI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0x3F:
		fmt.Printf("CMC\n")
	case 0x40:
		fmt.Printf("MOV     B,B\n")
	case 0x41:
		fmt.Printf("MOV     B,C\n")
	case 0x42:
		fmt.Printf("MOV     B,D\n")
	case 0x43:
		fmt.Printf("MOV     B,E\n")
	case 0x44:
		fmt.Printf("MOV     B,H\n")
	case 0x45:
		fmt.Printf("MOV     B,L\n")
	case 0x46:
		fmt.Printf("MOV     B,M\n")
	case 0x47:
		fmt.Printf("MOV     B,A\n")
	case 0x48:
		fmt.Printf("MOV     C,B\n")
	case 0x49:
		fmt.Printf("MOV     C,C\n")
	case 0x4A:
		fmt.Printf("MOV     C,D\n")
	case 0x4B:
		fmt.Printf("MOV     C,E\n")
	case 0x4C:
		fmt.Printf("MOV     C,H\n")
	case 0x4D:
		fmt.Printf("MOV     C,L\n")
	case 0x4E:
		fmt.Printf("MOV     C,M\n")
	case 0x4F:
		fmt.Printf("MOV     C,A\n")
	case 0x50:
		fmt.Printf("MOV     D,B\n")
	case 0x51:
		fmt.Printf("MOV     D,C\n")
	case 0x52:
		fmt.Printf("MOV     D,D\n")
	case 0x53:
		fmt.Printf("MOV     D,E\n")
	case 0x54:
		fmt.Printf("MOV     D,H\n")
	case 0x55:
		fmt.Printf("MOV     D,L\n")
	case 0x56:
		fmt.Printf("MOV     D,M\n")
	case 0x57:
		fmt.Printf("MOV     D,A\n")
	case 0x58:
		fmt.Printf("MOV     E,B\n")
	case 0x59:
		fmt.Printf("MOV     E,C\n")
	case 0x5A:
		fmt.Printf("MOV     E,D\n")
	case 0x5B:
		fmt.Printf("MOV     E,E\n")
	case 0x5C:
		fmt.Printf("MOV     E,H\n")
	case 0x5D:
		fmt.Printf("MOV     E,L\n")
	case 0x5E:
		fmt.Printf("MOV     E,M\n")
	case 0x5F:
		fmt.Printf("MOV     E,A\n")
	case 0x60:
		fmt.Printf("MOV     H,B\n")
	case 0x61:
		fmt.Printf("MOV     H,C\n")
	case 0x62:
		fmt.Printf("MOV     H,D\n")
	case 0x63:
		fmt.Printf("MOV     H,E\n")
	case 0x64:
		fmt.Printf("MOV     H,H\n")
	case 0x65:
		fmt.Printf("MOV     H,L\n")
	case 0x66:
		fmt.Printf("MOV     H,M\n")
	case 0x67:
		fmt.Printf("MOV     H,A\n")
	case 0x68:
		fmt.Printf("MOV     L,B\n")
	case 0x69:
		fmt.Printf("MOV     L,C\n")
	case 0x6A:
		fmt.Printf("MOV     L,D\n")
	case 0x6B:
		fmt.Printf("MOV     L,E\n")
	case 0x6C:
		fmt.Printf("MOV     L,H\n")
	case 0x6D:
		fmt.Printf("MOV     L,L\n")
	case 0x6E:
		fmt.Printf("MOV     L,M\n")
	case 0x6F:
		fmt.Printf("MOV     L,A\n")
	case 0x70:
		fmt.Printf("MOV     M,B\n")
	case 0x71:
		fmt.Printf("MOV     M,C\n")
	case 0x72:
		fmt.Printf("MOV     M,D\n")
	case 0x73:
		fmt.Printf("MOV     M,E\n")
	case 0x74:
		fmt.Printf("MOV     M,H\n")
	case 0x75:
		fmt.Printf("MOV     M,L\n")
	case 0x76:
		fmt.Printf("HLT\n")
	case 0x77:
		fmt.Printf("MOV     M,A\n")
	case 0x78:
		fmt.Printf("MOV     A,B\n")
	case 0x79:
		fmt.Printf("MOV     A,C\n")
	case 0x7A:
		fmt.Printf("MOV     A,D\n")
	case 0x7B:
		fmt.Printf("MOV     A,E\n")
	case 0x7C:
		fmt.Printf("MOV     A,H\n")
	case 0x7D:
		fmt.Printf("MOV     A,L\n")
	case 0x7E:
		fmt.Printf("MOV     A,M\n")
	case 0x7F:
		fmt.Printf("MOV     A,A\n")
	case 0x80:
		fmt.Printf("ADD     B\n")
	case 0x81:
		fmt.Printf("ADD     C\n")
	case 0x82:
		fmt.Printf("ADD     D\n")
	case 0x83:
		fmt.Printf("ADD     E\n")
	case 0x84:
		fmt.Printf("ADD     H\n")
	case 0x85:
		fmt.Printf("ADD     L\n")
	case 0x86:
		fmt.Printf("ADD     M\n")
	case 0x87:
		fmt.Printf("ADD     A\n")
	case 0x88:
		fmt.Printf("ADC     B\n")
	case 0x89:
		fmt.Printf("ADC     C\n")
	case 0x8A:
		fmt.Printf("ADC     D\n")
	case 0x8B:
		fmt.Printf("ADC     E\n")
	case 0x8C:
		fmt.Printf("ADC     H\n")
	case 0x8D:
		fmt.Printf("ADC     L\n")
	case 0x8E:
		fmt.Printf("ADC     M\n")
	case 0x8F:
		fmt.Printf("ADC     A\n")
	case 0x90:
		fmt.Printf("SUB     B\n")
	case 0x91:
		fmt.Printf("SUB     C\n")
	case 0x92:
		fmt.Printf("SUB     D\n")
	case 0x93:
		fmt.Printf("SUB     E\n")
	case 0x94:
		fmt.Printf("SUB     H\n")
	case 0x95:
		fmt.Printf("SUB     L\n")
	case 0x96:
		fmt.Printf("SUB     M\n")
	case 0x97:
		fmt.Printf("SUB     A\n")
	case 0x98:
		fmt.Printf("SBB     B\n")
	case 0x99:
		fmt.Printf("SBB     C\n")
	case 0x9A:
		fmt.Printf("SBB     D\n")
	case 0x9B:
		fmt.Printf("SBB     E\n")
	case 0x9C:
		fmt.Printf("SBB     H\n")
	case 0x9D:
		fmt.Printf("SBB     L\n")
	case 0x9E:
		fmt.Printf("SBB     M\n")
	case 0x9F:
		fmt.Printf("SBB     A\n")
	case 0xA0:
		fmt.Printf("ANA     B\n")
	case 0xA1:
		fmt.Printf("ANA     C\n")
	case 0xA2:
		fmt.Printf("ANA     D\n")
	case 0xA3:
		fmt.Printf("ANA     E\n")
	case 0xA4:
		fmt.Printf("ANA     H\n")
	case 0xA5:
		fmt.Printf("ANA     L\n")
	case 0xA6:
		fmt.Printf("ANA     M\n")
	case 0xA7:
		fmt.Printf("ANA     A\n")
	case 0xA8:
		fmt.Printf("XRA     B\n")
	case 0xA9:
		fmt.Printf("XRA     C\n")
	case 0xAA:
		fmt.Printf("XRA     D\n")
	case 0xAB:
		fmt.Printf("XRA     E\n")
	case 0xAC:
		fmt.Printf("XRA     H\n")
	case 0xAD:
		fmt.Printf("XRA     L\n")
	case 0xAE:
		fmt.Printf("XRA     M\n")
	case 0xAF:
		fmt.Printf("XRA     A\n")
	case 0xB0:
		fmt.Printf("ORA     B\n")
	case 0xB1:
		fmt.Printf("ORA     C\n")
	case 0xB2:
		fmt.Printf("ORA     D\n")
	case 0xB3:
		fmt.Printf("ORA     E\n")
	case 0xB4:
		fmt.Printf("ORA     H\n")
	case 0xB5:
		fmt.Printf("ORA     L\n")
	case 0xB6:
		fmt.Printf("ORA     M\n")
	case 0xB7:
		fmt.Printf("ORA     A\n")
	case 0xB8:
		fmt.Printf("CMP     B\n")
	case 0xB9:
		fmt.Printf("CMP     C\n")
	case 0xBA:
		fmt.Printf("CMP     D\n")
	case 0xBB:
		fmt.Printf("CMP     E\n")
	case 0xBC:
		fmt.Printf("CMP     H\n")
	case 0xBD:
		fmt.Printf("CMP     L\n")
	case 0xBE:
		fmt.Printf("CMP     M\n")
	case 0xBF:
		fmt.Printf("CMP     A\n")
	case 0xC0:
		fmt.Printf("RNZ\n")
	case 0xC1:
		fmt.Printf("POP     B\n")
	case 0xC2:
		fmt.Printf("JNZ     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC3:
		fmt.Printf("JMP     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC4:
		fmt.Printf("CNZ     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xC5:
		fmt.Printf("PUSH    B\n")
	case 0xC6:
		fmt.Printf("ADI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xC7:
		fmt.Printf("RST     0\n")
	case 0xC8:
		fmt.Printf("RZ\n")
	case 0xC9:
		fmt.Printf("RET\n")
	case 0xCA:
		fmt.Printf("JZ      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCB:
		fmt.Printf("NOP\n")
	case 0xCC:
		fmt.Printf("CZ      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCD:
		fmt.Printf("CALL    B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xCE:
		fmt.Printf("ACI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xCF:
		fmt.Printf("RST     1\n")
	case 0xD0:
		fmt.Printf("RNC\n")
	case 0xD1:
		fmt.Printf("POP     D\n")
	case 0xD2:
		fmt.Printf("JNC     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xD3:
		fmt.Printf("OUT     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xD4:
		fmt.Printf("CNC     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xD5:
		fmt.Printf("PUSH    D\n")
	case 0xD6:
		fmt.Printf("SUI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xD7:
		fmt.Printf("RST     2\n")
	case 0xD8:
		fmt.Printf("RC\n")
	case 0xD9:
		fmt.Printf("NOP\n")
	case 0xDA:
		fmt.Printf("JC      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xDB:
		fmt.Printf("IN      B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xDC:
		fmt.Printf("CC      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xDD:
		fmt.Printf("NOP\n")
	case 0xDE:
		fmt.Printf("SBI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xDF:
		fmt.Printf("RST     3\n")
	case 0xE0:
		fmt.Printf("RPO\n")
	case 0xE1:
		fmt.Printf("POP     H\n")
	case 0xE2:
		fmt.Printf("JPO     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xE3:
		fmt.Printf("XTHL\n")
	case 0xE4:
		fmt.Printf("CPO     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xE5:
		fmt.Printf("PUSH    H\n")
	case 0xE6:
		fmt.Printf("ANI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xE7:
		fmt.Printf("RST     4\n")
	case 0xE8:
		fmt.Printf("RPE\n")
	case 0xE9:
		fmt.Printf("PCHL\n")
	case 0xEA:
		fmt.Printf("JPE     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xEB:
		fmt.Printf("XCHG\n")
	case 0xEC:
		fmt.Printf("CPE     B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xED:
		fmt.Printf("NOP\n")
	case 0xEE:
		fmt.Printf("XRI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xEF:
		fmt.Printf("RST     5\n")
	case 0xF0:
		fmt.Printf("RP\n")
	case 0xF1:
		fmt.Printf("POP     PSW\n")
	case 0xF2:
		fmt.Printf("JP      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xF3:
		fmt.Printf("DI\n")
	case 0xF4:
		fmt.Printf("CP      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xF5:
		fmt.Printf("PUSH    PSW\n")
	case 0xF6:
		fmt.Printf("ORI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xF7:
		fmt.Printf("RST     6\n")
	case 0xF8:
		fmt.Printf("RM\n")
	case 0xF9:
		fmt.Printf("SPHL\n")
	case 0xFA:
		fmt.Printf("JM      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xFB:
		fmt.Printf("EI\n")
	case 0xFC:
		fmt.Printf("CM      B,D16\t$%02X%02X\n", buf[2], buf[1])
		opbytes = 3
	case 0xFD:
		fmt.Printf("NOP\n")
	case 0xFE:
		fmt.Printf("CPI     B,D8\t$%02X\n", buf[1])
		opbytes = 2
	case 0xFF:
		fmt.Printf("RST     7\n")
	default:
		fmt.Printf("Unknown OpCode: 0x%02X\n", buf[0])
	}

	return opbytes
}
