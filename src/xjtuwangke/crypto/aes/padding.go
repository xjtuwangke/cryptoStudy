package aes

import (
	"bytes"
)

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Padding(ciphertext []byte, blockSize int, mode int64) []byte {
	if mode&Padding_PKCS5 != 0 {
		return PKCS5Padding(ciphertext, blockSize)
	}
	if mode&Padding_Zero != 0 {
		return ZeroPadding(ciphertext, blockSize)
	}
	return PKCS5Padding(ciphertext, blockSize)
}

func UnPadding(origData []byte, mode int64) []byte {
	if mode&Padding_PKCS5 != 0 {
		return PKCS5UnPadding(origData)
	}
	if mode&Padding_Zero != 0 {
		return ZeroUnPadding(origData)
	}
	return PKCS5UnPadding(origData)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
