// Package Flowbits provides a robust, multi-purpose bitstream parser and encoder that can handle:
//		Big and little endianness
//		32 bit and 64 bit floats
//		Signed and unsigned ints from 1-64 bits in size
//		Booleans
//		Random access
//		Binary code search
package flobits

import (
	"errors"
	"io"
	"math"
)

const (
	BS_OUTPUT  uint32 = 1
	BS_INPUT   uint32 = 2
	BSHIFT     uint32 = 3
	BS_BUF_LEN uint32 = 32 // the minimum size of the internal read/write buffer - enough for 4 uint64
)

type Error_t int

const (
	E_NONE        Error_t = iota
	E_END_OF_DATA         // the underlying read device has no more available data, but the internal buffer may still contain data
	E_INVALID_ALIGNMENT
	E_READ_FAILED
	E_WRITE_FAILED
	E_SEEK_FAILED
)

// Bistream is a bitstream encoder and decoder.
type Bitstream struct {
	writer   io.Writer
	reader   io.Reader
	buf_len  uint32 // usable buffer size (for partially filled buffers)
	cur_bit  uint32 // current bit position in buf
	tot_bits uint64 // total bits read/written
	end      bool   // end of data flag
	// zcount   uint32 // number of zeros counted on most recent countZero
	iotype     uint32
	buf        []uint8
	err_code   Error_t
	seeker     io.Seeker // seeker interface for given reader/writer if available
	skip_check bool      // when true, skip checking for available data
	max_len    uint32    // max internal buffer size (may be larger than buf_len)
}

// NewBitstreamEncoder creates a new Bitstream encoder instance with the given io.writer and internal buffer size.
func NewBitstreamEncoder(w io.Writer, buffer_len int) *Bitstream {
	genMasks()

	// see if the writer also has the seeker interface
	seeker, _ := interface{}(w).(io.Seeker)

	nbuffer_len := uint32(math.Max(float64(BS_BUF_LEN), float64(buffer_len)))
	return &Bitstream{
		writer:   w,
		reader:   nil,
		buf_len:  nbuffer_len,
		cur_bit:  0,
		buf:      make([]uint8, nbuffer_len),
		end:      false,
		iotype:   BS_OUTPUT,
		err_code: E_NONE,
		seeker:   seeker,
		max_len:  nbuffer_len,
	}
}

// NewBitstreamDecoder creates a new Bitstream decoder instance with the given io.reader and internal buffer size.
func NewBitstreamDecoder(r io.Reader, buffer_len int) *Bitstream {
	genMasks()

	// see if the reader also has the seeker interface
	seeker, _ := interface{}(r).(io.Seeker)

	nbuffer_len := uint32(math.Max(float64(BS_BUF_LEN), float64(buffer_len)))
	retv := &Bitstream{
		writer:   nil,
		reader:   r,
		buf_len:  nbuffer_len,
		cur_bit:  0,
		buf:      make([]uint8, nbuffer_len),
		end:      false,
		iotype:   BS_INPUT,
		err_code: E_NONE,
		seeker:   seeker,
		max_len:  nbuffer_len,
	}

	// fake that we are at the end of buffer before we call fill_buff
	// or else fill_buf will think we have the entire available buffer as
	// ready to read data
	retv.cur_bit = nbuffer_len << BSHIFT
	retv.FillBuffer()

	return retv
}

func (me *Bitstream) seterror(err Error_t) {
	me.err_code = err
}

// availableBufferBits returns the count of BITs that are available in the internal read buffer
func (me *Bitstream) availableBufferBits() uint64 {
	return uint64(me.buf_len)<<uint64(BSHIFT) - uint64(me.cur_bit)
}

