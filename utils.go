package flobits

func fill_slice(s []uint8, c uint8) {
	for i := range s {
		s[i] = c
	}
}

func minUint64(x uint64, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

// compare two buffers of uint8 bytes
func unsigned_buffersEqual(b1 []uint8, b2 []uint8) bool {
	for i, b := range b1 {
		if b != b2[i] {
			return false
		}
	}

	return true
}

// compare two buffers of int8 bytes
func signed_buffersEqual(b1 []int8, b2 []int8) bool {
	for i, b := range b1 {
		if b != b2[i] {
			return false
		}
	}

	return true
}

// compare two uint64 arrays
func buffers64Equal(b1 []uint64, b2 []uint64) bool {
	for i, b := range b1 {
		if b != b2[i] {
			return false
		}
	}

	return true
}

// compare two buffers of uint8 bytes - (ugh, we can't get gerics fast enough)
func buffers32Equal(b1 []uint32, b2 []uint32) bool {
	for i, b := range b1 {
		if b != b2[i] {
			return false
		}
	}

	return true
}
