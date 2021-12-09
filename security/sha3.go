package main

/*
	sources:
	https://en.wikipedia.org/wiki/SHA-3
	https://www.youtube.com/watch?v=JWskjzgiIa4&t=5065s
*/

import (
	"encoding/hex"
	"fmt"
	"os"
)

var RHO_CONSTS = [][]int{
	{0, 1},
	{2, 0},
	{1, 2},
	{2, 1},
	{3, 2},
	{3, 3},
	{0, 3},
	{1, 0},
	{3, 1},
	{1, 3},
	{4, 1},
	{4, 4},
	{0, 4},
	{3, 0},
	{4, 3},
	{3, 4},
	{2, 3},
	{2, 2},
	{0, 2},
	{4, 0},
	{2, 4},
	{4, 2},
	{1, 4},
	{1, 1},
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

func printHex(buffer []byte) {
	fmt.Println(hex.EncodeToString(buffer))
}

func mod(a int, b int) int {
	return ((a % b) + b) % b
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
	bit := (x*5+z)*64 + y
	return (buffer[bit>>3] >> (bit & 7)) & 1
}

// set the **bit** at location x, y, z
func bitsBufferSet(buffer []byte, x int, y int, z int, value byte) {
	bit := (x*5+z)*64 + y
	if value == 0 {
		buffer[bit>>3] ^= (buffer[bit>>3] & (1 << (bit & 7)))
	} else {
		buffer[bit>>3] |= (1 << (bit & 7))
	}
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
	for y := 0; y < 64; y++ {
		for z := 0; z < 5; z++ {
			leftZ := (z - 1 + 5) % 5
			rightZ := (z + 1) % 5
			frontY := (y - 1 + 64) % 64
			parityLeft := byte(0)
			parityRight := byte(0)

			for x := 0; x < 5; x++ {
				parityLeft ^= bitsBufferGet(buffer, x, y, leftZ)
				parityRight ^= bitsBufferGet(buffer, x, frontY, rightZ)
			}

			for x := 0; x < 5; x++ {
				bitsBufferSet(buffer2, x, y, z, bitsBufferGet(buffer, x, y, z)^parityLeft^parityRight)
			}
		}
	}
}

// compute the keccak step 2
func keccakF2Rho(buffer []byte, buffer2 []byte) {
	// the line (0,0) is not rotated
	bitsBufferSetLine(buffer2, 0, 0, bitsBufferGetLine(buffer, 0, 0))

	// rotate the other lines
	for t := 0; t < 24; t++ {
		x := RHO_CONSTS[t][0]
		z := RHO_CONSTS[t][1]

		for y := 0; y < 64; y++ {
			bitsBufferSet(buffer2, x, y, z, bitsBufferGet(buffer, x, mod(y-(t+1)*(t+2)/2, 64), z))
		}
	}

}

// compute the keccak step 3
func keccakF3Pi(buffer []byte, buffer2 []byte) {
	for x := 0; x < 5; x++ {
		for z := 0; z < 5; z++ {
			bitsBufferSetLine(buffer2, (3*x+2*z)%5, z, bitsBufferGetLine(buffer, x, z))
		}
	}
}

// compute the keccak step 4
func keccakF4Chi(buffer []byte, buffer2 []byte) {
	for x := 0; x < 5; x++ {
		for z := 0; z < 5; z++ {
			a := bitsBufferGetLine(buffer, x, z)
			b := bitsBufferGetLine(buffer, x, (z+1)%5)
			c := bitsBufferGetLine(buffer, x, (z+2)%5)
			bitsBufferSetLine(buffer2, x, z, a^(neg(b)&c))
		}
	}

}

// compute the keccak step 5
func keccakF5Iota(buffer []byte, round int) {
	bitsBufferSetLine(buffer, 0, 0, bitsBufferGetLine(buffer, 0, 0)^IOTA_CONSTS[round])
}

// set a buffer to 0s (debug)
func zeros(buffer []byte) {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 0
	}
}

func keccakF(buffer []byte, buffer2 []byte) {
	// 12 + 2*l
	rounds := 24
	for i := 0; i < rounds; i++ {
		// notice the swap of the buffers in the function calls, it's intentional

		// 1 - theta
		keccakF1Theta(buffer, buffer2)
		// 2 - rho
		keccakF2Rho(buffer2, buffer)
		// 3 - pi
		keccakF3Pi(buffer, buffer2)
		// 4 - chi
		keccakF4Chi(buffer2, buffer)
		// 5 - iota
		keccakF5Iota(buffer, i)
	}
}

// sha3 of a byte array for 256 output size
// (here is where the fun begins)
func sha3(text []byte, salt []byte) []byte {
	toHash := append(salt, text...)
	// Keccak sha3
	// l := 6
	// 25*2**l
	b := 1600 / 8
	// T
	blockSize := 1088 / 8
	// c
	// capacity := b - blockSize

	var blocks int

	if len(toHash)%blockSize == 0 {
		blocks = len(toHash) / blockSize
	} else {
		blocks = len(toHash)/blockSize + 1
	}

	buffer := make([]byte, b)
	buffer2 := make([]byte, b)

	// absorbing phase

	for i := 0; i < blocks; i++ {
		xor(buffer[:blockSize], toHash[i*blockSize:])
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

	var buffer []byte

	switch function {
	case "f", "F":
		buffer = sha3f(os.Args[2], []byte(os.Args[3]))
		break
	case "s", "S":
		buffer = sha3s(os.Args[2], []byte(os.Args[3]))
		break
	default:
		usage()
		return
	}

	printHex(buffer)
}
