package main

import (
	"encryption/utils"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		utils.PrintHelp()
		os.Exit(-1)
	}
	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	filePath := os.Args[2]
	function := os.Args[1]
	switch function {
	case "help":
		utils.PrintHelp()
	case "encrypt":
		utils.EncryptHandler(filePath)
	case "decrypt":
		utils.DecryptHandler(filePath)
	default:
		os.Exit(0)
	}
}