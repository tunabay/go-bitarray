# go-bitarray

[![Go Reference](https://pkg.go.dev/badge/github.com/tunabay/go-bitarray.svg)](https://pkg.go.dev/github.com/tunabay/go-bitarray)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## Overview

Package bitarray provides data types and functions for manipulating bit arrays,
aka bit strings, of arbitrary length.

This is designed to handle bit arrays across byte boundaries naturally, without
error-prone bitwise operation code such as shifting, masking, and ORing. It may
be useful when dealing with Huffman coding, raw packet of various protocols, and
binary file formats, etc.

## Usage

```
import (
	"fmt"
	"github.com/tunabay/go-bitarray"
)

func main() {
	// Parse string representation
	ba1, err := bitarray.Parse("111000")
	if err != nil {
		panic(err)
	}
	fmt.Println(ba1) // 111000

	// Slice and Repeat
	ba2 := ba1.Slice(2, 5).Repeat(2)
	fmt.Println(ba2) // 100100

	// Append
	ba3 := ba2.Append(bitarray.MustParse("101011"))
	// alternative formatting
	fmt.Printf("% b\n", ba3) // 10010010 1011

	// Extract bits from []byte across byte boundary
	buf := []byte{0xff, 0x00}
	ba4 := bitarray.NewFromBytes(buf, 4, 7)
	fmt.Println(ba4) // 1111000
}
```
[Run in Go Playground](https://play.golang.org/)

## Documentation and examples:

- Read the [documentation](https://pkg.go.dev/github.com/tunabay/go-bitarray).

## License

go-bitarray is available under the MIT license. See the [LICENSE](LICENSE) file for more information.
