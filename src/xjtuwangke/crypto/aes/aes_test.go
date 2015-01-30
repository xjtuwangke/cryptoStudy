package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncryptAndDecrypt([]byte("polaris@studygolangxxxxxxxxxxxxx"), key, AES_ENCRYPT|AES_MODE_CBC|Padding_PKCS5)
	if err != nil {
		panic(err)
	}
	//fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesEncryptAndDecrypt(result, key, AES_DECRYPT|AES_MODE_CBC|Padding_PKCS5)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
