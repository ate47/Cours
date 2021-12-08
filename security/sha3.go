package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [f [file]|s [str]] [salt]\n  f - file\n  s - string\n", os.Args[0])
	os.Exit(-1)
}

// sha3 of a byte array
// (here is where the fun begins)
func sha3(text []byte, salt []byte) []byte {

	return text
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
