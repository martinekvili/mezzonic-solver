package utils

func TestBit(vector uint32, index uint8) bool {
	return vector&(1<<index) > 0
}

func SetBit(vector uint32, index uint8) uint32 {
	return vector | 1<<index
}

func ClearBit(vector uint32, index uint8) uint32 {
	return vector & ^(1 << index)
}
