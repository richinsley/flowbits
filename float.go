package flobits

import (
	"unsafe"
)

// PutFloat32Big writes the given float32 to the bitstream in Big Endian format.
func (me *Bitstream) PutFloat32Big(value float32) {
	fp := *(*uint32)(unsafe.Pointer(&value))
	me.PutBitsUnsignedBig(uint64(fp), 32)
}

// PutFloat32Little writes the given float32 to the bitstream in Little Endian format.
func (me *Bitstream) PutFloat32Little(value float32) {
	fp := *(*uint32)(unsafe.Pointer(&value))
	me.PutBitsUnsignedLittle(uint64(fp), 32)
}

// GetFloat32Big reads a float32 value from the bitstream in Big Endian format and advances the bit read position.
func (me *Bitstream) GetFloat32Big() (float32, error) {
	lbits, err := me.GetBitsUnsignedBig(32)
	var nb uint32 = uint32(lbits)
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp, err
}

// GetFloat32Little reads a float32 value from the bitstream in Little Endian format and advances the bit read position.
func (me *Bitstream) GetFloat32Little() (float32, error) {
	lbits, err := me.GetBitsUnsignedLittle(32)
	bits := uint32(lbits)
	return *(*float32)(unsafe.Pointer(&bits)), err
}

// NextFloat32Big reads a float32 value from the bitstream in Big Endian format, but does not advance the bit read position.
func (me *Bitstream) NextFloat32Big() (float32, error) {
	lbits, err := me.NextBitsUnsignedBig(32)
	var nb uint32 = uint32(lbits)
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp, err
}

// NextFloat32Little reads a float32 value from the bitstream in Little Endian format, but does not advance the bit read position.
func (me *Bitstream) NextFloat32Little() (float32, error) {
	b, err := me.NextBitsUnsignedLittle(32)
	bits := uint32(b)
	return *(*float32)(unsafe.Pointer(&bits)), err
}

// PutFloat64Big writes the given float64 to the bitstream in Big Endian format.
func (me *Bitstream) PutFloat64Big(value float64) {
	fp := *(*uint64)(unsafe.Pointer(&value))
	me.PutBitsUnsignedBig(uint64(fp), 64)
}

// PutFloat64Little writes the given float64 to the bitstream in Little Endian format.
func (me *Bitstream) PutFloat64Little(value float64) {
	fp := *(*uint64)(unsafe.Pointer(&value))
	me.PutBitsUnsignedLittle(uint64(fp), 64)
}

// GetFloat64Big reads a float64 value from the bitstream in Big Endian format and advances the bit read position.
func (me *Bitstream) GetFloat64Big() (float64, error) {
	nb, err := me.GetBitsUnsignedBig(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

// GetFloat64Little reads a float64 value from the bitstream in Little Endian format and advances the bit read position.
func (me *Bitstream) GetFloat64Little() (float64, error) {
	nb, err := me.GetBitsUnsignedLittle(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

// NextFloat64Big reads a float64 value from the bitstream in Big Endian format, but does not advance the bit read position.
func (me *Bitstream) NextFloat64Big() (float64, error) {
	nb, err := me.NextBitsUnsignedBig(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

// NextFloat64Little reads a float64 value from the bitstream in Little Endian format, but does not advance the bit read position.
func (me *Bitstream) NextFloat64Little() (float64, error) {
	b, err := me.NextBitsUnsignedLittle(64)
	bits := uint64(b)
	fp := *(*float64)(unsafe.Pointer(&bits))
	return fp, err
}
