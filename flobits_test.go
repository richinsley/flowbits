package flowbits

import (
	"bytes"
	"fmt"
	"math/bits"
	"math/rand"
	"testing"

	"github.com/richinsley/purtybits"
)

const (
	INTERNAL_BUFFER_LENGTH int = 1024
)

func TestBigEndianTypeHelpers(t *testing.T) {

	// some arbitrary int values
	var u8 uint8 = 127
	var i8 int8 = -127
	var u16 uint16 = 10000
	var i16 int16 = -10000
	var u32 uint32 = 4096
	var i32 int32 = -4096
	var u64 uint64 = 99999999
	var i64 int64 = -99999999

	w1 := bytes.Buffer{}
	foow1 := NewBitstreamEncoder(&w1, INTERNAL_BUFFER_LENGTH)
	foow1.PutWithBitCountUint8(u8, 8)
	foow1.PutWithBitCountInt8(i8, 8)
	foow1.PutWithBitCountBigUint16(u16, 16)
	foow1.PutWithBitCountBigInt16(i16, 16)
	foow1.PutWithBitCountBigUint32(u32, 32)
	foow1.PutWithBitCountBigInt32(i32, 32)
	foow1.PutWithBitCountBigUint64(u64, 64)
	foow1.PutWithBitCountBigInt64(i64, 64)
	foow1.Flushbits()

	w2 := bytes.Buffer{}
	foow2 := NewBitstreamEncoder(&w2, INTERNAL_BUFFER_LENGTH)
	foow2.PutUint8(u8)
	foow2.PutInt8(i8)
	foow2.PutBigUint16(u16)
	foow2.PutBigInt16(i16)
	foow2.PutBigUint32(u32)
	foow2.PutBigInt32(i32)
	foow2.PutBigUint64(u64)
	foow2.PutBigInt64(i64)
	foow2.Flushbits()

	// compare the 2 buffers
	if !unsigned_buffersEqual(w1.Bytes(), w2.Bytes()) {
		t.Errorf("Both write buffers should be equal")
	}

	r := bytes.NewReader(w1.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	if v, _ := foor.GetUint8(); v != u8 {
		t.Errorf("%d and %d should be equal", v, u8)
	}
	if v, _ := foor.GetInt8(); v != i8 {
		t.Errorf("%d and %d should be equal", v, i8)
	}
	if v, _ := foor.GetBigUint16(); v != u16 {
		t.Errorf("%d and %d should be equal", v, u16)
	}
	if v, _ := foor.GetBigInt16(); v != i16 {
		t.Errorf("%d and %d should be equal", v, i16)
	}
	if v, _ := foor.GetBigUint32(); v != u32 {
		t.Errorf("%d and %d should be equal", v, u16)
	}
	if v, _ := foor.GetBigInt32(); v != i32 {
		t.Errorf("%d and %d should be equal", v, i32)
	}
	if v, _ := foor.GetBigUint64(); v != u64 {
		t.Errorf("%d and %d should be equal", v, u64)
	}
	if v, _ := foor.GetBigInt64(); v != i64 {
		t.Errorf("%d and %d should be equal", v, i64)
	}

	w3 := bytes.Buffer{}
	foow3 := NewBitstreamEncoder(&w3, INTERNAL_BUFFER_LENGTH)
	foow3.PutWithBitCountUint8(u8, 8)
	foow3.PutWithBitCountInt8(i8, 8)
	foow3.PutWithBitCountLittleUint16(u16, 16)
	foow3.PutWithBitCountLittleInt16(i16, 16)
	foow3.PutWithBitCountLittleUint32(u32, 32)
	foow3.PutWithBitCountLittleInt32(i32, 32)
	foow3.PutWithBitCountLittleUint64(u64, 64)
	foow3.PutWithBitCountLittleInt64(i64, 64)
	foow3.Flushbits()

	w4 := bytes.Buffer{}
	foow4 := NewBitstreamEncoder(&w4, INTERNAL_BUFFER_LENGTH)
	foow4.PutUint8(u8)
	foow4.PutInt8(i8)
	foow4.PutLittleUint16(u16)
	foow4.PutLittleInt16(i16)
	foow4.PutLittleUint32(u32)
	foow4.PutLittleInt32(i32)
	foow4.PutLittleUint64(u64)
	foow4.PutLittleInt64(i64)
	foow4.Flushbits()

	// compare the 2 buffers
	if !unsigned_buffersEqual(w3.Bytes(), w4.Bytes()) {
		t.Errorf("Both write buffers should be equal")
	}

	r2 := bytes.NewReader(w3.Bytes())
	foor2 := NewBitstreamDecoder(r2, INTERNAL_BUFFER_LENGTH)

	if v, _ := foor2.GetUint8(); v != u8 {
		t.Errorf("%d and %d should be equal", v, u8)
	}
	if v, _ := foor2.GetInt8(); v != i8 {
		t.Errorf("%d and %d should be equal", v, i8)
	}
	if v, _ := foor2.GetLittleUint16(); v != u16 {
		t.Errorf("%d and %d should be equal", v, u16)
	}
	if v, _ := foor2.GetLittleInt16(); v != i16 {
		t.Errorf("%d and %d should be equal", v, i16)
	}
	if v, _ := foor2.GetLittleUint32(); v != u32 {
		t.Errorf("%d and %d should be equal", v, u16)
	}
	if v, _ := foor2.GetLittleInt32(); v != i32 {
		t.Errorf("%d and %d should be equal", v, i32)
	}
	if v, _ := foor2.GetLittleUint64(); v != u64 {
		t.Errorf("%d and %d should be equal", v, u64)
	}
	if v, _ := foor2.GetLittleInt64(); v != i64 {
		t.Errorf("%d and %d should be equal", v, i64)
	}
}

func TestBool(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)

	for i := 0; i < 128; i++ {
		foow.PutBool(i%2 == 0)
	}

	// flush the contents of the flowbits buffer
	foow.Flushbits()

	// show our work
	purty := purtybits.NewPurtyBits(4, purtybits.HexCodeGroupToRight)
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	// peek the first bool
	b, _ := foor.NextBool()
	if !b {
		t.Errorf("First bool should be true")
	}

	for i := 0; i < 128; i++ {
		b, _ := foor.GetBool()
		if (i%2 == 0) != b {
			t.Errorf("incorrect bool")
		}
	}
}

