package flobits

import (
	"bytes"
	"fmt"
	"math/bits"
	"math/rand"
	"testing"

	"github.com/richinsley/purtybits"
)

const (
	INTERNAL_BUFFER_LENGTH int = 32
)

func TestSeek(t *testing.T) {
	var bufferout [16]uint8 = [16]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	w := bytes.Buffer{}
	foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)
	foow.PutBuffer(bufferout[:])

	// flush the contents of the flobits buffer
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

	foor.SeekBits(8 * 3)

	var u = uint8(foor.GetBitsUnsignedBig(8))

	if u != 3 {
		t.Errorf("u = %d, want 3", u)
	}

	purty := purtybits.NewPurtyBits(4, purtybits.HexCodeGroupToRight)
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}
}

func TestFloats(t *testing.T) {
	// a set of arbitrary floats
	var d1 float64 = 123456789.34343433
	var f1 float32 = 999.9096069335938
	var d2 float64 = 34343433.123456789
	var f2 float32 = f1 + 0.5

	// the above floats and doubles generated on an ARM system with the C++ version of flobits
	var bufferin []uint8 = []uint8{65, 157, 111, 52, 85, 95, 173, 64, 68, 121, 250, 55, 65, 128,
		96, 80, 72, 252, 214, 234, 68, 122, 26, 55, 64, 173, 95, 85, 52, 111, 157, 65, 55, 250,
		121, 68, 234, 214, 252, 72, 80, 96, 128, 65, 55, 26, 122, 68}

	rbuffer := bytes.NewBuffer(bufferin)
	foor := NewFlobitsDecoder(rbuffer, INTERNAL_BUFFER_LENGTH)

	// test big endian floats and doubles
	id1 := foor.NextFloat64Big()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	id1 = foor.GetFloat64Big()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	if1 := foor.NextFloat32Big()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	if1 = foor.GetFloat32Big()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	id2 := foor.NextFloat64Big()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	id2 = foor.GetFloat64Big()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	if2 := foor.NextFloat32Big()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	if2 = foor.GetFloat32Big()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	// test little endian floats and doubles
	id1 = foor.NextFloat64Little()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	id1 = foor.GetFloat64Little()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	if1 = foor.NextFloat32Little()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	if1 = foor.GetFloat32Little()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	id2 = foor.NextFloat64Little()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	id2 = foor.GetFloat64Little()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	if2 = foor.NextFloat32Little()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	if2 = foor.GetFloat32Little()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}
}

