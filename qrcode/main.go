package main

import (
	"fmt"

	qrcode "github.com/yeqown/go-qrcode"
)

func main() {
	qrc, err := qrcode.New("Clifton Avil Dsouza")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err := qrc.Save("qrcode.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