func TestSeek(t *testing.T) {
	var bufferout [16]uint8 = [16]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)
	foow.PutBuffer(bufferout[:])

	// flush the contents of the flowbits buffer
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	foor.SeekBits(8 * 3)

	lbits, _ := foor.GetBitsUnsignedBig(8)
	var u = uint8(lbits)

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

	// the above floats and doubles generated on an ARM system with the C++ version of flowbits
	var bufferin []uint8 = []uint8{65, 157, 111, 52, 85, 95, 173, 64, 68, 121, 250, 55, 65, 128,
		96, 80, 72, 252, 214, 234, 68, 122, 26, 55, 64, 173, 95, 85, 52, 111, 157, 65, 55, 250,
		121, 68, 234, 214, 252, 72, 80, 96, 128, 65, 55, 26, 122, 68}

	rbuffer := bytes.NewBuffer(bufferin)
	foor := NewBitstreamDecoder(rbuffer, INTERNAL_BUFFER_LENGTH)

	// test big endian floats and doubles
	id1, _ := foor.NextFloat64Big()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	id1, _ = foor.GetFloat64Big()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	if1, _ := foor.NextFloat32Big()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	if1, _ = foor.GetFloat32Big()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	id2, _ := foor.NextFloat64Big()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	id2, _ = foor.GetFloat64Big()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	if2, _ := foor.NextFloat32Big()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	if2, _ = foor.GetFloat32Big()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	// test little endian floats and doubles
	id1, _ = foor.NextFloat64Little()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	id1, _ = foor.GetFloat64Little()
	if id1 != d1 {
		t.Errorf("id1 = %f, want %f", id1, d1)
	}

	if1, _ = foor.NextFloat32Little()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	if1, _ = foor.GetFloat32Little()
	if if1 != f1 {
		t.Errorf("if1 = %f, want %f", if1, f1)
	}

	id2, _ = foor.NextFloat64Little()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	id2, _ = foor.GetFloat64Little()
	if id2 != d2 {
		t.Errorf("id2 = %f, want %f", id2, d2)
	}

	if2, _ = foor.NextFloat32Little()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}

	if2, _ = foor.GetFloat32Little()
	if if2 != f2 {
		t.Errorf("if2 = %f, want %f", if2, f2)
	}
}

