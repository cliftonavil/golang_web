package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// OtpContent returns a random number of size four
func OtpContent() string {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		fmt.Println("xxx", e)
	}
	return "We are currently processing ticket number " + strconv.FormatInt(nBig.Int64()+1000, 10)
}

func main() {
	fmt.Println(OtpContent())
}
