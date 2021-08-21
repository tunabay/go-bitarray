// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray

import (
	"fmt"
)

// D returns the string representing of its internal state.
func (buf *Buffer) D() string {
	if buf == nil {
		return "<nil>"
	}
	return fmt.Sprintf("BUF{nbit=%d, b=%08b}", buf.nBits, buf.b)
}

// V validate the internal data representation. It panics on failure.
func (buf *Buffer) V() {
	switch {
	case buf == nil:
		return

	case buf.nBits < 0:
		panicf("V: negative nBits %d", buf.nBits)

	case buf.b != nil && len(buf.b) == 0:
		panicf("V: buf.b is an empty slice, must be nil: %08b", buf.b)

	case buf.b == nil && buf.nBits != 0:
		panicf("V: buf.b is nil, must be non nil for nbits=%d", buf.nBits)

	case buf.b == nil:
		return

	case len(buf.b) != (buf.nBits+7)>>3:
		panicf("V: wrong len: len=%d, nBits=%d: %08b", len(buf.b), buf.nBits, buf.b)

		// case cap(buf.b)&7 != 0:
		// 	panicf("V: wrong cap: cap=%d, len=%d, nBits=%d", cap(buf.b), len(buf.b), buf.nBits)
	}

	if fb := buf.nBits & 7; fb != 0 {
		mask := byte(0xff) >> fb
		if lb := buf.b[len(buf.b)-1] & mask; lb != 0 {
			panicf("V: non-zero padding bits: nfrac=%d, lastbyte=%08b", fb, lb)
		}
	}
}
