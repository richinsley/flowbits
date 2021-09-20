package flobits

// Uint8 helprs

func (me *Bitstream) PutWithSizeUint8(v uint8, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithSizeUint8(n uint32) (uint8, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint8(v), err
}

func (me *Bitstream) NextWithSizeUint8(n uint32) (uint8, error) {
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

func (me *Bitstream) PutWithSizeInt8(v int8, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithSizeInt8(n uint32) (int8, error) {
	v, err := me.GetBitsSignedBig(n)
	return int8(v), err
}

func (me *Bitstream) NextWithSizeInt8(n uint32) (int8, error) {
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

// Uint16 helpers

func (me *Bitstream) PutWithSizeUint16(v uint16, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithSizeUint16(n uint32) (uint16, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint16(v), err
}

func (me *Bitstream) NextWithSizeUint16(n uint32) (uint16, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint16(v), err
}

func (me *Bitstream) PutUint16(v uint16) error {
	return me.PutBitsUnsignedBig(uint64(v), 16)
}

func (me *Bitstream) GetUint16() (uint16, error) {
	v, err := me.GetBitsUnsignedBig(16)
	return uint16(v), err
}

func (me *Bitstream) NextUint16() (uint16, error) {
	v, err := me.NextBitsUnsignedBig(16)
	return uint16(v), err
}

// Int16 helpers

func (me *Bitstream) PutWithSizeInt16(v int16, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithSizeInt16(n uint32) (int16, error) {
	v, err := me.GetBitsSignedBig(n)
	return int16(v), err
}

func (me *Bitstream) NextWithSizeInt16(n uint32) (int16, error) {
	v, err := me.NextBitsSignedBig(n)
	return int16(v), err
}

func (me *Bitstream) PutInt16(v int16) error {
	return me.PutBitsSignedBig(int64(v), 16)
}

func (me *Bitstream) GetInt16() (int16, error) {
	v, err := me.GetBitsSignedBig(16)
	return int16(v), err
}

func (me *Bitstream) NextInt16() (int16, error) {
	v, err := me.NextBitsSignedBig(16)
	return int16(v), err
}

// Uint32 helpers

func (me *Bitstream) PutWithSizeUint32(v uint32, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithSizeUint32(n uint32) (uint32, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint32(v), err
}

func (me *Bitstream) NextWithSizeUint32(n uint32) (uint32, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint32(v), err
}

func (me *Bitstream) PutUint32(v uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), 32)
}

func (me *Bitstream) GetUint32() (uint32, error) {
	v, err := me.GetBitsUnsignedBig(32)
	return uint32(v), err
}

func (me *Bitstream) NextUint32() (uint32, error) {
	v, err := me.NextBitsUnsignedBig(32)
	return uint32(v), err
}

// Int32 helpers

func (me *Bitstream) PutWithSizeInt32(v int32, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithSizeInt32(n uint32) (int32, error) {
	v, err := me.GetBitsSignedBig(n)
	return int32(v), err
}

func (me *Bitstream) NextWithSizeInt32(n uint32) (int32, error) {
	v, err := me.NextBitsSignedBig(n)
	return int32(v), err
}

func (me *Bitstream) PutInt32(v int32) error {
	return me.PutBitsSignedBig(int64(v), 32)
}

func (me *Bitstream) GetInt32() (int32, error) {
	v, err := me.GetBitsSignedBig(32)
	return int32(v), err
}

func (me *Bitstream) NextInt32() (int32, error) {
	v, err := me.NextBitsSignedBig(32)
	return int32(v), err
}

// Uint64 helpers

func (me *Bitstream) PutWithSizeUint64(v uint64, n uint32) error {
	return me.PutBitsUnsignedBig(uint64(v), n)
}

func (me *Bitstream) GetWithSizeUint64(n uint32) (uint64, error) {
	v, err := me.GetBitsUnsignedBig(n)
	return uint64(v), err
}

func (me *Bitstream) NextWithSizeUint64(n uint32) (uint64, error) {
	v, err := me.NextBitsUnsignedBig(n)
	return uint64(v), err
}

func (me *Bitstream) PutUint64(v uint64) error {
	return me.PutBitsUnsignedBig(uint64(v), 64)
}

func (me *Bitstream) GetUint64() (uint64, error) {
	v, err := me.GetBitsUnsignedBig(64)
	return uint64(v), err
}

func (me *Bitstream) NextUint64() (uint64, error) {
	v, err := me.NextBitsUnsignedBig(64)
	return uint64(v), err
}

// Int64 helpers

func (me *Bitstream) PutWithSizeInt64(v int64, n uint32) error {
	return me.PutBitsSignedBig(int64(v), n)
}

func (me *Bitstream) GetWithSizeInt64(n uint32) (int64, error) {
	v, err := me.GetBitsSignedBig(n)
	return int64(v), err
}

func (me *Bitstream) NextWithSizeInt64(n uint32) (int64, error) {
	v, err := me.NextBitsSignedBig(n)
	return int64(v), err
}

func (me *Bitstream) PutInt64(v int64) error {
	return me.PutBitsSignedBig(int64(v), 64)
}

func (me *Bitstream) GetInt64() (int64, error) {
	v, err := me.GetBitsSignedBig(64)
	return int64(v), err
}

func (me *Bitstream) NextInt64() (int64, error) {
	v, err := me.NextBitsSignedBig(64)
	return int64(v), err
}
