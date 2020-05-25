package op

import (
	"fmt"
)

// HandleOp handles opcodes
func HandleOp(op byte, buf []byte) int {
	opbytes := 1
	switch op {
	case 0x00:
		fmt.Printf("NOP\n")
	case 0x01:
		fmt.Printf("LXI    B,$%02X%02X\n", buf[1], buf[0])
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
		fmt.Printf("MVI    B,$%02X\n", buf[0])
		opbytes = 2
	case 0x07:
		fmt.Printf("RLC\n")
	case 0x08:
		fmt.Printf("NOP\n")
	case 0x09:
		fmt.Println("Handling OpCode: 0x09")
	case 0x0A:
		fmt.Println("Handling OpCode: 0x0A")
	case 0x0B:
		fmt.Println("Handling OpCode: 0x0B")
	case 0x0C:
		fmt.Println("Handling OpCode: 0x0C")
	case 0x0D:
		fmt.Println("Handling OpCode: 0x0D")
	case 0x0E:
		fmt.Println("Handling OpCode: 0x0E")
		opbytes = 2
	case 0x0F:
		fmt.Println("Handling OpCode: 0x0F")
	case 0x10:
		fmt.Println("Handling OpCode: 0x10")
	case 0x11:
		fmt.Println("Handling OpCode: 0x11")
		opbytes = 3
	case 0x12:
		fmt.Println("Handling OpCode: 0x12")
	case 0x13:
		fmt.Println("Handling OpCode: 0x13")
	case 0x14:
		fmt.Println("Handling OpCode: 0x14")
	case 0x15:
		fmt.Println("Handling OpCode: 0x15")
	case 0x16:
		fmt.Println("Handling OpCode: 0x16")
		opbytes = 2
	case 0x17:
		fmt.Println("Handling OpCode: 0x17")
	case 0x18:
		fmt.Println("Handling OpCode: 0x18")
	case 0x19:
		fmt.Println("Handling OpCode: 0x19")
	case 0x1A:
		fmt.Println("Handling OpCode: 0x1A")
	case 0x1B:
		fmt.Println("Handling OpCode: 0x1B")
	case 0x1C:
		fmt.Println("Handling OpCode: 0x1C")
	case 0x1D:
		fmt.Println("Handling OpCode: 0x1D")
	case 0x1E:
		fmt.Println("Handling OpCode: 0x1E")
		opbytes = 2
	case 0x1F:
		fmt.Println("Handling OpCode: 0x1F")
	case 0x20:
		fmt.Println("Handling OpCode: 0x20")
	case 0x21:
		fmt.Println("Handling OpCode: 0x21")
		opbytes = 3
	case 0x22:
		fmt.Println("Handling OpCode: 0x22")
		opbytes = 3
	case 0x23:
		fmt.Println("Handling OpCode: 0x23")
	case 0x24:
		fmt.Println("Handling OpCode: 0x24")
	case 0x25:
		fmt.Println("Handling OpCode: 0x25")
	case 0x26:
		fmt.Println("Handling OpCode: 0x26")
		opbytes = 2
	case 0x27:
		fmt.Println("Handling OpCode: 0x27")
	case 0x28:
		fmt.Println("Handling OpCode: 0x28")
	case 0x29:
		fmt.Println("Handling OpCode: 0x29")
	case 0x2A:
		fmt.Println("Handling OpCode: 0x2A")
		opbytes = 3
	case 0x2B:
		fmt.Println("Handling OpCode: 0x2B")
	case 0x2C:
		fmt.Println("Handling OpCode: 0x2C")
	case 0x2D:
		fmt.Println("Handling OpCode: 0x2D")
	case 0x2E:
		fmt.Println("Handling OpCode: 0x2E")
		opbytes = 2
	case 0x2F:
		fmt.Println("Handling OpCode: 0x2F")
	case 0x30:
		fmt.Println("Handling OpCode: 0x30")
	case 0x31:
		fmt.Println("Handling OpCode: 0x31")
		opbytes = 3
	case 0x32:
		fmt.Println("Handling OpCode: 0x32")
		opbytes = 3
	case 0x33:
		fmt.Println("Handling OpCode: 0x33")
	case 0x34:
		fmt.Println("Handling OpCode: 0x34")
	case 0x35:
		fmt.Println("Handling OpCode: 0x35")
	case 0x36:
		fmt.Println("Handling OpCode: 0x36")
		opbytes = 2
	case 0x37:
		fmt.Println("Handling OpCode: 0x37")
	case 0x38:
		fmt.Println("Handling OpCode: 0x38")
	case 0x39:
		fmt.Println("Handling OpCode: 0x39")
	case 0x3A:
		fmt.Println("Handling OpCode: 0x3A")
		opbytes = 3
	case 0x3B:
		fmt.Println("Handling OpCode: 0x3B")
	case 0x3C:
		fmt.Println("Handling OpCode: 0x3C")
	case 0x3D:
		fmt.Println("Handling OpCode: 0x3D")
	case 0x3e:
		fmt.Printf("MVI    A,$%02X\n", buf[0])
		opbytes = 2
	case 0x3F:
		fmt.Println("Handling OpCode: 0x3F")
	case 0x40:
		fmt.Println("Handling OpCode: 0x40")
	case 0x41:
		fmt.Println("Handling OpCode: 0x41")
	case 0x42:
		fmt.Println("Handling OpCode: 0x42")
	case 0x43:
		fmt.Println("Handling OpCode: 0x43")
	case 0x44:
		fmt.Println("Handling OpCode: 0x44")
	case 0x45:
		fmt.Println("Handling OpCode: 0x45")
	case 0x46:
		fmt.Println("Handling OpCode: 0x46")
	case 0x47:
		fmt.Println("Handling OpCode: 0x47")
	case 0x48:
		fmt.Println("Handling OpCode: 0x48")
	case 0x49:
		fmt.Println("Handling OpCode: 0x49")
	case 0x4A:
		fmt.Println("Handling OpCode: 0x4A")
	case 0x4B:
		fmt.Println("Handling OpCode: 0x4B")
	case 0x4C:
		fmt.Println("Handling OpCode: 0x4C")
	case 0x4D:
		fmt.Println("Handling OpCode: 0x4D")
	case 0x4E:
		fmt.Println("Handling OpCode: 0x4E")
	case 0x4F:
		fmt.Println("Handling OpCode: 0x4F")
	case 0x50:
		fmt.Println("Handling OpCode: 0x50")
	case 0x51:
		fmt.Println("Handling OpCode: 0x51")
	case 0x52:
		fmt.Println("Handling OpCode: 0x52")
	case 0x53:
		fmt.Println("Handling OpCode: 0x53")
	case 0x54:
		fmt.Println("Handling OpCode: 0x54")
	case 0x55:
		fmt.Println("Handling OpCode: 0x55")
	case 0x56:
		fmt.Println("Handling OpCode: 0x56")
	case 0x57:
		fmt.Println("Handling OpCode: 0x57")
	case 0x58:
		fmt.Println("Handling OpCode: 0x58")
	case 0x59:
		fmt.Println("Handling OpCode: 0x59")
	case 0x5A:
		fmt.Println("Handling OpCode: 0x5A")
	case 0x5B:
		fmt.Println("Handling OpCode: 0x5B")
	case 0x5C:
		fmt.Println("Handling OpCode: 0x5C")
	case 0x5D:
		fmt.Println("Handling OpCode: 0x5D")
	case 0x5E:
		fmt.Println("Handling OpCode: 0x5E")
	case 0x5F:
		fmt.Println("Handling OpCode: 0x5F")
	case 0x60:
		fmt.Println("Handling OpCode: 0x60")
	case 0x61:
		fmt.Println("Handling OpCode: 0x61")
	case 0x62:
		fmt.Println("Handling OpCode: 0x62")
	case 0x63:
		fmt.Println("Handling OpCode: 0x63")
	case 0x64:
		fmt.Println("Handling OpCode: 0x64")
	case 0x65:
		fmt.Println("Handling OpCode: 0x65")
	case 0x66:
		fmt.Println("Handling OpCode: 0x66")
	case 0x67:
		fmt.Println("Handling OpCode: 0x67")
	case 0x68:
		fmt.Println("Handling OpCode: 0x68")
	case 0x69:
		fmt.Println("Handling OpCode: 0x69")
	case 0x6A:
		fmt.Println("Handling OpCode: 0x6A")
	case 0x6B:
		fmt.Println("Handling OpCode: 0x6B")
	case 0x6C:
		fmt.Println("Handling OpCode: 0x6C")
	case 0x6D:
		fmt.Println("Handling OpCode: 0x6D")
	case 0x6E:
		fmt.Println("Handling OpCode: 0x6E")
	case 0x6F:
		fmt.Println("Handling OpCode: 0x6F")
	case 0x70:
		fmt.Println("Handling OpCode: 0x70")
	case 0x71:
		fmt.Println("Handling OpCode: 0x71")
	case 0x72:
		fmt.Println("Handling OpCode: 0x72")
	case 0x73:
		fmt.Println("Handling OpCode: 0x73")
	case 0x74:
		fmt.Println("Handling OpCode: 0x74")
	case 0x75:
		fmt.Println("Handling OpCode: 0x75")
	case 0x76:
		fmt.Println("Handling OpCode: 0x76")
	case 0x77:
		fmt.Println("Handling OpCode: 0x77")
	case 0x78:
		fmt.Println("Handling OpCode: 0x78")
	case 0x79:
		fmt.Println("Handling OpCode: 0x79")
	case 0x7A:
		fmt.Println("Handling OpCode: 0x7A")
	case 0x7B:
		fmt.Println("Handling OpCode: 0x7B")
	case 0x7C:
		fmt.Println("Handling OpCode: 0x7C")
	case 0x7D:
		fmt.Println("Handling OpCode: 0x7D")
	case 0x7E:
		fmt.Println("Handling OpCode: 0x7E")
	case 0x7F:
		fmt.Println("Handling OpCode: 0x7F")
	case 0x80:
		fmt.Println("Handling OpCode: 0x80")
	case 0x81:
		fmt.Println("Handling OpCode: 0x81")
	case 0x82:
		fmt.Println("Handling OpCode: 0x82")
	case 0x83:
		fmt.Println("Handling OpCode: 0x83")
	case 0x84:
		fmt.Println("Handling OpCode: 0x84")
	case 0x85:
		fmt.Println("Handling OpCode: 0x85")
	case 0x86:
		fmt.Println("Handling OpCode: 0x86")
	case 0x87:
		fmt.Println("Handling OpCode: 0x87")
	case 0x88:
		fmt.Println("Handling OpCode: 0x88")
	case 0x89:
		fmt.Println("Handling OpCode: 0x89")
	case 0x8A:
		fmt.Println("Handling OpCode: 0x8A")
	case 0x8B:
		fmt.Println("Handling OpCode: 0x8B")
	case 0x8C:
		fmt.Println("Handling OpCode: 0x8C")
	case 0x8D:
		fmt.Println("Handling OpCode: 0x8D")
	case 0x8E:
		fmt.Println("Handling OpCode: 0x8E")
	case 0x8F:
		fmt.Println("Handling OpCode: 0x8F")
	case 0x90:
		fmt.Println("Handling OpCode: 0x90")
	case 0x91:
		fmt.Println("Handling OpCode: 0x91")
	case 0x92:
		fmt.Println("Handling OpCode: 0x92")
	case 0x93:
		fmt.Println("Handling OpCode: 0x93")
	case 0x94:
		fmt.Println("Handling OpCode: 0x94")
	case 0x95:
		fmt.Println("Handling OpCode: 0x95")
	case 0x96:
		fmt.Println("Handling OpCode: 0x96")
	case 0x97:
		fmt.Println("Handling OpCode: 0x97")
	case 0x98:
		fmt.Println("Handling OpCode: 0x98")
	case 0x99:
		fmt.Println("Handling OpCode: 0x99")
	case 0x9A:
		fmt.Println("Handling OpCode: 0x9A")
	case 0x9B:
		fmt.Println("Handling OpCode: 0x9B")
	case 0x9C:
		fmt.Println("Handling OpCode: 0x9C")
	case 0x9D:
		fmt.Println("Handling OpCode: 0x9D")
	case 0x9E:
		fmt.Println("Handling OpCode: 0x9E")
	case 0x9F:
		fmt.Println("Handling OpCode: 0x9F")
	case 0xA0:
		fmt.Println("Handling OpCode: 0xA0")
	case 0xA1:
		fmt.Println("Handling OpCode: 0xA1")
	case 0xA2:
		fmt.Println("Handling OpCode: 0xA2")
	case 0xA3:
		fmt.Println("Handling OpCode: 0xA3")
	case 0xA4:
		fmt.Println("Handling OpCode: 0xA4")
	case 0xA5:
		fmt.Println("Handling OpCode: 0xA5")
	case 0xA6:
		fmt.Println("Handling OpCode: 0xA6")
	case 0xA7:
		fmt.Println("Handling OpCode: 0xA7")
	case 0xA8:
		fmt.Println("Handling OpCode: 0xA8")
	case 0xA9:
		fmt.Println("Handling OpCode: 0xA9")
	case 0xAA:
		fmt.Println("Handling OpCode: 0xAA")
	case 0xAB:
		fmt.Println("Handling OpCode: 0xAB")
	case 0xAC:
		fmt.Println("Handling OpCode: 0xAC")
	case 0xAD:
		fmt.Println("Handling OpCode: 0xAD")
	case 0xAE:
		fmt.Println("Handling OpCode: 0xAE")
	case 0xAF:
		fmt.Println("Handling OpCode: 0xAF")
	case 0xB0:
		fmt.Println("Handling OpCode: 0xB0")
	case 0xB1:
		fmt.Println("Handling OpCode: 0xB1")
	case 0xB2:
		fmt.Println("Handling OpCode: 0xB2")
	case 0xB3:
		fmt.Println("Handling OpCode: 0xB3")
	case 0xB4:
		fmt.Println("Handling OpCode: 0xB4")
	case 0xB5:
		fmt.Println("Handling OpCode: 0xB5")
	case 0xB6:
		fmt.Println("Handling OpCode: 0xB6")
	case 0xB7:
		fmt.Println("Handling OpCode: 0xB7")
	case 0xB8:
		fmt.Println("Handling OpCode: 0xB8")
	case 0xB9:
		fmt.Println("Handling OpCode: 0xB9")
	case 0xBA:
		fmt.Println("Handling OpCode: 0xBA")
	case 0xBB:
		fmt.Println("Handling OpCode: 0xBB")
	case 0xBC:
		fmt.Println("Handling OpCode: 0xBC")
	case 0xBD:
		fmt.Println("Handling OpCode: 0xBD")
	case 0xBE:
		fmt.Println("Handling OpCode: 0xBE")
	case 0xBF:
		fmt.Println("Handling OpCode: 0xBF")
	case 0xC0:
		fmt.Println("Handling OpCode: 0xC0")
	case 0xC1:
		fmt.Println("Handling OpCode: 0xC1")
	case 0xC2:
		fmt.Println("Handling OpCode: 0xC2")
		opbytes = 3
	case 0xC3:
		fmt.Printf("JMP    $%02X%02X\n", buf[1], buf[0])
		opbytes = 3
	case 0xC4:
		fmt.Println("Handling OpCode: 0xC4")
		opbytes = 3
	case 0xC5:
		fmt.Println("Handling OpCode: 0xC5")
	case 0xC6:
		fmt.Println("Handling OpCode: 0xC6")
		opbytes = 2
	case 0xC7:
		fmt.Println("Handling OpCode: 0xC7")
	case 0xC8:
		fmt.Println("Handling OpCode: 0xC8")
	case 0xC9:
		fmt.Println("Handling OpCode: 0xC9")
	case 0xCA:
		fmt.Println("Handling OpCode: 0xCA")
		opbytes = 3
	case 0xCB:
		fmt.Println("Handling OpCode: 0xCB")
	case 0xCC:
		fmt.Println("Handling OpCode: 0xCC")
		opbytes = 3
	case 0xCD:
		fmt.Println("Handling OpCode: 0xCD")
		opbytes = 3
	case 0xCE:
		fmt.Println("Handling OpCode: 0xCE")
		opbytes = 2
	case 0xCF:
		fmt.Println("Handling OpCode: 0xCF")
	case 0xD0:
		fmt.Println("Handling OpCode: 0xD0")
	case 0xD1:
		fmt.Println("Handling OpCode: 0xD1")
	case 0xD2:
		fmt.Println("Handling OpCode: 0xD2")
		opbytes = 3
	case 0xD3:
		fmt.Println("Handling OpCode: 0xD3")
		opbytes = 2
	case 0xD4:
		fmt.Println("Handling OpCode: 0xD4")
		opbytes = 3
	case 0xD5:
		fmt.Println("Handling OpCode: 0xD5")
	case 0xD6:
		fmt.Println("Handling OpCode: 0xD6")
		opbytes = 2
	case 0xD7:
		fmt.Println("Handling OpCode: 0xD7")
	case 0xD8:
		fmt.Println("Handling OpCode: 0xD8")
	case 0xD9:
		fmt.Println("Handling OpCode: 0xD9")
	case 0xDA:
		fmt.Println("Handling OpCode: 0xDA")
		opbytes = 3
	case 0xDB:
		fmt.Println("Handling OpCode: 0xDB")
		opbytes = 2
	case 0xDC:
		fmt.Println("Handling OpCode: 0xDC")
		opbytes = 3
	case 0xDD:
		fmt.Println("Handling OpCode: 0xDD")
	case 0xDE:
		fmt.Println("Handling OpCode: 0xDE")
		opbytes = 2
	case 0xDF:
		fmt.Println("Handling OpCode: 0xDF")
	case 0xE0:
		fmt.Println("Handling OpCode: 0xE0")
	case 0xE1:
		fmt.Println("Handling OpCode: 0xE1")
	case 0xE2:
		fmt.Println("Handling OpCode: 0xE2")
		opbytes = 3
	case 0xE3:
		fmt.Println("Handling OpCode: 0xE3")
	case 0xE4:
		fmt.Println("Handling OpCode: 0xE4")
		opbytes = 3
	case 0xE5:
		fmt.Println("Handling OpCode: 0xE5")
	case 0xE6:
		fmt.Println("Handling OpCode: 0xE6")
		opbytes = 2
	case 0xE7:
		fmt.Println("Handling OpCode: 0xE7")
	case 0xE8:
		fmt.Println("Handling OpCode: 0xE8")
	case 0xE9:
		fmt.Println("Handling OpCode: 0xE9")
	case 0xEA:
		fmt.Println("Handling OpCode: 0xEA")
		opbytes = 3
	case 0xEB:
		fmt.Println("Handling OpCode: 0xEB")
	case 0xEC:
		fmt.Println("Handling OpCode: 0xEC")
		opbytes = 3
	case 0xED:
		fmt.Println("Handling OpCode: 0xED")
	case 0xEE:
		fmt.Println("Handling OpCode: 0xEE")
		opbytes = 2
	case 0xEF:
		fmt.Println("Handling OpCode: 0xEF")
	case 0xF0:
		fmt.Println("Handling OpCode: 0xF0")
	case 0xF1:
		fmt.Println("Handling OpCode: 0xF1")
	case 0xF2:
		fmt.Println("Handling OpCode: 0xF2")
		opbytes = 3
	case 0xF3:
		fmt.Println("Handling OpCode: 0xF3")
	case 0xF4:
		fmt.Println("Handling OpCode: 0xF4")
		opbytes = 3
	case 0xF5:
		fmt.Println("Handling OpCode: 0xF5")
	case 0xF6:
		fmt.Println("Handling OpCode: 0xF6")
		opbytes = 2
	case 0xF7:
		fmt.Println("Handling OpCode: 0xF7")
	case 0xF8:
		fmt.Println("Handling OpCode: 0xF8")
	case 0xF9:
		fmt.Println("Handling OpCode: 0xF9")
	case 0xFA:
		fmt.Println("Handling OpCode: 0xFA")
		opbytes = 3
	case 0xFB:
		fmt.Println("Handling OpCode: 0xFB")
	case 0xFC:
		fmt.Println("Handling OpCode: 0xFC")
		opbytes = 3
	case 0xFD:
		fmt.Println("Handling OpCode: 0xFD")
	case 0xFE:
		fmt.Println("Handling OpCode: 0xFE")
		opbytes = 2
	case 0xFF:
		fmt.Println("Handling OpCode: 0xFF")
	default:
		panic("Unknown OpCode")
	}

	return opbytes
}
