// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray

// BitAt returns a single bit at the specified offset as 0 or 1. It panics if
// the off is out of range.
func (buf *Buffer) BitAt(off int) byte {
	switch {
	case off < 0:
		panicf("BitAt: negative off %d.", off)
	case buf.nBits <= off:
		panicf("BitAt: out of range: off=%d >= len=%d.", off, buf.nBits)
	}

	return buf.b[off>>3] >> (7 - off&7) & 1
}

// PutBitAt writes a single bit at the position specified by off in the buffer.
// The bit should be 0 or 1, otherwise its LSB is silently used.
func (buf *Buffer) PutBitAt(off int, bit byte) {
	switch {
	case off < 0:
		panicf("PutBitAt: negative off %d.", off)
	case buf.nBits <= off:
		panicf("PutBitAt: out of range: off=%d >= len=%d.", off, buf.nBits)
	}
	buf.b[off>>3] = buf.b[off>>3] & ^(byte(0x80)>>(off&7)) | ((bit & 1) << (7 - off&7))
}

// BitArrayAt returns bits within the specified range as a BitArray.
func (buf *Buffer) BitArrayAt(off, nBits int) *BitArray {
	switch {
	case off < 0:
		panicf("BitArrayAt: negative off %d.", off)
	case nBits < 0:
		panicf("BitArrayAt: negative nBits %d.", nBits)
	case buf.nBits < off+nBits:
		panicf("BitArrayAt: out of range: off=%d + nBits=%d > len=%d.", off, nBits, buf.nBits)
	case nBits == 0:
		return zeroBitArray
	}

	return NewFromBytes(buf.b, off, nBits)
}

// PutBitArrayAt writes bits from a BitArray onto the specified offset off.
func (buf *Buffer) PutBitArrayAt(off int, ba BitArrayer) {
	switch {
	case off < 0:
		panicf("PutBitArrayAt: negative off %d.", off)
	case ba == nil:
		return
	}
	bab := ba.BitArray()
	switch {
	case buf.nBits < off+bab.nBits:
		panicf("PutBitArrayAt: out of range: off=%d + ba.len=%d > len=%d.", off, bab.nBits, buf.nBits)
	case bab.IsZero():
		return
	case bab.b == nil:
		clearBits(buf.b, off, bab.nBits)
		return
	}
	_ = copyBits(buf.b, bab.b, off, 0, bab.nBits)
}

// ByteAt reads 8 bits starting at the offset off and returns them as a single
// byte. Note that off is in bits, not bytes. If the off is not a multiple of 8,
// 8 bits across a byte boundary are returned.
func (buf *Buffer) ByteAt(off int) byte {
	switch {
	case off < 0:
		panicf("ByteAt: negative off %d.", off)
	case buf.nBits < off+8:
		panicf("ByteAt: out of range: off=%d + 8 > len=%d.", off, buf.nBits)
	}
	i, f := off>>3, off&7
	if f == 0 {
		return buf.b[i]
	}
	return buf.b[i]<<f | buf.b[i+1]>>(8-f)
}

// PutByteAt writes 8 bits of b to the position specified by off in the buffer.
// Note that off is in bits, not bytes. If the off is not a multiple of 8, it
// writes 8 bits across a byte boundary.
func (buf *Buffer) PutByteAt(off int, b byte) {
	switch {
	case off < 0:
		panicf("PutByteAt: negative off %d.", off)
	case buf.nBits < off+8:
		panicf("PutByteAt: out of range: off=%d + 8 > len=%d.", off, buf.nBits)
	}
	i, f := off>>3, off&7
	if f == 0 {
		buf.b[i] = b
	} else {
		copyBits(buf.b[i:], []byte{b}, f, 0, 8) // TODO: optimize
	}
}

// BytesAt reads 8 * nBytes bits starting at the offset off and returns them as
// a byte slice. Note that off is in bits, not bytes. If the off is not a
// multiple of 8, it returns a properly shifted byte slice.
func (buf *Buffer) BytesAt(off, nBytes int) []byte {
	nBits := nBytes << 3
	switch {
	case off < 0:
		panicf("ByteAt: negative off %d.", off)
	case nBytes < 0:
		panicf("ByteAt: negative nBytes %d.", nBytes)
	case buf.nBits < off+nBits:
		panicf("BytesAt: out of range: off=%d + 8 * nBytes=%d > len=%d.", off, nBytes, buf.nBits)
	case nBytes == 0:
		return []byte{}
	}
	ret := make([]byte, nBytes)
	copyBits(ret, buf.b, 0, off, nBits)

	return ret
}

// PutBytesAt writes 8 * len(b) bits of b to the position specified by off in
// the buffer. Note that off is in bits, not bytes. If the off is not a multiple
// of 8, it writes bytes across byte boundaries of the buffer.
func (buf *Buffer) PutBytesAt(off int, b []byte) {
	nBits := len(b) << 3
	switch {
	case off < 0:
		panicf("PutByteAt: negative off %d.", off)
	case buf.nBits < off+nBits:
		panicf("PutByteAt: out of range: off=%d + 8 * b.len=%d > len=%d.", off, len(b), buf.nBits)
	case len(b) == 0:
		return
	}
	copyBits(buf.b, b, off, 0, nBits)
}
