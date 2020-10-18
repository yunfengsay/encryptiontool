package utils

func Encrypt(path string) {
	if IsDir(path) {
		EncryptDir(path)
	}
}

func EncryptDir(path string)  {

}
