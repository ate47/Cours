package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

// xor two byte arrays, if len(a) != len(b), only the first min(len(a), len(b)) are affected
func xor(a []byte, b []byte) {
	for i := 0; i < len(a) && i < len(b); i++ {
		a[i] ^= b[i]
	}
}

// get the **bit** at location x, y, z
func bitsBufferGet(buffer []byte, x int, y int, z int) byte {
	bit := (z*5+x)*64 + y
	return (buffer[bit>>3] >> (bit & 7)) & 1
}

// set the **bit** at location x, y, z
func bitsBufferSet(buffer []byte, x int, y int, z int, value byte) {
	bit := (x*5+y)*64 + z
	valueShifted := (value & 1) << (bit & 7)
	buffer[bit>>3] = (buffer[bit>>3] | valueShifted) & valueShifted
}

// compute the keccak step 1
func keccakF1Theta(buffer []byte) {
	for x := 0; x < 5; x++ {
		for y := 0; y < 64; y++ {
			leftX := (x - 1 + 5) % 5
			rightX := (x + 1) % 5
			frontY := (y - 1 + 64) % 64
			parityLeft := byte(0)
			parityRight := byte(0)

			for z := 0; z < 5; z++ {
				parityLeft ^= bitsBufferGet(buffer, leftX, y, z)
				parityRight ^= bitsBufferGet(buffer, rightX, frontY, z)
			}

			for z := 0; z < 5; z++ {
				bitsBufferSet(buffer, x, y, z, bitsBufferGet(buffer, x, y, z)^parityLeft^parityRight)
			}
		}
	}
}

// compute the keccak step 2
func keccakF2Rho(buffer []byte) {

}

// compute the keccak step 3
func keccakF3Pi(buffer []byte) {

}

// compute the keccak step 4
func keccakF4Chi(buffer []byte) {

}

// compute the keccak step 5
func keccakF5Iota(buffer []byte) {

}

func keccakF(buffer []byte) {
	// 12 + 2*l
	rounds := 24
	for i := 0; i < rounds; i++ {
		// 1 - theta
		keccakF1Theta(buffer)
		// 2 - rho
		keccakF2Rho(buffer)
		// 3 - pi
		keccakF3Pi(buffer)
		// 4 - chi
		keccakF4Chi(buffer)
		// 5 - iota
		keccakF5Iota(buffer)

	}
}

// sha3 of a byte array for 256 output size
// (here is where the fun begins)
func sha3(text []byte, salt []byte) []byte {
	// Keccak sha3
	// l := 6
	// 25*2**l
	b := 1600 / 8
	// T
	blockSize := 1088 / 8
	// c
	// capacity := b - blockSize

	var blocks int

	if len(text)%blockSize == 0 {
		blocks = len(text) / blockSize
	} else {
		blocks = len(text)/blockSize + 1
	}

	buffer := make([]byte, b)

	// TODO: fill initial buffer

	// absorbing phase

	for i := 0; i < blocks; i++ {
		xor(buffer[:blockSize], text[i*blockSize:])
		keccakF(buffer)
	}

	// squeezing phase (lol)
	return buffer[:256/8]
}

// sha3 of a file
func sha3f(filename string, salt []byte) []byte {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file %s\n", filename)
		os.Exit(-1)
	}

	return sha3(bytes, salt)
}

// sha3 of a string
func sha3s(text string, salt []byte) []byte {
	return sha3([]byte(text), salt)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [f [file]|s [str]] [salt]\n  f - file\n  s - string\n", os.Args[0])
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 4 {
		usage()
	}

	function := os.Args[1]

	switch function {
	case "f", "F":
		fmt.Println(hex.EncodeToString(sha3f(os.Args[2], []byte(os.Args[3]))))
		break
	case "s", "S":
		fmt.Println(hex.EncodeToString(sha3s(os.Args[2], []byte(os.Args[3]))))
		break
	default:
		usage()
		break
	}
}
