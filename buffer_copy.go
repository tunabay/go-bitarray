// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray

// CopyBitsFromBytes reads nBits bits from b at the offset bOff, and write them
// into the buffer at the offset off.
func (buf *Buffer) CopyBitsFromBytes(off int, b []byte, bOff, nBits int) {
	switch {
	case off < 0:
		panicf("CopyBitsFromBytes: negative off %d.", off)
	case buf.nBits < off+nBits:
		panicf("CopyBitsFromBytes: out of range: off=%d + nBits=%d > len=%d.", off, nBits, buf.nBits)
	case nBits == 0:
		return
	}
	copyBits(buf.b, b, off, bOff, nBits)
}

// CopyBitsToBytes reads nBits bits of the buffer starting at the offset off,
// and write them into the byte slice b at the offset bOff.
func (buf *Buffer) CopyBitsToBytes(off int, b []byte, bOff, nBits int) {
	switch {
	case off < 0:
		panicf("CopyBitsToBytes: negative off %d.", off)
	case buf.nBits < off+nBits:
		panicf("CopyBitsToBytes: out of range: off=%d + nBits=%d > len=%d.", off, nBits, buf.nBits)
	case nBits == 0:
		return
	}
	copyBits(b, buf.b, bOff, off, nBits)
}