// FillBuffer re-fills the internal buffer in cases where a io.writer as added additional data
// to an io.reader used by a Bitstream decoder.
func (me *Bitstream) FillBuffer() error {
	if me.iotype == BS_OUTPUT {
		return nil
	}

	// reset the err code incase more data was added after the last call
	me.err_code = E_NONE

	var n uint32 = me.cur_bit >> BSHIFT // number of spent bytes in the buffer
	var l int                           // how many bytes we did fetch (available)
	var u uint32 = me.buf_len - n       // how many bytes are still unread
	var err error

	if u != 0 {
		// shift unread contents to the beginning of the buffer
		copy(me.buf, me.buf[n:])
	}

	// zero out the rest of buf
	fill_slice(me.buf[u:], 0)

	// we've moved all remaining data to the beginning, so the cur_bit should point to byte 0
	me.cur_bit &= 7

	l, err = me.reader.Read(me.buf[u:])

	if l == 0 {
		// we can read no bytes
		me.end = true
		me.buf_len = u // the buffer length is whatever was left over
		me.seterror(E_END_OF_DATA)
		return io.EOF
	} else if uint32(l) < me.max_len {
		// we got some, so we'll assume this is the end of the stream, but not an io error
		me.end = true
		me.buf_len = u + uint32(l)
		me.seterror(E_END_OF_DATA)
		return nil
	} else if err != nil {
		// the dog pulled the plug out of the wall again
		me.end = true
		me.seterror(E_READ_FAILED)
		return io.ErrUnexpectedEOF
	}

	return nil
}

// flush_buf flushes the internal buffer and outputs the buffer excluding the left-over bits
func (me *Bitstream) flush_buf() error {
	if me.iotype == BS_OUTPUT {
		var l int = int(me.cur_bit >> BSHIFT) // number of bytes written already
		n, err := me.writer.Write(me.buf[:l])

		if err != nil {
			me.seterror(E_WRITE_FAILED)
			return err
		}

		if n != l {
			me.seterror(E_WRITE_FAILED)
			return io.ErrShortWrite
		}

		// are there any left over bits?
		if me.cur_bit&0x7 != 0 {
			me.buf[0] = me.buf[l]     // copy the left-over bits
			fill_slice(me.buf[1:], 0) // zero-out rest of buffer
		} else {
			fill_slice(me.buf[:], 0) // zero-out entire buffer
		}
		me.cur_bit &= 7
	}
	return nil
}

// GetError returns the most recent internal error code.
func (me *Bitstream) GetError() Error_t {
	return me.err_code
}

// Flushbits flushes the buffer.  If the current bit position is not byte aligned, the left-over
// bits are also output with zero padding.
func (me *Bitstream) Flushbits() error {
	err := me.flush_buf()

	if me.cur_bit == 0 || err != nil {
		return err
	}

	l, err := me.writer.Write(me.buf[:1])
	if l != 1 || err != nil {
		me.seterror(E_WRITE_FAILED)
		return errors.New("write failed")
	}
	me.buf[0] = 0
	me.cur_bit = 0

	return nil
}

// Skipbits advances the read/write bit position n bits.
func (me *Bitstream) Skipbits(n uint32) error {
	x := n
	buf_size := me.buf_len << BSHIFT
	var err error = nil

	// make sure we have enough data
	for me.cur_bit+x > buf_size {
		x -= (buf_size - me.cur_bit)
		me.cur_bit = buf_size
		if me.iotype == BS_INPUT {
			err = me.FillBuffer()
		} else {
			err = me.flush_buf()
		}

		if err != nil {
			return err
		}
	}
	me.cur_bit += x
	me.tot_bits += uint64(n)

	return nil
}

// Align will align the read or write bit position.
// alignToBits must be multiple of 8
// returns number of bits skipped
func (me *Bitstream) Align(alignToBits uint64) uint64 {
	var s uint64 = 0

	// we only allow alignment on multiples of bytes
	if alignToBits%8 != 0 {
		me.seterror(E_INVALID_ALIGNMENT)
		return 0
	}

	// align on next byte
	if me.tot_bits%8 != 0 {
		s = 8 - (me.tot_bits % 8)
		me.Skipbits(uint32(s))
	}

	for (me.tot_bits % alignToBits) != 0 {
		me.Skipbits(8)
		s = s + 8
	}
	return s
}

// CanSeek returns true if reader/writer supports random access seeking
func (me *Bitstream) CanSeek() bool {
	return me.seeker != nil
}

