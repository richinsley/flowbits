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

func (me *Flobitsstream) GetFloat32Big() (float32, error) {
	lbits, err := me.GetBitsUnsignedBig(32)
	var nb uint32 = uint32(lbits)
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp, err
}

func (me *Flobitsstream) GetFloat32Little() (float32, error) {
	lbits, err := me.GetBitsUnsignedLittle(32)
	bits := uint32(lbits)
	return *(*float32)(unsafe.Pointer(&bits)), err
}

func (me *Flobitsstream) NextFloat32Big() (float32, error) {
	lbits, err := me.NextBitsUnsignedBig(32)
	var nb uint32 = uint32(lbits)
	fp := *(*float32)(unsafe.Pointer(&nb))
	return fp, err
}

func (me *Flobitsstream) PutFloat64Big(value float64) {
	fp := *(*uint64)(unsafe.Pointer(&value))
	me.PutBitsUnsignedBig(uint64(fp), 64)
}

func (me *Flobitsstream) GetFloat64Big() (float64, error) {
	nb, err := me.GetBitsUnsignedBig(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

func (me *Flobitsstream) NextFloat64Big() (float64, error) {
	nb, err := me.NextBitsUnsignedBig(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

func (me *Flobitsstream) NextFloat32Little() (float32, error) {
	b, err := me.NextBitsUnsignedLittle(32)
	bits := uint32(b)
	return *(*float32)(unsafe.Pointer(&bits)), err
}

func (me *Flobitsstream) GetFloat64Little() (float64, error) {
	nb, err := me.GetBitsUnsignedLittle(64)
	fp := *(*float64)(unsafe.Pointer(&nb))
	return fp, err
}

func (me *Flobitsstream) NextFloat64Little() (float64, error) {
	b, err := me.NextBitsUnsignedLittle(64)
	bits := uint64(b)
	fp := *(*float64)(unsafe.Pointer(&bits))
	return fp, err
}
