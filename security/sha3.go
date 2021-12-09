package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

var ROTATION_CONSTS = [][]int{
	{0, 1, 62, 28, 27},
	{36, 44, 6, 55, 20},
	{3, 10, 43, 25, 39},
	{41, 45, 15, 21, 8},
	{18, 2, 61, 56, 14},
}

var IOTA_CONSTS = []uint64{
	0x0000_0000_0000_0001,
	0x0000_0000_0000_8082,
	0x8000_0000_0000_808a,
	0x8000_0000_8000_8000,
	0x0000_0000_0000_808b,
	0x0000_0000_8000_0001,
	0x8000_0000_8000_8081,
	0x8000_0000_0000_8009,
	0x0000_0000_0000_008a,
	0x0000_0000_0000_0088,
	0x0000_0000_8000_8009,
	0x0000_0000_8000_000a,
	0x0000_0000_8000_808b,
	0x8000_0000_0000_008b,
	0x8000_0000_0000_8089,
	0x8000_0000_0000_8003,
	0x8000_0000_0000_8002,
	0x8000_0000_0000_0080,
	0x0000_0000_0000_800a,
	0x8000_0000_8000_000a,
	0x8000_0000_8000_8081,
	0x8000_0000_0000_8080,
	0x0000_0000_8000_0001,
	0x8000_0000_8000_8008,
}

func neg(v uint64) uint64 {
	return v ^ uint64(0xFFFF_FFFF_FFFF_FFFF)
}

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
	bit := (z*5+x)*64 + y
	valueShifted := (value & 1) << (bit & 7)
	buffer[bit>>3] = (buffer[bit>>3] | valueShifted) & valueShifted
}

// set the 64 bits line at location x, z
func bitsBufferSetLine(buffer []byte, x int, z int, value uint64) {
	for y := 0; y < 64; y++ {
		bitsBufferSet(buffer, x, y, z, byte((value>>y)&1))
	}
}

// get the 64 bits line at location x, z
func bitsBufferGetLine(buffer []byte, x int, z int) uint64 {
	line := uint64(0)
	for y := 0; y < 64; y++ {
		line |= uint64(bitsBufferGet(buffer, x, y, z)) << y
	}
	return line
}

// compute the keccak step 1
func keccakF1Theta(buffer []byte, buffer2 []byte) {
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
				bitsBufferSet(buffer2, x, y, z, bitsBufferGet(buffer, x, y, z)^parityLeft^parityRight)
			}
		}
	}
}

// compute the keccak step 2 and 3
func keccakF23RhoPi(buffer []byte, buffer2 []byte) {
	for x := 0; x < 5; x++ {
		for z := 0; z < 5; z++ {
			rotation := ROTATION_CONSTS[x][z]

			// step 2 rho - rotate the (x,z) into line
			line := uint64(0)
			for y := 0; y < 64; y++ {
				line |= uint64(bitsBufferGet(buffer, x, (y+rotation)%64, z)) << y
			}

			// step 3 pi - set the line on another location
			bitsBufferSetLine(buffer2, z, (2*x+3*z)%5, line)
		}
	}
}

// compute the keccak step 4
func keccakF4Chi(buffer []byte, buffer2 []byte) {
	for x := 0; x < 5; x++ {
		for z := 0; z < 5; z++ {
			line := bitsBufferGetLine(buffer, x, z)
			line2 := bitsBufferGetLine(buffer, (x+1)%5, z)
			line3 := bitsBufferGetLine(buffer, (x+2)%5, z)
			bitsBufferSetLine(buffer2, x, z, line^(neg(line2)&line3))
		}
	}

}

// compute the keccak step 5
func keccakF5Iota(buffer []byte, buffer2 []byte, round int) {
	copy(buffer2, buffer)
	line := bitsBufferGetLine(buffer, 0, 0)
	bitsBufferSetLine(buffer2, 0, 0, line^IOTA_CONSTS[round])
}

func keccakF(buffer []byte, buffer2 []byte) {
	// 12 + 2*l
	rounds := 24
	for i := 0; i < rounds; i++ {
		// notice the swap the buffer in the process calls, it's intentional

		// 1 - theta
		keccakF1Theta(buffer, buffer2)
		// 2 - rho and 3 - pi
		keccakF23RhoPi(buffer2, buffer)
		// 4 - chi
		keccakF4Chi(buffer, buffer2)
		// 5 - iota
		keccakF5Iota(buffer2, buffer, i)
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
	buffer2 := make([]byte, b)

	// TODO: fill initial buffer

	// absorbing phase

	for i := 0; i < blocks; i++ {
		xor(buffer[:blockSize], text[i*blockSize:])
		keccakF(buffer, buffer2)
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
