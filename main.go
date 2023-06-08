package main

import (
	"fmt"
	"os"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	url := "https://example.org"
	size := 256
	name := "qr"

	// go trough each argument
	for _, arg := range os.Args[1:] {
		if len(arg) > 4 && arg[:4] == "url=" {
			url = arg[4:]
		}

		if len(arg) > 5 && arg[:5] == "size=" {
			arg, err := strconv.Atoi(arg[5:])

			if err != nil {
				fmt.Println("Error: size must be an integer")
				os.Exit(1)
			}

			size = arg
		}

		if len(arg) > 5 && arg[:5] == "name=" {
			name = arg[5:]
		}
	}

	// Create a folder to store the QR code
	os.Mkdir(name, 0755)
	os.Chdir(name)

	// Create a QR code
	qrcode.WriteFile(url, qrcode.Medium, size, name+".png")
}
