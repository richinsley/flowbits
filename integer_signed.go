package flobits

import (
	"unsafe"
)

func (me *Flobitsstream) NextBitsSignedBig(n uint32) int64 {
	var x uint64 = me.NextBitsUnsignedBig(n)
	if (n > 1) && (x&smask[n]) != 0 {
		x |= cmask[n]
	}
	return *(*int64)(unsafe.Pointer(&x))
}

func (me *Flobitsstream) GetBitsSignedBig(n uint32) int64 {
	var x uint64 = me.GetBitsUnsignedBig(n)
	if n > 1 && (x&smask[n]) != 0 {
		x = x | cmask[n]
	}
	return *(*int64)(unsafe.Pointer(&x))
}

func (me *Flobitsstream) PutBitsSignedBig(value int64, n uint32) int64 {
	up := *(*uint64)(unsafe.Pointer(&value))
	retv := me.PutBitsUnsignedBig(up, n)
	return *(*int64)(unsafe.Pointer(&retv))
}

// returns 'n' bits as unsigned int with sign extension
// does not advance bit pointer (sign extension only if n>1)
func (me *Flobitsstream) NextBitsSignedLittle(n uint32) int64 {
	x := me.NextBitsUnsignedLittle(n)
	if n > 1 && (x&smask[n] != 0) {
		return int64(x | cmask[n])
	}
	return int64(x)
}

func (me *Flobitsstream) GetBitsSignedLittle(n uint32) int64 {
	var x uint64 = me.GetBitsUnsignedLittle(n)
	if n > 1 && (x&smask[n]) != 0 {
		return int64(x | cmask[n])
	}
	return int64(x)
}

func (me *Flobitsstream) PutBitsSignedLittle(value int64, n uint32) int64 {
	up := *(*uint64)(unsafe.Pointer(&value))
	retv := me.PutBitsUnsignedLittle(up, n)
	return *(*int64)(unsafe.Pointer(&retv))
}
