package flobits

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReaderOverSimple(t *testing.T) {
	// create a 16 byte buffer
	testslice := [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	testbuffer := bytes.NewBuffer(testslice[:])
	foor := NewBitstreamDecoder(testbuffer, 16)

	// read two uint64s
	foor.GetBitsUnsignedBig(64)
	foor.GetBitsUnsignedBig(64)

	fmt.Println(foor.availableBufferBits())

	// the bitstream should be EOF, so the following NextBitsUnsignedBig must yeild an EOF error
	_, err := foor.NextFloat32Little()
	if err == nil {
		t.Errorf("NextBitsUnsignedBig must return non-nil error")
	}

	available_bits := foor.availableBufferBits()
	if available_bits != 0 {
		t.Errorf("available_bits should be 0, got %d", available_bits)
	}

	// write 4 more bytes to the testbuffer then refill our read buffer
	testbuffer.Write([]uint8{0, 0, 0, 0})
	foor.FillReadBuffer()

	// we should now report 32 bits are available in the read buffer
	available_bits = foor.availableBufferBits()
	if available_bits != 32 {
		t.Errorf("available_bits should be 32, got %d", available_bits)
	}

	// we should now be able to read in those 4 bytes
	_, err = foor.NextFloat32Little()
	if err != nil {
		t.Errorf("NextBitsUnsignedBig must return non-nil error")
	}
}
