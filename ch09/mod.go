package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"go-execise/ch09/formatter"
	"go-execise/ch09/math"
	"math/rand"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}

func seedRand() *rand.Rand {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed with cryptographic random number generator")
	}

	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}
