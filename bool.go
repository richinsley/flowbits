package flowbits

// NextBool will return the next boolean value in the bitstream without moving the bit read position.
func (me *Bitstream) NextBool() (bool, error) {
	// check if we have enough bits available for the operation
	// if skip_check is true, another method already checked beforehand
	var err error = nil
	if !me.skip_check {
		err := me.checkReadEnoughAvailable(1)
		if err != nil {
			return false, err
		}
	}

	bpos := me.cur_bit % 8
	return me.buf[me.cur_bit>>BSHIFT]&charbitmask[bpos] != 0, err
}

// GetBool will return the next boolean value in the bitstream.  The bit read position is advanced by one.
func (me *Bitstream) GetBool() (bool, error) {
	// check if we have enough bits available for the operation
	// if skip_check is true, another method already checked beforehand
	var err error = nil
	if !me.skip_check {
		err := me.checkReadEnoughAvailable(1)
		if err != nil {
			return false, err
		}
	}

	if err != nil {
		return false, err
	}
	bpos := me.cur_bit % 8
	retv := me.buf[me.cur_bit>>BSHIFT]&charbitmask[bpos] != 0
	me.cur_bit++
	return retv, err
}

// PutBool will write a single bit boolean value to the bitstream.
func (me *Bitstream) PutBool(value bool) error {
	var err error = nil
	if me.availableBufferBits() < uint64(1) {
		err = me.flush_write_buffer()
	}

	// if we're putting TRUE, or in the charbitmask
	// if we're putting FALSE, just increment the cur_bit because the underlying value
	// should already be FALSE
	bpos := me.cur_bit % 8
	if value {

		me.buf[me.cur_bit>>BSHIFT] |= charbitmask[bpos]
	}
	me.cur_bit++
	return err
}
