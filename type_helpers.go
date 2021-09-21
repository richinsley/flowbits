package flowbits

// Uint8 helpers

func (me *Bitstream) PutWithBitCountUint8(v uint8, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountUint8(n uint32) (uint8, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint8(v), err
}

func (me *Bitstream) NextWithBitCountUint8(n uint32) (uint8, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint8(v), err
}

func (me *Bitstream) PutUint8(v uint8) error {
	return me.PutBitsUnsignedBig(uint64(v), 8)
}

func (me *Bitstream) GetUint8() (uint8, error) {
	v, err := me.GetBitsUnsignedBig(8)
	return uint8(v), err
}

func (me *Bitstream) NextUint8() (uint8, error) {
	v, err := me.NextBitsUnsignedBig(8)
	return uint8(v), err
}

// Int8 helpers

func (me *Bitstream) PutWithBitCountInt8(v int8, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithBitCountInt8(n uint32) (int8, error) {
	v, err := me.GetBitsSignedBig(n)
	return int8(v), err
}

func (me *Bitstream) NextWithBitCountInt8(n uint32) (int8, error) {
	v, err := me.NextBitsSignedBig(n)
	return int8(v), err
}

func (me *Bitstream) PutInt8(v int8) error {
	return me.PutBitsSignedBig(int64(v), 8)
}

func (me *Bitstream) GetInt8() (int8, error) {
	v, err := me.GetBitsSignedBig(8)
	return int8(v), err
}

func (me *Bitstream) NextInt8() (int8, error) {
	v, err := me.NextBitsSignedBig(8)
	return int8(v), err
}

// Uint16 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigUint16(v uint16, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountBigUint16(n uint32) (uint16, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint16(v), err
}

func (me *Bitstream) NextWithBitCountBigUint16(n uint32) (uint16, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint16(v), err
}

func (me *Bitstream) PutBigUint16(v uint16) error {
	return me.PutBitsUnsignedBig(uint64(v), 16)
}

func (me *Bitstream) GetBigUint16() (uint16, error) {
	v, err := me.GetBitsUnsignedBig(16)
	return uint16(v), err
}

func (me *Bitstream) NextBigUint16() (uint16, error) {
	v, err := me.NextBitsUnsignedBig(16)
	return uint16(v), err
}

// Int16 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigInt16(v int16, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithBitCountBigInt16(n uint32) (int16, error) {
	v, err := me.GetBitsSignedBig(n)
	return int16(v), err
}

func (me *Bitstream) NextWithBitCountBigInt16(n uint32) (int16, error) {
	v, err := me.NextBitsSignedBig(n)
	return int16(v), err
}

func (me *Bitstream) PutBigInt16(v int16) error {
	return me.PutBitsSignedBig(int64(v), 16)
}

func (me *Bitstream) GetBigInt16() (int16, error) {
	v, err := me.GetBitsSignedBig(16)
	return int16(v), err
}

func (me *Bitstream) NextBigInt16() (int16, error) {
	v, err := me.NextBitsSignedBig(16)
	return int16(v), err
}

// Uint32 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigUint32(v uint32, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountBigUint32(n uint32) (uint32, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint32(v), err
}

func (me *Bitstream) NextWithBitCountBigUint32(n uint32) (uint32, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint32(v), err
}

func (me *Bitstream) PutBigUint32(v uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), 32)
}

func (me *Bitstream) GetBigUint32() (uint32, error) {
	v, err := me.GetBitsUnsignedBig(32)
	return uint32(v), err
}

func (me *Bitstream) NextBigUint32() (uint32, error) {
	v, err := me.NextBitsUnsignedBig(32)
	return uint32(v), err
}

// Int32 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigInt32(v int32, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithBitCountBigInt32(n uint32) (int32, error) {
	v, err := me.GetBitsSignedBig(n)
	return int32(v), err
}

func (me *Bitstream) NextWithBitCountBigInt32(n uint32) (int32, error) {
	v, err := me.NextBitsSignedBig(n)
	return int32(v), err
}

func (me *Bitstream) PutBigInt32(v int32) error {
	return me.PutBitsSignedBig(int64(v), 32)
}

func (me *Bitstream) GetBigInt32() (int32, error) {
	v, err := me.GetBitsSignedBig(32)
	return int32(v), err
}

func (me *Bitstream) NextBigInt32() (int32, error) {
	v, err := me.NextBitsSignedBig(32)
	return int32(v), err
}

// Uint64 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigUint64(v uint64, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountBigUint64(n uint32) (uint64, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint64(v), err
}

func (me *Bitstream) NextWithBitCountBigUint64(n uint32) (uint64, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint64(v), err
}

func (me *Bitstream) PutBigUint64(v uint64) error {
	return me.PutBitsUnsignedBig(uint64(v), 64)
}

func (me *Bitstream) GetBigUint64() (uint64, error) {
	v, err := me.GetBitsUnsignedBig(64)
	return uint64(v), err
}

func (me *Bitstream) NextBigUint64() (uint64, error) {
	v, err := me.NextBitsUnsignedBig(64)
	return uint64(v), err
}

// Int64 Big Endian helpers

