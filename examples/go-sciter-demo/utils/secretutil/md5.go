package secretutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Md5Hash(s string) []byte {
	m := md5.New()
	m.Write([]byte(s))
	return m.Sum(nil)
}

func Md5Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func ToMd5(source string) string {
	data := []byte(source)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func To16Md5(input string) string {
	return ToMd5(input)[8:24]
}

func ToFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
