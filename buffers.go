package flobits

import "io"

// GetBuffer reads 'n' bytes into the given buffer and returns count of bytes read.
func (me *Bitstream) GetBuffer(buffer []uint8, n uint64) (uint64, error) {
	if n == 0 {
		return 0, nil
	}

	var total_bytes_read uint64 = 0
	if me.cur_bit%8 != 0 {
		var i uint64
		for i = 0; i < n; i++ {
			lbits, err := me.GetBitsUnsignedBig(8)
			if err != nil {
				return total_bytes_read, err
			}
			buffer[i] = uint8(lbits)
			total_bytes_read++
		}
	} else {
		var leftover uint64 = 0
		// see if we have any available bytes in our buffer
		mbufsize := minUint64(uint64(me.buf_len-(me.cur_bit>>BSHIFT)), n)

		if mbufsize != 0 {
			// buf is fixed size
			// cur_bit is the current BIT position (read/write) in buf
			// (cur_bit >> BSHIFT) will give the current BYTE position (read/write)

			// create a slice of me.buf starting from the read byte position up to the end
			readpos := (me.cur_bit >> BSHIFT)
			endpos := readpos + uint32(mbufsize)
			v := me.buf[readpos:endpos]

			// copy that slice into the target buffer starting at it's begining
			copied := copy(buffer[:mbufsize], v)

			total_bytes_read += uint64(copied)

			// adjust internal cur_bit pointer
			me.cur_bit += uint32(mbufsize) << BSHIFT

			// reduce the size parameter by the amount we read from our internal buffer
			leftover = n - mbufsize
		} else {
			// the current read buffer was empty so the entire buffer needs to read from the reader
			leftover = n
		}

		if leftover != 0 {
			// read any remaining bytes directly from the reader
			v := buffer[mbufsize:]
			l, err := me.reader.Read(v)
			total_bytes_read += uint64(l)
			if err != nil {
				me.end = true
				me.seterror(E_READ_FAILED)
				return total_bytes_read, io.EOF
			}
		}

		me.tot_bits += total_bytes_read << BSHIFT
	}

	return total_bytes_read, nil
}

// PutBuffer writes the given buffer to the bitstream advancing the write position to current+(8*buffer_length).
func (me *Bitstream) PutBuffer(buffer []uint8) uint64 {
	size := uint64(len(buffer))
	if size == 0 {
		return 0
	}

	if me.cur_bit%8 != 0 {
		// we are not byte aligned.  Putbits each byte individually
		for _, v := range buffer {
			me.PutBitsUnsignedBig(uint64(v), 8)
		}
		return size
	}

	// we are byte aligned, just flush the internal buffer and write the entire given buffer directly to the writer
	me.flush_buf()
	wsize, err := me.writer.Write(buffer)
	me.cur_bit = 0
	me.tot_bits += uint64(wsize) >> BSHIFT
	if err != nil {
		me.seterror(E_WRITE_FAILED)
	}
	return uint64(wsize)
}
