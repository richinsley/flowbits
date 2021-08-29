package flobits

import (
	"unsafe"
)

func (me *Flobitsstream) PutFloat32Big(value float32) {
	fp := *(*uint32)(unsafe.Pointer(&value))
	me.PutBitsUnsignedBig(uint64(fp), 32)
}

func (me *Flobitsstream) PutFloat32Little(value float32) {
	fp := *(*uint32)(unsafe.Pointer(&value))
	me.PutBitsUnsignedLittle(uint64(fp), 32)
}

func (me *Flobitsstream) GetFloat32Big() float32 {
	var nb uint32 = uint32(me.GetBitsUnsignedBig(32))
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp
}

func (me *Flobitsstream) GetFloat32Little() float32 {
	bits := uint32(me.GetBitsUnsignedLittle(32))
	return *(*float32)(unsafe.Pointer(&bits))
}

func (me *Flobitsstream) NextFloat32Big() float32 {
	var nb uint32 = uint32(me.NextBitsUnsignedBig(32))
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp
}

func (me *Flobitsstream) PutFloat64Big(value float64) {
	fp := *(*uint64)(unsafe.Pointer(&value))
	me.PutBitsUnsignedBig(uint64(fp), 64)
}

func (me *Flobitsstream) GetFloat64Big() float64 {
	var nb uint64 = uint64(me.GetBitsUnsignedBig(64))
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp
}

func (me *Flobitsstream) NextFloat64Big() float64 {
	var nb uint64 = uint64(me.NextBitsUnsignedBig(64))
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp
}

func (me *Flobitsstream) NextFloat32Little() float32 {
	bits := uint32(me.NextBitsUnsignedLittle(32))
	return *(*float32)(unsafe.Pointer(&bits))
}

func (me *Flobitsstream) GetFloat64Little() float64 {
	var nb uint64 = uint64(me.GetBitsUnsignedLittle(64))
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp
}

func (me *Flobitsstream) NextFloat64Little() float64 {
	var nb uint64 = uint64(me.NextBitsUnsignedLittle(64))
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp
}
