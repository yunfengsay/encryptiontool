package utils

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"path/filepath"
)

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password1, _ := terminal.ReadPassword(0)
	if !validatePassword(password, password1) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}
	return password
}

// 判断给定路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func getAllFilesInFolder(path string) ([]string, error)  {
	var result []string
	filepath.Walk(path, func(path string, fi os.FileInfo, err error) error {
		if err != nil{
			log.Println(err.Error())
			return err
		}
		if !fi.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, nil
}