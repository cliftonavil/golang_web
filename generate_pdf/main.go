package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, Clifton")
	err := pdf.OutputFileAndClose("clifton.pdf")
	if err != nil {
		log.Fatal("Some thing went wrong!!!", err)
	}
}
