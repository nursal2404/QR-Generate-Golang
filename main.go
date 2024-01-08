package main

import (
	"log"
	"github.com/skip2/go-qrcode"
)

func main() {
	url := "https://github.com/nursal2404/QR-Generate-Golang"

	newQR := "myqr.png"

	err := qrcode.WriteFile(url, qrcode.Medium, 512, newQR)

	if err != nil {
		log.Fatal(err)
	}
}
