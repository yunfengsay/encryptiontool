package utils

import (
	"bytes"
	"fmt"
	"github.com/isfonzar/filecrypt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func PrintHelp() {
	fmt.Println("encrypt or decrypt")
}

func EncryptHandler(file string) {
	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nFile successfully protected")

}


func DecryptHandler(file string) {

	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)

	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\nFile successfully decrypted.")

}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}