func TestReaderInit(t *testing.T) {
	var bufferin []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0}
	rbuffer := bytes.NewBuffer(bufferin)
	foor := NewBitstreamDecoder(rbuffer, INTERNAL_BUFFER_LENGTH)
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
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)

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

	// flush the contents of the flowbits buffer
	foow.Flushbits()

	// show work
	fmt.Printf("TestCodes buffer length: %d\n", w.Len())

	// purtyfy the bits
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	// NextCode will search for the given code and return the count of
	// bits that were skipped over to land at the start of the code
	code1skipped, _ := foor.NextCode(code1, 32, 0)
	t.Logf("NextCode test skipped %d bits", code1skipped)

	// get the current read bit position
	icode1pos := foor.GetPos()
	if icode1pos != code_one_start_bit {
		t.Errorf("icode1pos = %d, want %d", icode1pos, code_one_start_bit)
	}

	code2skipped, _ := foor.NextCode(code2, 32, 0)
	t.Logf("NextCode test skipped %d bits", code2skipped)
	icode2pos := foor.GetPos()
	if icode2pos != code_two_start_bit {
		t.Errorf("icode2pos = %d, want %d", icode2pos, code_two_start_bit)
	}
}

func TestEndianFloats(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)
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
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	// read in big endian bits
	uin, _ := foor.GetBitsUnsignedBig(64)
	if uin != uout {
		t.Errorf("uin = %d, want %d", uin, uout)
	}

	// PROBING the next uint64 as big endian should yield an incorrect value because it
	// is encoded as little endian
	uin, _ = foor.NextBitsUnsignedBig(64)
	if uin == uout {
		t.Errorf("uin should not be equal to uout")
	}

	// Getting the next uint64 should be equal to uout
	uin, _ = foor.GetBitsUnsignedLittle(64)
	if uin != uout {
		t.Errorf("uin = %d, want %d", uin, uout)
	}
}

func TestLittleIntDelta(t *testing.T) {
	var out uint64 = 1543572285742637646
	// 0b000.10101 01101011 11011110 11010000 11010000 11010011 11100010 01001110
	// 15          6B       DE       D0       D0       D3       E2       4E

	// if we are at a bit 3 offset
	var skip int = 3
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)
	foow.Skipbits(uint32(skip))
	foow.PutBitsUnsignedLittle(out, 61)
	foow.Flushbits()

	// 000 01001.110 11100010 11010011 11010000 11010000 11011110 01101011 10101

	purty := purtybits.NewPurtyBits(4, purtybits.HexCodeGroupToRight)
	purtyrows := purty.BufferToStrings(w.Bytes())
	for _, s := range purtyrows {
		fmt.Println(s)
	}

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)
	foor.Skipbits(3)

	in, _ := foor.GetBitsUnsignedLittle(61)
	if in != out {
		t.Errorf("in = %d, want %d", in, out)
	}
}

func TestFuzz(t *testing.T) {
	var test_array_length int = 10000
	var iterations int = 100

	fmt.Printf("TestFuzz iterating over %d arrays of %d values\n", iterations, test_array_length)
	for u := 0; u < iterations; u++ {
		// create a new random seed for each iteration
		random := rand.New(rand.NewSource(int64(u)))

		// create an array of random uint64
		var buffer_out []uint64 = make([]uint64, test_array_length)
		for i := 0; i < test_array_length; i++ {
			buffer_out[i] = uint64(random.Int63())
		}

		w := bytes.Buffer{}
		foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)
		for i := 0; i < test_array_length; i++ {
			// write out a 6-bit value for the number of bits that uint32 value
			// requires (this can be done with https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/math/bits/bits.go;l=318 )
			bcount := bits.Len64(buffer_out[i])
			foow.PutBitsUnsignedBig(uint64(bcount), 7)
			// write out the actual value 0 bigE for even, littleE for odd
			if i%2 == 0 {
				foow.PutBitsUnsignedBig(uint64(buffer_out[i]), uint32(bcount))
			} else {
				foow.PutBitsUnsignedLittle(uint64(buffer_out[i]), uint32(bcount))
			}
		}

		// we don't know if the buffer will be byte aligned, so byte align it
		skipped_bits, _ := foow.Align(8)
		fmt.Printf("Fuzz pass align skipped %d bits\n", skipped_bits)

		foow.Flushbits()

		// read the array back in
		var buffer_in []uint64 = make([]uint64, test_array_length)
		r := bytes.NewReader(w.Bytes())
		foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

		for i := 0; i < test_array_length; i++ {
			lbits, _ := foor.GetBitsUnsignedBig(7)
			bcount := uint32(lbits)
			if i%2 == 0 {
				lbits, _ = foor.GetBitsUnsignedBig(bcount)
			} else {
				lbits, _ = foor.GetBitsUnsignedLittle(bcount)
			}
			buffer_in[i] = uint64(lbits)
		}

		if !buffers64Equal(buffer_in, buffer_out) {
			t.Errorf("buffer_in and buffer_out must be equal at seed %d", u)
		}
	}
}

