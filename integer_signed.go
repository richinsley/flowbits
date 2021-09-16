package flobits

import (
	"unsafe"
)

// NextBitsSignedBig returns 'n' bits in Big Endian as an int64 with sign extension
// does not advance bit pointer. (sign extension only if n>1)
func (me *Bitstream) NextBitsSignedBig(n uint32) (int64, error) {
	x, err := me.NextBitsUnsignedBig(n)
	if (n > 1) && (x&smask[n]) != 0 {
		x |= cmask[n]
	}
	return *(*int64)(unsafe.Pointer(&x)), err
}

// GetBitsSignedBig returns 'n' bits in Big Endian as an int64 with sign extension
// and advances the bit read position. (sign extension only if n>1)
func (me *Bitstream) GetBitsSignedBig(n uint32) (int64, error) {
	x, err := me.GetBitsUnsignedBig(n)
	if n > 1 && (x&smask[n]) != 0 {
		x = x | cmask[n]
	}
	return *(*int64)(unsafe.Pointer(&x)), err
}

// PutBitsSignedBig writes 'n' bits as a signed int in Big Endian. (sign extension only if n>1)
func (me *Bitstream) PutBitsSignedBig(value int64, n uint32) int64 {
	up := *(*uint64)(unsafe.Pointer(&value))
	retv := me.PutBitsUnsignedBig(up, n)
	return *(*int64)(unsafe.Pointer(&retv))
}

// NextBitsSignedLittle returns 'n' bits in Little Endian as an int64 with sign extension
// does not advance bit pointer. (sign extension only if n>1)
func (me *Bitstream) NextBitsSignedLittle(n uint32) (int64, error) {
	x, err := me.NextBitsUnsignedLittle(n)
	if n > 1 && (x&smask[n] != 0) {
		return int64(x | cmask[n]), err
	}
	return int64(x), err
}

// GetBitsSignedLittle returns 'n' bits in Little Endian as an int64 with sign extension
// and advances the bit read position. (sign extension only if n>1)
func (me *Bitstream) GetBitsSignedLittle(n uint32) (int64, error) {
	x, err := me.GetBitsUnsignedLittle(n)
	if n > 1 && (x&smask[n]) != 0 {
		return int64(x | cmask[n]), err
	}
	return int64(x), err
}

// PutBitsSignedLittle writes 'n' bits as a signed int in Little Endian. (sign extension only if n>1)
func (me *Bitstream) PutBitsSignedLittle(value int64, n uint32) int64 {
	up := *(*uint64)(unsafe.Pointer(&value))
	retv := me.PutBitsUnsignedLittle(up, n)
	return *(*int64)(unsafe.Pointer(&retv))
}
