package flobits

import (
	"bytes"
	"testing"
)

func TestReaderOverSimple(t *testing.T) {
	// create a 16 byte buffer
	testslice := [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	testbuffer := bytes.NewBuffer(testslice[:])
	foor := NewFlobitsDecoder(testbuffer, 16)

	// read two uint64s
	foor.GetBitsUnsignedBig(64)
	foor.GetBitsUnsignedBig(64)

	// the bitstream should be EOF, so the following NextBitsUnsignedBig must yeild an EOF error
	// _, err := foor.NextFloat32Little()
	// if err == nil {
	// 	t.Errorf("NextBitsUnsignedBig must return non-nil error")
	// }
}

// GENERICS!!
/*
package main

import (
	"fmt"
)

// AllInteger is a type constriant that restrics an allowed geric type to only integers
type AllInteger interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
}

func min[T AllInteger](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func getsix[T AllInteger]() (T, error) {
	return 6, nil
}

func main() {
	var i1 uint32 = 3
	var i2 uint32 = 4
	m := min(i1,i2)
	fmt.Printf("%T\n",m)
	fmt.Println(min(i1, i2))

	zzz, _ := getsix[uint64]()
	fmt.Printf("%d,%T\n",zzz,zzz)
}
*/