// Tell returns the current read/write position in the io device in BITS or -1 if not seekable.
func (me *Bitstream) Tell() int64 {
	if !me.CanSeek() {
		return -1
	}

	// seeking to 0 with io.SeekCurrent will return the current device potition
	offset, err := me.seeker.Seek(0, io.SeekCurrent)
	if err != nil {
		me.seterror(E_SEEK_FAILED)
		return -1
	}

	if me.iotype == BS_INPUT {
		// the offset represents the total bytes that have been read into the internal buffer overall
		// we need to subtract the size of the internal buffer and add back the cur_bit position
		return int64(offset<<BSHIFT) - int64(me.buf_len<<BSHIFT) + int64(me.cur_bit)
	}

	return offset*8 + int64(me.cur_bit)
}

// GetPos returns the current absolute read or write position in bits.  For input streams, this will be the
// next bit read.  For output streams, this will be the position the bit will be written.
func (me *Bitstream) GetPos() uint64 {
	if me.CanSeek() {
		return uint64(me.Tell())
	}

	return me.tot_bits
}

// EOF returns true if at bitstream is EOF.
func (me *Bitstream) EOF() bool {
	return me.end
}

// SeekBits will seek the read or write bit position to the absolute bit position given.
func (me *Bitstream) SeekBits(pos int64) {
	if !me.CanSeek() {
		me.seterror(E_SEEK_FAILED)
		return
	}

	// reset end
	me.end = false
	me.seterror(E_NONE)

	// convert the bit position to byte position
	byteoffset := pos >> int64(BSHIFT)

	if me.iotype == BS_INPUT {
		// to seek on input, we'll reload the buffer at new stream position

		noff, err := me.seeker.Seek(byteoffset, io.SeekStart)
		if err != nil || noff != byteoffset {
			me.seterror(E_SEEK_FAILED)
			return
		}

		// clear the buffer
		fill_slice(me.buf[:], 0)

		var l int
		l, err = me.reader.Read(me.buf)

		// check for end of data
		if l == 0 {
			me.end = true
			me.seterror(E_END_OF_DATA)
			me.cur_bit = uint32(pos) & 7
			return
		} else if err != nil {
			me.end = true
			me.seterror(E_READ_FAILED)
			return
		} else if l < int(me.buf_len) {
			me.end = true
			me.buf_len = uint32(l)
			me.seterror(E_END_OF_DATA)
		}

		me.cur_bit = uint32(pos & 7)
	} else {
		// flush and then seek
		me.Flushbits()

		// clear the buffer
		fill_slice(me.buf[:], 0)

		_, err := me.seeker.Seek(byteoffset, io.SeekStart)
		if err != nil {
			me.seterror(E_SEEK_FAILED)
			return
		}

		me.cur_bit = uint32(pos & 7)
	}
}

// NextCode searches for a specified code (input); returns number of bits skipped, excluding the code.
//	If alen > 0, then output bits up to the specified alen-bit boundary (output); returns number of bits written
//	The code is represented using n bits at alen-bit boundary.
func (me *Bitstream) NextCode(code uint64, num_bits uint32, align_length uint32) (uint64, error) {
	var retv uint64 = 0

	if me.iotype == BS_INPUT {
		if align_length == 0 {
			for {
				// for code != me.NextBitsUnsignedBig(num_bits) {
				ncode, err := me.NextBitsUnsignedBig(num_bits)
				if ncode == code {
					break
				} else if err != nil {
					return retv, err
				}
				retv += 1
				me.Skipbits(1)
			}
		} else {
			retv += me.Align(uint64(align_length))
			for {
				// for code != me.NextBitsUnsignedBig(num_bits) {
				ncode, err := me.NextBitsUnsignedBig(num_bits)
				if ncode == code {
					break
				} else if err != nil {
					return retv, err
				}
				retv += uint64(align_length)
				me.Skipbits(align_length)
			}
		}
	} else {
		retv = me.Align(uint64(align_length))
	}
	return retv, nil
}
