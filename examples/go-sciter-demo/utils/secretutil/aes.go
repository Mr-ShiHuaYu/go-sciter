package secretutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(toEncryptArray []byte, keyArray []byte) ([]byte, error) {
	block, err := aes.NewCipher(keyArray)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	toEncryptArray = PKCS5Padding(toEncryptArray, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyArray[:blockSize])
	crypted := make([]byte, len(toEncryptArray))
	blockMode.CryptBlocks(crypted, toEncryptArray)
	return crypted, nil
}

func AesDecrypt(toDecryptArray []byte, keyArray []byte) ([]byte, error) {
	block, err := aes.NewCipher(keyArray)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyArray[:blockSize])
	origData := make([]byte, len(toDecryptArray))
	blockMode.CryptBlocks(origData, toDecryptArray)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
