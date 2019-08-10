package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("Clifton")
	fmt.Println(c)
	c = getCode("Avil")
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
