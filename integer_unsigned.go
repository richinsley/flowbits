package flobits

import (
	"fmt"
	"io"
)

// NextBitsUnsignedBig returns 'n' bits in big endian order as an unsigned int.
// This does not advance bit pointer.
func (me *Flobitsstream) NextBitsUnsignedBig(n uint32) (uint64, error) {
	var x uint64    // the value we will return
	var v []uint8   // the byte where cur_bit points to
	var delta int32 // number of bits to shift

	// check if we have enough bits available for the operation
	// if skip_check is true, another method already checked beforehand
	if !me.skip_check {
		err := me.checkEnoughAvailable(n)
		if err != nil {
			return 0, err
		}
	}
	// starting byte in buffer - the internal buffer may be larger than the actual available data in it,
	// so we'll clamp it to the value of buf_len
	v = me.buf[me.cur_bit>>BSHIFT : me.buf_len]

	// load up to 8 bytes at a time - this way endianess is automatically taken care of.
	// we'll use a switch statement and unroll loops for speed
	byterange := minUint64(uint64(len(v)), 8)
	switch byterange {
	case 8:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40) |
			(uint64(v[3]) << 32) |
			(uint64(v[4]) << 24) |
			(uint64(v[5]) << 16) |
			(uint64(v[6]) << 8) |
			uint64(v[7])
	case 7:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40) |
			(uint64(v[3]) << 32) |
			(uint64(v[4]) << 24) |
			(uint64(v[5]) << 16) |
			(uint64(v[6]) << 8)
	case 6:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40) |
			(uint64(v[3]) << 32) |
			(uint64(v[4]) << 24) |
			(uint64(v[5]) << 16)
	case 5:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40) |
			(uint64(v[3]) << 32) |
			(uint64(v[4]) << 24)
	case 4:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40) |
			(uint64(v[3]) << 32)
	case 3:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48) |
			(uint64(v[2]) << 40)
	case 2:
		x = (uint64(v[0]) << 56) |
			(uint64(v[1]) << 48)
	case 1:
		x = (uint64(v[0]) << 56)
	default:
		x = 0
	}

	// figure out how much shifting is required
	delta = 64 - int32((me.cur_bit%8)+n)

	if delta >= 0 {
		x = (x >> delta) // need right shift to get proper value
	} else {
		// shift left and read an extra byte
		x = x << -delta
		x |= uint64(v[8]) >> (8 + delta)
	}
	return x & mask[n], nil
}

func (me *Flobitsstream) PutBitsUnsignedBig(value uint64, n uint32) uint64 {
	var delta int                    // required input shift amount
	var v []uint8                    // current byte
	var tmp uint64                   // temp value for shifted bits
	var val uint64 = value & mask[n] // the n-bit value

	if me.AvailableBufferBits() < 64 {
		me.flush_buf()
	}

	// delta can be negative, so cast uints to int
	delta = 64 - int(n) - (int(me.cur_bit) % 8)
	v = me.buf[me.cur_bit>>BSHIFT:]

	if delta >= 0 {
		tmp = val << delta
		v[0] |= uint8(tmp >> 56)
		v[1] |= uint8(tmp >> 48)
		v[2] |= uint8(tmp >> 40)
		v[3] |= uint8(tmp >> 32)
		v[4] |= uint8(tmp >> 24)
		v[5] |= uint8(tmp >> 16)
		v[6] |= uint8(tmp >> 8)
		v[7] |= uint8(tmp)
	} else {
		tmp = val >> (-delta)
		v[0] |= uint8(tmp >> 56)
		v[1] |= uint8(tmp >> 48)
		v[2] |= uint8(tmp >> 40)
		v[3] |= uint8(tmp >> 32)
		v[4] |= uint8(tmp >> 24)
		v[5] |= uint8(tmp >> 16)
		v[6] |= uint8(tmp >> 8)
		v[7] |= uint8(tmp)
		v[8] |= (uint8)(value << (8 + delta))
	}

	me.cur_bit += n
	me.tot_bits += uint64(n)
	return value
}

func (me *Flobitsstream) GetBitsUnsignedBig(n uint32) (uint64, error) {
	x, err := me.NextBitsUnsignedBig(n)
	me.cur_bit += n
	me.tot_bits += uint64(n)
	return x & mask[n], err
}

