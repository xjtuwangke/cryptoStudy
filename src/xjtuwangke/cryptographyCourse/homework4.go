package cryptographyCourse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"xjtuwangke/utils"
)

const (
	cbcBlockSize = 16
	url_target   = "http://crypto-class.appspot.com/po?er="
	cipherText   = "f20bdba6ff29eed7b046d1df9fb7000058b1ffb4210a580f748b4ac714c001bd4a61044426fb515dad3f21f18aa577c0bdf302936266926ff37dbf7035d5eeb4"
)

func RunCourseProjWeek4() {
	//iv := "f20bdba6ff29eed7b046d1df9fb70000"
	cipherText := []string{
		"58b1ffb4210a580f748b4ac714c001bd",
		"4a61044426fb515dad3f21f18aa577c0",
		"bdf302936266926ff37dbf7035d5eeb4"}
	plainText := ""
	//plainText += paddingOracleGuess(iv, cipherText[0])
	//fmt.Println(plainText)
	//plainText += paddingOracleGuess(cipherText[0], cipherText[1])
	//fmt.Println(plainText)
	plainText += paddingOracleGuess(cipherText[1], cipherText[2])
	fmt.Println(plainText)
}

func paddingMask(blockSize, padding int) []byte {
	return append(bytes.Repeat([]byte{byte(0)}, blockSize-padding), bytes.Repeat([]byte{byte(padding)}, padding)...)
}

func paddingOracleGuess(iv string, block string) string {
	ivByteFlow := utils.HexStringToByte(iv)
	blockByteFlow := utils.HexStringToByte(block)
	results := []byte{}
	for i := 0; i < cbcBlockSize; i++ {
		fmt.Printf("\n")
		padding := paddingMask(cbcBlockSize, i+1)
		for g := 0; g < 256; g++ {
			guess := append(bytes.Repeat([]byte{byte(0)}, cbcBlockSize-i-1), append([]byte{byte(g)}, results...)...)
			//ivByteFlow ^ guess ^ padding
			iv := utils.SliceXor(ivByteFlow, utils.SliceXor(guess, padding))
			//fmt.Println(utils.ByteToHexString(ivByteFlow))
			//fmt.Println(utils.ByteToHexString(guess))
			//fmt.Println(utils.ByteToHexString(padding))
			//fmt.Println(utils.ByteToHexString(iv))
			isHit := paddingOracleTry(append(iv, blockByteFlow...))
			fmt.Printf(".")
			if isHit {
				results = append([]byte{byte(g)}, results...)
				fmt.Printf("\ni=%d,g=%d\n", i, g)
				break
			}
		}
	}
	return string(results)
}

func paddingOracleTry(bytesFlow []byte) bool {
	/*
	   when a decrypted CBC ciphertext ends in an invalid pad the web server returns a 403 error code (forbidden request).
	   When the CBC padding is valid, but the message is malformed, the web server returns a 404 error code (URL not found).
	   The padding oracle will let you decrypt the given ciphertext one byte at a time.
	   To decrypt a single byte you will need to send up to 256 HTTP requests to the site.
	   Keep in mind that the first ciphertext block is the random IV. The decrypted message is ASCII encoded.
	*/
	para := utils.ByteToHexString(bytesFlow)
	statusCode, _ := httpGetQuery(para)
	//fmt.Println(para)
	//fmt.Println(statusCode)
	if statusCode == 404 {
		return true
	} else {
		return false
	}
}

func httpGetQuery(para string) (int, []byte) {
	resp, err := http.Get(url_target + para)
	if err != nil {
		// handle error
	}
	//resp.StatusCode
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return resp.StatusCode, body
}