func TestReaderInit(t *testing.T) {
	var bufferin []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0}
	rbuffer := bytes.NewBuffer(bufferin)
	foor := NewFlobitsDecoder(rbuffer, INTERNAL_BUFFER_LENGTH)
	if foor.cur_bit != 0 {
		t.Errorf("foor.cur_bit = %d, want 0", foor.cur_bit)
	}

	if foor.buf_len != 8 {
		t.Errorf("foor.buf_len = %d, want 8", foor.cur_bit)
	}
}
func TestCodes(t *testing.T) {
	// the codes we'll search for
	var code1 uint64 = 0x01020304
	var code2 uint64 = 0x01020305

	// create a purtybits object to show the absolute placement of the codes in the stream
	purty := purtybits.NewPurtyBits(8, purtybits.HexCodeGroupToRight)

	w := bytes.Buffer{}
	foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)

	// the default internal buffer size is 64 bytes.
	// write out 63 bytes of '0xff' then write out the code we'll search for
	for i := 0; i < 63; i++ {
		foow.PutBitsUnsignedBig(255, 8)
	}

	// annotate the position in bits and write out code 1
	code_one_start_bit := foow.GetPos()
	foow.PutBitsUnsignedBig(code1, 32)
	code_one_end_bit := int(foow.GetPos() - 1)
	purty.ColorBitRange(int(code_one_start_bit), code_one_end_bit, purtybits.PurtyBitColorWhiteOnRed("1"), purtybits.PurtyBitColorBlackOnRed("0"))

	// pad 256 bytes of '0xff' then write out code 2
	for i := 0; i < 256; i++ {
		foow.PutBitsUnsignedBig(255, 8)
	}

	// annotate the position in bits and write out code 1
	code_two_start_bit := foow.GetPos()
	foow.PutBitsUnsignedBig(code2, 32)
	code_two_end_bit := int(foow.GetPos() - 1)
	purty.ColorBitRange(int(code_two_start_bit), code_two_end_bit, purtybits.PurtyBitColorWhiteOnRed("1"), purtybits.PurtyBitColorBlackOnRed("0"))

	// flush the contents of the flobits buffer
	foow.Flushbits()

	// show work
	fmt.Printf("TestCodes buffer length: %d\n", w.Len())

	// purtyfy the bits
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}

	r := bytes.NewReader(w.Bytes())
	foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

	// NextCode will search for the given code and return the count of
	// bits that were skipped over to land at the start of the code
	code1skipped := foor.NextCode(code1, 32, 0)
	t.Logf("NextCode test skipped %d bits", code1skipped)

	// get the current read bit position
	icode1pos := foor.GetPos()
	if icode1pos != code_one_start_bit {
		t.Errorf("icode1pos = %d, want %d", icode1pos, code_one_start_bit)
	}

	code2skipped := foor.NextCode(code2, 32, 0)
	t.Logf("NextCode test skipped %d bits", code2skipped)
	icode2pos := foor.GetPos()
	if icode2pos != code_two_start_bit {
		t.Errorf("icode2pos = %d, want %d", icode2pos, code_two_start_bit)
	}
}

func TestEndianFloats(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)
	var uout uint64 = 0x00aabbccddeeff11
	foow.PutBitsUnsignedBig(uout, 64)
	foow.PutBitsUnsignedLittle(uout, 64)
	foow.Flushbits()

	// purtyfy the bits
	purty := purtybits.NewPurtyBits(4, purtybits.HexCodeGroupToRight)
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}

	r := bytes.NewReader(w.Bytes())
	foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

	// read in big endian bits
	uin := foor.GetBitsUnsignedBig(64)
	if uin != uout {
		t.Errorf("uin = %d, want %d", uin, uout)
	}

	// PROBING the next uint64 as big endian should yield an incorrect value because it
	// is encoded as little endian
	uin = foor.NextBitsUnsignedBig(64)
	if uin == uout {
		t.Errorf("uin should not be equal to uout")
	}

	// Getting the next uint64 should be equal to uout
	uin = foor.GetBitsUnsignedLittle(64)
	if uin != uout {
		t.Errorf("uin = %d, want %d", uin, uout)
	}
}

func TestMonkey(t *testing.T) {
	var test_array_length int = 100000
	var iterations int = 100

	fmt.Printf("TestMonkey iterating over %d arrays of %d values\n", iterations, test_array_length)
	for u := 0; u < iterations; u++ {
		// create a new random seed for each iteration
		random := rand.New(rand.NewSource(int64(u)))

		// create an array of random uint32 with a range of 0 - 32000
		var buffer_out []uint32 = make([]uint32, test_array_length)
		for i := 0; i < test_array_length; i++ {
			buffer_out[i] = uint32(random.Int31())
		}

		w := bytes.Buffer{}
		foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)
		for i := 0; i < test_array_length; i++ {
			// write out a 6-bit value for the number of bits that uint32 value
			// requires (this can be done with https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/math/bits/bits.go;l=318 )
			bcount := bits.Len32(buffer_out[i])
			foow.PutBitsUnsignedBig(uint64(bcount), 6)
			// write out the actual value
			foow.PutBitsUnsignedBig(uint64(buffer_out[i]), uint32(bcount))
		}

		// we don't know if the buffer will be byte aligned, so byte align it
		foow.Align(8)
		foow.Flushbits()

		// read the array back in
		var buffer_in []uint32 = make([]uint32, test_array_length)
		r := bytes.NewReader(w.Bytes())
		foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

		for i := 0; i < test_array_length; i++ {
			bcount := uint32(foor.GetBitsUnsignedBig(6))
			buffer_in[i] = uint32(foor.GetBitsUnsignedBig(bcount))
		}

		if !buffers32Equal(buffer_in, buffer_out) {
			t.Errorf("buffer_in and buffer_out must be equal at seed %d", u)
		}
	}
}

