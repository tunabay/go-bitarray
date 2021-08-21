// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package bitarray_test

import (
	"fmt"

	"github.com/tunabay/go-bitarray"
)

func ExampleBuffer() {
	buf := bitarray.NewBuffer(32)
	fmt.Println(buf)
	buf.PutBitAt(0, 1)
	buf.PutBitAt(1, 1)
	fmt.Println(buf)
	buf.PutBitArrayAt(8, bitarray.MustParse("1010101"))
	fmt.Println(buf)
	buf.FillBitsAt(16, 4, 1)
	fmt.Println(buf)
	buf.PutBitArrayAt(24, bitarray.MustParse("1111-0000"))
	fmt.Println(buf)
	buf.ToggleBitsAt(24, 8)
	fmt.Println(buf)

	fmt.Printf("% b\n", buf.BitArray())

	// Output:
	// 00000000000000000000000000000000
	// 11000000000000000000000000000000
	// 11000000101010100000000000000000
	// 11000000101010101111000000000000
	// 11000000101010101111000011110000
	// 11000000101010101111000000001111
	// 11000000 10101010 11110000 00001111
}

func ExampleNewBuffer() {
	buf0 := bitarray.NewBuffer(0)
	buf64 := bitarray.NewBuffer(64)

	fmt.Println(buf0)
	fmt.Println(buf64)

	// Output:
	// 0000000000000000000000000000000000000000000000000000000000000000
}

func ExampleNewBufferFromBitArray() {
	ba := bitarray.MustParse("1111-1111 0000-0000 1111-1111")
	buf := bitarray.NewBufferFromBitArray(ba)

	fmt.Println(buf)

	// Output:
	// 111111110000000011111111
}

func ExampleBuffer_Len() {
	buf := bitarray.NewBuffer(4096)

	fmt.Println(buf.Len())

	// Output:
	// 4096
}
