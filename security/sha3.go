package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [f [file]|s [str]]\n  f - file\n  s - string\n", os.Args[0])
	os.Exit(-1)
}

// sha3 of a byte array
// (here is where the fun begins)
func sha3(text []byte) []byte {

	return text
}

// sha3 of a file
func sha3f(filename string) []byte {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read file %s\n", filename)
		os.Exit(-1)
	}

	return sha3(bytes)
}

// sha3 of a string
func sha3s(text string) []byte {
	return sha3([]byte(text))
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	function := os.Args[1]

	switch function {
	case "f":
		fmt.Println(hex.EncodeToString(sha3f(os.Args[2])))
		break
	case "s":
		fmt.Println(hex.EncodeToString(sha3s(os.Args[2])))
		break
	default:
		usage()
		break
	}
}
