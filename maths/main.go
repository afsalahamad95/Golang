package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// random number from crypto
	randomnum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomnum)
}
