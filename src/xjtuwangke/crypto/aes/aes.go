package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	Padding_Zero  int64 = 1 << iota
	Padding_PKCS5 int64 = 1 << iota
	AES_ENCRYPT   int64 = 1 << iota
	AES_DECRYPT   int64 = 1 << iota
	AES_MODE_CBC  int64 = 1 << iota
	AES_MODE_CTR  int64 = 1 << iota
)

func AesEncryptAndDecrypt(text, key []byte, mode int64) ([]byte, error) {
	if mode&AES_MODE_CTR != 0 {
		return AesCTREncryptAndDecrypt(text, key, mode)
	}
	if mode&AES_MODE_CBC != 0 {
		return AesCBCEncryptAndDecrypt(text, key, mode)
	}
	return AesCBCEncryptAndDecrypt(text, key, mode)
}

func AesCBCEncryptAndDecrypt(text, key []byte, mode int64) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if mode&AES_ENCRYPT != 0 {

		// AES CBC加密
		// Step 1: 做padding
		origData := Padding(text, blockSize, mode)
		// 为加密后的数据申请空间
		crypted := make([]byte, len(origData))
		// key的长度与blockSize一致 因此截取key[:blockSize]
		encrypter := cipher.NewCBCEncrypter(block, key[:blockSize])
		encrypter.CryptBlocks(crypted, origData)
		return crypted, nil

	} else {
		decrypter := cipher.NewCBCDecrypter(block, key[:blockSize])
		// 不考虑Padding plainText与cipherText长度相同
		plainText := make([]byte, len(text))
		decrypter.CryptBlocks(plainText, text)
		// Final Step: unpadding
		plainText = UnPadding(plainText, mode)
		return plainText, nil
	}
}

func AesCTREncryptAndDecrypt(text, key []byte, mode int64) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if mode&AES_ENCRYPT != 0 {
		// The IV needs to be unique, but not secure. Therefore it's common to
		// include it at the beginning of the ciphertext.
		ciphertext := make([]byte, blockSize+len(text))
		iv := ciphertext[:blockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}

		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(ciphertext[blockSize:], text)
		return ciphertext, nil
	} else {
		// It's important to remember that ciphertexts must be authenticated
		// (i.e. by using crypto/hmac) as well as being encrypted in order to
		// be secure.

		// CTR mode is the same for both encryption and decryption, so we can
		// also decrypt that ciphertext with NewCTR.

		plaintext := make([]byte, len(text)-blockSize)
		iv := text[:blockSize]
		stream := cipher.NewCTR(block, iv)
		stream.XORKeyStream(plaintext, text[blockSize:])
		return plaintext, nil
	}
}