func TestInt64(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)

	var uint_out uint64 = 0x1122334455667788
	var tiny_code uint64 = 3
	var uint_in uint64

	foow.PutBitsUnsignedBig(uint_out, 64)
	foow.PutBitsUnsignedBig(tiny_code, 3)
	foow.PutBitsUnsignedBig(uint_out, 64)
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

	uint_in = foor.GetBitsUnsignedBig(64)
	if uint_in != uint_out {
		t.Errorf("uint_in = %d, want %d", uint_in, uint_out)
	}
	foor.Skipbits(3)
	uint_in = foor.GetBitsUnsignedBig(64)
	if uint_in != uint_out {
		t.Errorf("uint_in = %d, want %d", uint_in, uint_out)
	}
}

func TestGeneral(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewFlobitsEncoder(&w, INTERNAL_BUFFER_LENGTH)

	if foow.CanSeek() {
		t.Error("foow.CanSeek should be false (bytes.Buffer does not implement io.Seeker")
	}

	// an arbitrary unsigned buffer of 16 bytes
	var buffer_out [16]uint8 = [16]uint8{0xfa /* 0xfa == 250 == b11111010 */, 1, 2, 3, 4, 5, 6, 7, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7}
	var buffer_in []uint8 = make([]uint8, 16)

	// an arbitrary signed buffer of 8 bytes with a range of -4..3
	var signed_buffer_out [8]int8 = [8]int8{-4, -3, -2, -1, 0, 1, 2, 3}
	var signed_buffer_in []int8 = make([]int8, 8)

	// a large buffer of unsigned bytes with each value = v[i & 0xff]
	large_buffer_len := 2096
	var large_buffer_out []uint8 = make([]uint8, large_buffer_len)
	var large_buffer_in []uint8 = make([]uint8, large_buffer_len)
	for i := 0; i < large_buffer_len; i++ {
		large_buffer_out[i] = uint8(i)
	}

	// a large buffer of uint64 values
	large_64_len := 100
	var large_64buffer_out []uint64 = make([]uint64, large_64_len)
	var large_64buffer_in []uint64 = make([]uint64, large_64_len)
	for i := 0; i < large_64_len; i++ {
		large_64buffer_out[i] = uint64(i)
	}

	var float1_out float32 = 9.6969
	var double1_out float64 = 3.14
	var uint1_out uint64 = 2
	var uint2_out uint64 = 19
	var uint3_out uint64 = 1970
	var int1_out int64 = -1970
	var end_code_out uint64 = 0x1122334455667788

	var float1_in float32
	var double1_in float64
	var uint1_in uint64
	var uint2_in uint64
	var uint3_in uint64
	var int1_in int64
	var end_code_in uint64

	// encode a float32 with 32 bits
	foow.PutFloat32Big(float1_out)

	// encode a float64 with 64 bits
	foow.PutFloat64Big(double1_out)

	// encode a uint with 3 bits
	foow.PutBitsUnsignedBig(uint1_out, 3)

	// encode a uint with 5 bits
	foow.PutBitsUnsignedBig(uint2_out, 5)

	// encode a uint with 60 bits
	foow.PutBitsUnsignedBig(uint3_out, 60)

	// encode a signed int with 59 bits
	foow.PutBitsSignedBig(int1_out, 59)

	// encode an entire unsigned uint8 buffer
	foow.PutBuffer(buffer_out[:])

	// encode an int8 buffer using 3 bits per byte
	// a 3-bit two's complement signed int can represent a range of -4 to 3
	for i := 0; i < 8; i = i + 1 {
		foow.PutBitsSignedBig(int64(signed_buffer_out[i]), 3)
	}

	// align back to 8 bit boundry and re-write 16 byte buffer
	skippedbits := foow.Align(8)
	if skippedbits != 1 {
		t.Errorf("skippedbits = %d, want 1", skippedbits)
	}

	foow.PutBuffer(buffer_out[:])

	// put the uint8 large buffer
	foow.PutBuffer(large_buffer_out[:])

	// ------------------------------------------------------------------------------
	// put 3 arbitrary bits to make the buffer non-aligned and put a buffer of uint64
	// to test the delta offset in Putbits
	foow.PutBitsUnsignedBig(3, 3)

	// put the uint64 buffer
	for i := 0; i < large_64_len; i++ {
		foow.PutBitsUnsignedBig(large_64buffer_out[i], 64)
	}

	// write out the end code
	foow.PutBitsUnsignedBig(end_code_out, 64)

	// ------------------------------------------------------------------------------

	// flush the contents of the flobits buffer
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewFlobitsDecoder(r, INTERNAL_BUFFER_LENGTH)

	if !foor.CanSeek() {
		t.Error("foor.CanSeek should be true (bytes.Reader implements io.Seeker")
	}

	float1_in = foor.GetFloat32Big()
	if float1_in != float1_out {
		t.Errorf("float1_in = %f, want %f", float1_in, float1_out)
	}

	double1_in = foor.GetFloat64Big()
	if double1_in != double1_out {
		t.Errorf("double1_in = %f, want %f", double1_in, double1_out)
	}

	uint1_in = foor.GetBitsUnsignedBig(3)
	if uint1_in != uint1_out {
		t.Errorf("uint1_in = %d, want %d", uint1_in, uint1_out)
	}

	uint2_in = foor.GetBitsUnsignedBig(5)
	if uint2_in != uint2_out {
		t.Errorf("uint2_in = %d, want %d", uint2_in, uint2_out)
	}

	uint3_in = foor.GetBitsUnsignedBig(60)
	if uint3_in != uint3_out {
		t.Errorf("uint3_in = %d, want %d", uint3_in, uint3_out)
	}

	int1_in = foor.GetBitsSignedBig(59)
	if int1_in != int1_out {
		t.Errorf("int1_in = %d, want %d", int1_in, int1_out)
	}

	foor.GetBuffer(buffer_in, 16)
	if !unsigned_buffersEqual(buffer_in, buffer_out[:]) {
		t.Errorf("buffer_in and buffer_must be equal")
	}

	for i := 0; i < 8; i = i + 1 {
		signed_buffer_in[i] = int8(foor.GetBitsSignedBig(3))
	}
	if !signed_buffersEqual(signed_buffer_in, signed_buffer_out[:]) {
		t.Errorf("signed_buffer_in and signed_buffer_out be equal")
	}

	// align back to 8 bit boundry and read 16 byte buffer
	foor.Align(8)
	foor.GetBuffer(buffer_in, 16)
	if !unsigned_buffersEqual(buffer_in, buffer_out[:]) {
		t.Errorf("buffer_in and buffer_must be equal")
	}

	// read large buffer
	foor.GetBuffer(large_buffer_in, uint64(large_buffer_len))
	if !unsigned_buffersEqual(large_buffer_in, large_buffer_out) {
		t.Errorf("buffer_in and buffer_must be equal")
	}

	// read 3 bits plus an array of uint64
	foor.GetBitsUnsignedBig(3)
	for i := 0; i < large_64_len; i++ {
		large_64buffer_in[i] = foor.GetBitsUnsignedBig(64)
	}

	if !buffers64Equal(large_64buffer_in, large_64buffer_out) {
		t.Errorf("large_64buffer_in and large_64buffer_out must be equal")
	}

	// read the end code
	end_code_in = foor.GetBitsUnsignedBig(64)
	if end_code_in != end_code_out {
		t.Errorf("end_code_in and end_code_out must be equal")
	}
}
