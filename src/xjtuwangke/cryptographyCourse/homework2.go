package cryptographyCourse

import (
	"fmt"
	"xjtuwangke/crypto/aes"
	"xjtuwangke/utils"
)

type oneJob struct {
	Key  string
	Text string
	Mode int64
}

func RunCourseProjWeek2() {
	jobs := make([]oneJob, 0)
	jobs = append(jobs, oneJob{
		Key:  "140b41b22a29beb4061bda66b6747e14",
		Text: "4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81",
		Mode: aes.AES_DECRYPT | aes.Padding_PKCS5 | aes.AES_MODE_CBC,
	})
	jobs = append(jobs, oneJob{
		Key:  "140b41b22a29beb4061bda66b6747e14",
		Text: "5b68629feb8606f9a6667670b75b38a5b4832d0f26e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253",
		Mode: aes.AES_DECRYPT | aes.Padding_PKCS5 | aes.AES_MODE_CBC,
	})
	jobs = append(jobs, oneJob{
		Key:  "36f18357be4dbd77f050515c73fcf9f2",
		Text: "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329",
		Mode: aes.AES_DECRYPT | aes.AES_MODE_CTR,
	})
	jobs = append(jobs, oneJob{
		Key:  "36f18357be4dbd77f050515c73fcf9f2",
		Text: "770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451",
		Mode: aes.AES_DECRYPT | aes.AES_MODE_CTR,
	})
	for index, job := range jobs {
		fmt.Printf("index of %d\n", index)
		aesDecode(job.Key, job.Text, job.Mode)
	}
}

func aesDecode(key, text string, mode int64) {
	keyFlow := utils.HexStringToByte(key)
	textFlow := utils.HexStringToByte(text)
	//result, err := aes.AesEncryptAndDecrypt(textFlow, keyFlow, mode)
	result, err := aes.AesEncryptAndDecrypt(textFlow, keyFlow, mode)
	if err != nil {
		fmt.Printf("whoops there was an error:%s", err)
	}
	fmt.Printf("cipherText length is: %d Bytes\n", len(string(textFlow)))
	fmt.Printf("key length is: %d Bytes\n", len(string(keyFlow)))
	fmt.Printf("the decrypted string is:%s\n", string(result))
	fmt.Printf("with length:%d\n", len(string(result)))
}