func (me *Flobitsstream) PutBitsUnsignedLittle(value uint64, n uint32) uint64 {
	var bytes uint32 = n >> uint32(BSHIFT)
	var leftbits uint32 = n % 8
	var byte_x uint64 = 0
	var i uint32
	for i = 0; i < bytes; i++ {
		byte_x = (value >> (8 * i)) & mask[8]
		me.PutBitsUnsignedBig(byte_x, 8)
	}
	if leftbits > 0 {
		byte_x = (value >> (8 * i)) & mask[leftbits]
		me.PutBitsUnsignedBig(byte_x, leftbits)
	}
	return value
}

func (me *Flobitsstream) GetBitsUnsignedLittle(n uint32) (uint64, error) {
	var x uint64 = 0               // the value we will return
	var bytes uint32 = n >> BSHIFT // number of bytes to read
	var leftbits uint32 = n % 8    // number of bits to read
	var byte_x uint64 = 0
	var i uint32 = 0

	// check if we have enough bits available for the operation
	err := me.checkEnoughAvailable(n)
	if err != nil {
		return 0, err
	}

	// we checked for available above, so we can ignore errors from GetBitsUnsignedBig
	me.skip_check = true
	for i = 0; i < bytes; i++ {
		byte_x, _ = me.GetBitsUnsignedBig(8)
		byte_x <<= (8 * i)
		x |= byte_x
	}

	if leftbits > 0 {
		byte_x, _ = me.GetBitsUnsignedBig(leftbits)
		byte_x <<= (8 * i)
		x |= byte_x
	}
	me.skip_check = false
	return x, nil
}

func (me *Flobitsstream) checkEnoughAvailable(n uint32) error {
	// ensure we have enough data to read an entire 64 bits.
	// we are going to walk cur_bit forward 8 bits for every read so we'll
	// need to walk back cur_bit before we return.  If NextBitsUnsignedLittle were to fill_buf
	// while we were reading the bytes, this would break that.  We'll ensure there are
	// at least 64-bits (uint64 + 8 bits) available to read BEFORE we start probing to prevent that
	need := uint64(n)
	available := me.AvailableBufferBits()
	if available < need {
		// refill and check available again
		me.FillBuffer()
		available = me.AvailableBufferBits()
		if available < need {
			return io.EOF
		}
	}

	if me.err_code != E_NONE {
		if me.err_code == E_END_OF_DATA {
			// re-get available after fill
			// available = me.AvailableBufferBits()
			if available == 0 {
				return io.EOF
			} else if available < uint64(n) {
				return io.ErrShortBuffer
			}
		} else {
			// read failed for unknown reasons
			return io.ErrUnexpectedEOF
		}
	}

	return nil
}

// returns 'n' bits as unsigned int
// does not advance bit pointer
func (me *Flobitsstream) NextBitsUnsignedLittle(n uint32) (uint64, error) {
	var x uint64 = 0               // the value we will return
	var bytes uint32 = n >> BSHIFT // number of bytes to read
	var leftbits uint32 = n % 8    // number of left-over bits to read
	var byte_x uint64 = 0
	var i uint32

	// check if we have enough bits available for the operation
	err := me.checkEnoughAvailable(n)
	if err != nil {
		return 0, err
	}

	// we checked for available above, so we can ignore errors from GetBitsUnsignedBig
	// we will also be walking back the cur_bit, so we don't want the internal buffer to be filled while
	// we're performing this operation
	me.skip_check = true
	for i = 0; i < bytes; i++ {
		byte_x, _ = me.NextBitsUnsignedBig(8)
		me.cur_bit += 8
		byte_x <<= (8 * i)
		x |= byte_x
	}

	// Note that it doesn't make much sense to have a number in little-endian
	// byte-ordering, where the number of bits used to represent the number is
	// not a multiple of 8.  Neverthless, we provide a way to take care of
	// such case.
	if leftbits > 0 {
		byte_x, _ = me.NextBitsUnsignedBig(leftbits)
		byte_x <<= (8 * i)
		x |= byte_x
	}
	me.skip_check = false

	// we temporarily moved the cur_bit value above, so now we need to step it back
	back := i * 8
	if back <= me.cur_bit {
		me.cur_bit = me.cur_bit - (i * 8)
	} else {
		fmt.Println("this is a problem")
		me.cur_bit = 0
	}
	return x, nil
}
