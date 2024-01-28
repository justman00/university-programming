package main

import (
	_ "embed"
	"encoding/binary"
	"fmt"
	"slices"
)

// https://protobuf.dev/programming-guides/encoding/#varints
// encoding unsigned int64 into varint

var (
	//go:embed maxint.uint64
	maxIntByteSlice []byte

	//go:embed 1.uint64
	oneByteSlice []byte

	//go:embed 150.uint64
	oneHundredFiftyByteSlice []byte
)

func main() {
	maxInt := binary.BigEndian.Uint64(maxIntByteSlice)
	one := binary.BigEndian.Uint64(oneByteSlice)
	oneHundredFifty := binary.BigEndian.Uint64(oneHundredFiftyByteSlice)

	testCases := []uint64{maxInt, one, oneHundredFifty, uint64(73676)}

	for _, testCase := range testCases {
		fmt.Println("-----------------------------------")
		encoded := encode(testCase)
		fmt.Printf("encoded as protobuf varint: %x (hex), the original value in hex %x (hex), %d \n", encoded, testCase, testCase)

		decoded := decode(encoded)
		fmt.Printf("decoded from protobuf varint: %d (decimal) \n", decoded)

		if decoded != testCase {
			panic(fmt.Sprintf("decoded %d does not match original %d", decoded, testCase))
		}
		fmt.Println("-----------------------------------")
	}

	fmt.Println("all test cases passed")
}

func encode(num uint64) []byte {
	var encoded []byte

	for num > 0 {
		// take lowest 7 bits
		lowest7Bits := num & 0b01111111

		// reduce n by 7 bits
		num >>= 7

		fmt.Printf("lowest7Bits in binary format before appending: %b \n", lowest7Bits)
		// add correct msb: 1 unless final last bits
		if num > 0 {
			// adds 1 as the most significant bit
			lowest7Bits |= 0b10000000
		}
		fmt.Printf("lowest7Bits in binary format after appending: %b \n", lowest7Bits)
		// push to some sequence of bytes
		encoded = append(encoded, byte(lowest7Bits))
	}

	// return sequence of bytes
	return encoded
}

func decode(encoded []byte) (decoded uint64) {
	slices.Reverse(encoded)
	for _, b := range encoded {
		// shift decoded by 7 bits so that we increase the number by 7 bits
		// on each iteration
		decoded = decoded << 7
		// add the lowest 7 bits of the current byte, we remove the msb
		decoded |= uint64(b & 0x7F)
	}
	return decoded
}
