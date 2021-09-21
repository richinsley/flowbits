package flowbits

import (
	"io"
)

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

func (me *Bitstream) seterror(err Error_t) {
	me.err_code = err
}

// availableBufferBits returns the count of BITs that are available in the internal read buffer
func (me *Bitstream) availableBufferBits() uint64 {
	return uint64(me.buf_len)<<uint64(BSHIFT) - uint64(me.cur_bit)
}

func (me *Bitstream) checkReadEnoughAvailable(n uint32) error {
	// ensure we have enough data to read an entire 64 bits.
	// we are going to walk cur_bit forward 8 bits for every read so we'll
	// need to walk back cur_bit before we return.  If NextBitsUnsignedLittle were to fill_buf
	// while we were reading the bytes, this would break that.  We'll ensure there are
	// at least 64-bits (uint64 + 8 bits) available to read BEFORE we start probing to prevent that
	need := uint64(n)
	available := me.availableBufferBits()
	if available < need {
		// refill and check available again
		me.FillReadBuffer()
		available = me.availableBufferBits()
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
