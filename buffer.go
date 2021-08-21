// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray

import (
	"fmt"
)

// Buffer is a bit array buffer whose contents can be updated by partial reading
// and writing with an offset. It is not safe for concurrent use by multiple
// goroutines. The zero value for Buffer represents a zero length buffer that
// can be resized and used.
type Buffer BitArray

// Buffer fields:
// .b     []byte // nil only for zero length
// .nBits int

// NewBuffer creates a Buffer with the specified bit length.
func NewBuffer(nBits int) *Buffer {
	switch {
	case nBits < 0:
		panicf("NewBuffer: negative nBits %d.", nBits)
	case nBits == 0:
		return &Buffer{}
	}

	return &Buffer{
		b:     allocByteSlice((nBits + 7) >> 3),
		nBits: nBits,
	}
}

// NewBufferFromBitArray creates a new Buffer with the same bit length and
// initial content as the specified BitArray.
func NewBufferFromBitArray(ba BitArrayer) *Buffer {
	if ba == nil {
		return &Buffer{}
	}
	bab := ba.BitArray()
	buf := NewBuffer(bab.Len())
	if 0 < buf.nBits {
		copy(buf.b, bab.b)
	}
	return buf
}

// IsZero returns whether the Buffer is zero length.
func (buf *Buffer) IsZero() bool {
	return buf.Len() == 0
}

// Len returns the number of bits contained in the buffer.
func (buf *Buffer) Len() int {
	if buf == nil {
		return 0
	}
	return buf.nBits
}

// Clone clones the Buffer with its content.
func (buf *Buffer) Clone() *Buffer {
	if buf.Len() == 0 {
		return &Buffer{}
	}
	b := allocByteSlice(len(buf.b))
	copy(b, buf.b)

	return &Buffer{b: b, nBits: buf.nBits}
}

// BitArray creates an imuurable BitArray from the current content.
func (buf *Buffer) BitArray() *BitArray {
	return NewFromBytes(buf.b, 0, buf.nBits)
}

// String returns the string representation of the current content.
func (buf Buffer) String() string {
	sb := make([]byte, buf.nBits)
	for i := 0; i < buf.nBits; i++ {
		sb[i] = '0' + buf.b[i>>3]>>(7-i&7)&1
	}
	return string(sb)
}

// Resize resizes the Buffer to the size specified by nBits. When expanding, all
// bits in the new range to be extended are initialized with 0. When shrinking,
// the extra bits are truncated. In either case, the align specifies whether to
// fix the MSBs or the LSBs.
func (buf *Buffer) Resize(nBits int, align Alignment) {
	switch {
	case nBits < 0:
		panicf("Resize: negative nBits %d.", nBits)
	case nBits == buf.nBits:
		return
	case nBits == 0:
		buf.b = nil
		buf.nBits = 0
		return
	}
	nBytes := (nBits + 7) >> 3

	// AlignLeft
	if align == AlignLeft {
		// shrink
		if nBits < buf.nBits {
			buf.b = buf.b[:nBytes]
			if nBits&7 != 0 {
				buf.b[nBits>>3] &= 0xff << (8 - nBits&7)
			}
			buf.nBits = nBits
			return
		}

		// extend
		if nBytes <= cap(buf.b) { // enough cap
			n := len(buf.b)
			buf.b = buf.b[:nBytes]
			for ; n < nBytes; n++ {
				buf.b[n] = 0
			}
			buf.nBits = nBits
			return
		}

		newb := allocByteSlice(nBytes)
		copy(newb, buf.b)
		buf.b = newb
		buf.nBits = nBits
		return
	}

	// AlignRight
	// shrink
	if nBits < buf.nBits {
		if nBits&7 == buf.nBits&7 { // no shift
			buf.b = buf.b[len(buf.b)-nBytes:]
			buf.nBits = nBits
			return
		}
		newb := allocByteSlice(nBytes)
		_ = copyBits(newb, buf.b, 0, buf.nBits-nBits, nBits)
		buf.b = newb
		buf.nBits = nBits
		return
	}

	// extend
	newb := allocByteSlice(nBytes)
	if buf.b != nil {
		_ = copyBits(newb, buf.b, nBits-buf.nBits, 0, buf.nBits)
	}
	buf.b = newb
	buf.nBits = nBits
}

// FillBitsAt sets the nBits bits starting at off to the value bit.
func (buf *Buffer) FillBitsAt(off, nBits int, bit byte) {
	switch {
	case off < 0:
		panicf("FillBitsAt: negative off %d.", off)
	case nBits < 0:
		panicf("FillBitsAt: negative nBits %d.", nBits)
	case buf.nBits < off+nBits:
		panicf("FillBitsAt: out of range: off=%d + nBits=%d > len=%d.", off, nBits, buf.nBits)
	case bit&1 == 0:
		clearBits(buf.b, off, nBits)
	default:
		setBits(buf.b, off, nBits)
	}
}

// Format implements the fmt.Formatter interface to format Buffer value using
// the standard fmt.Printf family functions.
func (buf Buffer) Format(s fmt.State, verb rune) { BitArray(buf).Format(s, verb) }
