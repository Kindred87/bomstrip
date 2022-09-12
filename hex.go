package bomtrim

type hexes int

const (
	ef hexes = iota
	bb
	bf
	ff
	fe
	zeroZero
	zeroZeroZeroZero
	twoB
	twoF
	sevenSix
	f7
	sixFour
	fourC
	dd
	sevenThree
	sixSix
	zeroE
	twoEight
	eightFour
	threeOne
	nineFive
	threeThree
	endList
)

// hex returns the byte equivalent of the hexadecimal.
func (h hexes) hex() []byte {
	switch h {
	case ef:
		return []byte{0xEF}
	case bb:
		return []byte{0xBB}
	case bf:
		return []byte{0xBF}
	case ff:
		return []byte{0xFF}
	case fe:
		return []byte{0xFE}
	case zeroZero:
		return []byte{0x00}
	case zeroZeroZeroZero:
		return []byte{0x0000}
	case twoB:
		return []byte{0x2B}
	case twoF:
		return []byte{0x2F}
	case sevenSix:
		return []byte{0x76}
	case f7:
		return []byte{0xF7}
	case sixFour:
		return []byte{0x64}
	case fourC:
		return []byte{0x4C}
	case dd:
		return []byte{0xDD}
	case sevenThree:
		return []byte{0x73}
	case sixSix:
		return []byte{0x66}
	case zeroE:
		return []byte{0x0E}
	case twoEight:
		return []byte{0x28}
	case eightFour:
		return []byte{0x84}
	case threeOne:
		return []byte{0x31}
	case nineFive:
		return []byte{0x95}
	case threeThree:
		return []byte{0x33}
	default:
		return []byte("")
	}
}

// string returns the string equivalent of the hexadecimal.
func (h hexes) string() string {
	switch h {
	case ef:
		return "0xEF"
	case bb:
		return "0xBB"
	case bf:
		return "0xBF"
	case ff:
		return "0xFF"
	case fe:
		return "0xFE"
	case zeroZero:
		return "0x00"
	case zeroZeroZeroZero:
		return "0x0000"
	case twoB:
		return "0x2B"
	case twoF:
		return "0x2F"
	case sevenSix:
		return "0x76"
	case f7:
		return "0xF7"
	case sixFour:
		return "0x64"
	case fourC:
		return "0x4C"
	case dd:
		return "0xDD"
	case sevenThree:
		return "0x73"
	case sixSix:
		return "0x66"
	case zeroE:
		return "0x0E"
	case twoEight:
		return "0x28"
	case eightFour:
		return "0x84"
	case threeOne:
		return "0x31"
	case nineFive:
		return "0x95"
	case threeThree:
		return "0x33"
	default:
		return ""
	}
}
