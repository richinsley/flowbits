package flobits

// NextBitsUnsignedBig returns 'n' bits in big endian order as an unsigned int.
// This does not advance bit pointer.
func (me *Flobitsstream) NextBitsUnsignedBig(n uint32) uint64 {
	var x uint64    // the value we will return
	var v []uint8   // the byte where cur_bit points to
	var delta int32 // number of bits to shift

	// make sure we have enough data
	if me.availableBufferBits() < 64 {
		// try to fill the buffer and handle any errors it may encounter
		me.fill_buf()

		if me.err_code == E_END_OF_DATA {
			// how many bits do we actually have after the failed the fill_buf?
			remaining := uint32(me.availableBufferBits())
			if remaining == 0 {
				return 0
			} else if remaining < n {
				// we have some left, but not enough
				return 0
			}
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
	return x & mask[n]
}

func (me *Flobitsstream) PutBitsUnsignedBig(value uint64, n uint32) uint64 {
	var delta int                    // required input shift amount
	var v []uint8                    // current byte
	var tmp uint64                   // temp value for shifted bits
	var val uint64 = value & mask[n] // the n-bit value

	if me.availableBufferBits() < 64 {
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

func (me *Flobitsstream) GetBitsUnsignedBig(n uint32) uint64 {
	var x uint64 = me.NextBitsUnsignedBig(n)
	me.cur_bit += n
	me.tot_bits += uint64(n)
	return x & mask[n]
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

func (me *Flobitsstream) GetBitsUnsignedLittle(n uint32) uint64 {
	var x uint64 = 0               // the value we will return
	var bytes uint32 = n >> BSHIFT // number of bytes to read
	var leftbits uint32 = n % 8    // number of bits to read
	var byte_x uint64 = 0
	var i uint32 = 0
	for i = 0; i < bytes; i++ {
		byte_x = me.GetBitsUnsignedBig(8)
		byte_x <<= (8 * i)
		x |= byte_x
	}

	if leftbits > 0 {
		byte_x = me.GetBitsUnsignedBig(leftbits)
		byte_x <<= (8 * i)
		x |= byte_x
	}
	return x
}

// returns 'n' bits as unsigned int
// does not advance bit pointer
func (me *Flobitsstream) NextBitsUnsignedLittle(n uint32) uint64 {
	var x uint64 = 0               // the value we will return
	var bytes uint32 = n >> BSHIFT // number of bytes to read
	var leftbits uint32 = n % 8    // number of left-over bits to read
	var byte_x uint64 = 0
	var i uint32

	// ensure we have enough data to read an entire 64 bits.
	// we are going to walk cur_bit forward 8 bits for every read so we'll
	// need to walk back cur_bit before we return.  If NextBitsUnsignedLittle were to fill_buf
	// while we were reading the bytes, this would break that.  We'll ensure there are
	// at least 128-bits available to read BEFORE we start probing to prevent that
	available := me.availableBufferBits()
	if available <= 128 {
		me.fill_buf()
	}

	for i = 0; i < bytes; i++ {
		byte_x = me.NextBitsUnsignedBig(8)
		me.cur_bit += 8
		byte_x <<= (8 * i)
		x |= byte_x
	}

	// Note that it doesn't make much sense to have a number in little-endian
	// byte-ordering, where the number of bits used to represent the number is
	// not a multiple of 8.  Neverthless, we provide a way to take care of
	// such case.
	if leftbits > 0 {
		byte_x = me.NextBitsUnsignedBig(leftbits)
		byte_x <<= (8 * i)
		x |= byte_x
	}

	// we temporarily moved the cur_bit value above, so now we need to step it back
	me.cur_bit -= i * 8
	return x
}