func TestInt64(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)

	var uint_out uint64 = 0x1122334455667788
	var tiny_out uint64 = 0b11 // 3
	var tiny_in uint64
	var uint_in uint64

	foow.PutBitsUnsignedBig(uint_out, 64)
	foow.PutBitsUnsignedBig(tiny_out, 3)
	foow.PutBitsUnsignedBig(uint_out, 64)
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	uint_in, _ = foor.GetBitsUnsignedBig(64)
	if uint_in != uint_out {
		t.Errorf("uint_in = %d, want %d", uint_in, uint_out)
	}

	tiny_in, _ = foor.GetBitsUnsignedBig(3)
	if tiny_in != tiny_out {
		t.Errorf("uint_in = %d, want %d", tiny_in, tiny_out)
	}

	uint_in, _ = foor.GetBitsUnsignedBig(64)
	if uint_in != uint_out {
		t.Errorf("uint_in = %d, want %d", uint_in, uint_out)
	}
}

func TestGeneral(t *testing.T) {
	w := bytes.Buffer{}
	foow := NewBitstreamEncoder(&w, INTERNAL_BUFFER_LENGTH)

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
	skippedbits, _ := foow.Align(8)
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

	// flush the contents of the flowbits buffer
	foow.Flushbits()

	r := bytes.NewReader(w.Bytes())
	foor := NewBitstreamDecoder(r, INTERNAL_BUFFER_LENGTH)

	if !foor.CanSeek() {
		t.Error("foor.CanSeek should be true (bytes.Reader implements io.Seeker")
	}

	float1_in, _ = foor.GetFloat32Big()
	if float1_in != float1_out {
		t.Errorf("float1_in = %f, want %f", float1_in, float1_out)
	}

	double1_in, _ = foor.GetFloat64Big()
	if double1_in != double1_out {
		t.Errorf("double1_in = %f, want %f", double1_in, double1_out)
	}

	uint1_in, _ = foor.GetBitsUnsignedBig(3)
	if uint1_in != uint1_out {
		t.Errorf("uint1_in = %d, want %d", uint1_in, uint1_out)
	}

	uint2_in, _ = foor.GetBitsUnsignedBig(5)
	if uint2_in != uint2_out {
		t.Errorf("uint2_in = %d, want %d", uint2_in, uint2_out)
	}

	uint3_in, _ = foor.GetBitsUnsignedBig(60)
	if uint3_in != uint3_out {
		t.Errorf("uint3_in = %d, want %d", uint3_in, uint3_out)
	}

	int1_in, _ = foor.GetBitsSignedBig(59)
	if int1_in != int1_out {
		t.Errorf("int1_in = %d, want %d", int1_in, int1_out)
	}

	foor.GetBuffer(buffer_in, 16)
	if !unsigned_buffersEqual(buffer_in, buffer_out[:]) {
		t.Errorf("buffer_in and buffer_must be equal")
	}

	for i := 0; i < 8; i = i + 1 {
		x, _ := foor.GetBitsSignedBig(3)
		signed_buffer_in[i] = int8(x)
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
		large_64buffer_in[i], _ = foor.GetBitsUnsignedBig(64)
	}

	if !buffers64Equal(large_64buffer_in, large_64buffer_out) {
		t.Errorf("large_64buffer_in and large_64buffer_out must be equal")
	}

	// read the end code
	end_code_in, _ = foor.GetBitsUnsignedBig(64)
	if end_code_in != end_code_out {
		t.Errorf("end_code_in and end_code_out must be equal")
	}
}
