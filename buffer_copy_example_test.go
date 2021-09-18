// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray_test

import (
	"fmt"

	"github.com/tunabay/go-bitarray"
)

func ExampleBuffer_CopyBitsFromBytes() {
	ba := bitarray.MustParse("1100-1010 0001-0011 1010-00")
	buf := bitarray.NewBufferFromBitArray(ba)

	fmt.Println(buf)
	buf.CopyBitsFromBytes(2, []byte{0xff, 0x00}, 6, 4)
	fmt.Println(buf)
	buf.CopyBitsFromBytes(6, []byte{0xAA, 0xFF, 0xAA}, 4, 16)
	fmt.Println(buf)

	// Output:
	// 1100101000010011101000
	// 1111001000010011101000
	// 1111001010111111111010
}

func ExampleBuffer_CopyBitsToBytes() {
	ba := bitarray.MustParse("1100-1010 0001")
	buf := bitarray.NewBufferFromBitArray(ba)

	b := make([]byte, 3)

	buf.CopyBitsToBytes(0, b, 0, 12)
	fmt.Printf("%08b\n", b)

	buf.CopyBitsToBytes(4, b, 16, 4)
	buf.CopyBitsToBytes(4, b, 20, 4)
	fmt.Printf("%08b\n", b)

	// Output:
	// [11001010 00010000 00000000]
	// [11001010 00010000 10101010]
}