func (me *Bitstream) PutWithBitCountBigInt64(v int64, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithBitCountBigInt64(n uint32) (int64, error) {
	v, err := me.GetBitsSignedBig(n)
	return int64(v), err
}

func (me *Bitstream) NextWithBitCountBigInt64(n uint32) (int64, error) {
	v, err := me.NextBitsSignedBig(n)
	return int64(v), err
}

func (me *Bitstream) PutBigInt64(v int64) error {
	return me.PutBitsSignedBig(int64(v), 64)
}

func (me *Bitstream) GetBigInt64() (int64, error) {
	v, err := me.GetBitsSignedBig(64)
	return int64(v), err
}

func (me *Bitstream) NextBigInt64() (int64, error) {
	v, err := me.NextBitsSignedBig(64)
	return int64(v), err
}

// Uint16 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleUint16(v uint16, n uint32) error {
	return me.PutBitsUnsignedLittle(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleUint16(n uint32) (uint16, error) {
	v, err := me.GetBitsUnsignedLittle(n)
	return uint16(v), err
}

func (me *Bitstream) NextWithBitCountLittleUint16(n uint32) (uint16, error) {
	v, err := me.NextBitsUnsignedLittle(n)
	return uint16(v), err
}

func (me *Bitstream) PutLittleUint16(v uint16) error {
	return me.PutBitsUnsignedLittle(uint64(v), 16)
}

func (me *Bitstream) GetLittleUint16() (uint16, error) {
	v, err := me.GetBitsUnsignedLittle(16)
	return uint16(v), err
}

func (me *Bitstream) NextLittleUint16() (uint16, error) {
	v, err := me.NextBitsUnsignedLittle(16)
	return uint16(v), err
}

// Int16 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleInt16(v int16, n uint32) error {
	return me.PutBitsSignedLittle(int64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleInt16(n uint32) (int16, error) {
	v, err := me.GetBitsSignedLittle(n)
	return int16(v), err
}

func (me *Bitstream) NextWithBitCountLittleInt16(n uint32) (int16, error) {
	v, err := me.NextBitsSignedLittle(n)
	return int16(v), err
}

func (me *Bitstream) PutLittleInt16(v int16) error {
	return me.PutBitsSignedLittle(int64(v), 16)
}

func (me *Bitstream) GetLittleInt16() (int16, error) {
	v, err := me.GetBitsSignedLittle(16)
	return int16(v), err
}

func (me *Bitstream) NextLittleInt16() (int16, error) {
	v, err := me.NextBitsSignedLittle(16)
	return int16(v), err
}

// Uint32 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleUint32(v uint32, n uint32) error {
	return me.PutBitsUnsignedLittle(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleUint32(n uint32) (uint32, error) {
	v, err := me.GetBitsUnsignedLittle(n)
	return uint32(v), err
}

func (me *Bitstream) NextWithBitCountLittleUint32(n uint32) (uint32, error) {
	v, err := me.NextBitsUnsignedLittle(n)
	return uint32(v), err
}

func (me *Bitstream) PutLittleUint32(v uint32) error {
	return me.PutBitsUnsignedLittle(uint64(v), 32)
}

func (me *Bitstream) GetLittleUint32() (uint32, error) {
	v, err := me.GetBitsUnsignedLittle(32)
	return uint32(v), err
}

func (me *Bitstream) NextLittleUint32() (uint32, error) {
	v, err := me.NextBitsUnsignedLittle(32)
	return uint32(v), err
}

// Int32 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleInt32(v int32, n uint32) error {
	return me.PutBitsSignedLittle(int64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleInt32(n uint32) (int32, error) {
	v, err := me.GetBitsSignedLittle(n)
	return int32(v), err
}

func (me *Bitstream) NextWithBitCountLittleInt32(n uint32) (int32, error) {
	v, err := me.NextBitsSignedLittle(n)
	return int32(v), err
}

func (me *Bitstream) PutLittleInt32(v int32) error {
	return me.PutBitsSignedLittle(int64(v), 32)
}

func (me *Bitstream) GetLittleInt32() (int32, error) {
	v, err := me.GetBitsSignedLittle(32)
	return int32(v), err
}

func (me *Bitstream) NextLittleInt32() (int32, error) {
	v, err := me.NextBitsSignedLittle(32)
	return int32(v), err
}

// Uint64 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleUint64(v uint64, n uint32) error {
	return me.PutBitsUnsignedLittle(uint64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleUint64(n uint32) (uint64, error) {
	v, err := me.GetBitsUnsignedLittle(n)
	return uint64(v), err
}

func (me *Bitstream) NextWithBitCountLittleUint64(n uint32) (uint64, error) {
	v, err := me.NextBitsUnsignedLittle(n)
	return uint64(v), err
}

func (me *Bitstream) PutLittleUint64(v uint64) error {
	return me.PutBitsUnsignedLittle(uint64(v), 64)
}

func (me *Bitstream) GetLittleUint64() (uint64, error) {
	v, err := me.GetBitsUnsignedLittle(64)
	return uint64(v), err
}

func (me *Bitstream) NextLittleUint64() (uint64, error) {
	v, err := me.NextBitsUnsignedLittle(64)
	return uint64(v), err
}

// Int64 Little Endian helpers

func (me *Bitstream) PutWithBitCountLittleInt64(v int64, n uint32) error {
	return me.PutBitsSignedLittle(int64(v), n)
}

func (me *Bitstream) GetWithBitCountLittleInt64(n uint32) (int64, error) {
	v, err := me.GetBitsSignedLittle(n)
	return int64(v), err
}

func (me *Bitstream) NextWithBitCountLittleInt64(n uint32) (int64, error) {
	v, err := me.NextBitsSignedLittle(n)
	return int64(v), err
}

func (me *Bitstream) PutLittleInt64(v int64) error {
	return me.PutBitsSignedLittle(int64(v), 64)
}

func (me *Bitstream) GetLittleInt64() (int64, error) {
	v, err := me.GetBitsSignedLittle(64)
	return int64(v), err
}

func (me *Bitstream) NextLittleInt64() (int64, error) {
	v, err := me.NextBitsSignedLittle(64)
	return int64(v), err
}
