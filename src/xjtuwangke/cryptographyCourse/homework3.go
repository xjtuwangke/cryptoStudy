package cryptographyCourse

import (
	"fmt"
	"math"
	"os"
	//"crypto/sha1"
	//"io"
	"crypto/sha256"
	"xjtuwangke/utils"
)

const (
	blockSize = 1024
	fileName  = "/Volumes/WANGKE/courses/Stanford Cryptography I/6 Collision Resistance/6 - 1 - Introduction (11 min).mp4"
	//fileName = "/Volumes/WANGKE/courses/Stanford Cryptography I/6 Collision Resistance/6 - 2 - Generic birthday attack (16 min).mp4"
)

func RunCourseProjWeek3() {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("file open error")
		return
	}
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("file open error")
		return
	}
	defer fp.Close()
	//h := sha1.New()
	//fmt.Printf("%x\n", h.Sum(nil))
	blocks := int64(math.Ceil(float64(fileInfo.Size() / blockSize)))
	fmt.Printf("we have %d blocks for this file\n", blocks)
	//checkSum := make([]byte, sha256.Size)
	var checkSum [sha256.Size]byte
	for i := blocks; i >= 0; i-- {
		buf := make([]byte, blockSize)
		n, _ := fp.ReadAt(buf, int64(i*blockSize))
		if i == blocks {
			fmt.Printf("n is %d\n", n)
			checkSum = sha256.Sum256(buf[0:n:blockSize])
			fmt.Printf("first block , %s\n", utils.ByteToHexString(checkSum[:]))
		} else {
			checkSum = sha256.Sum256(append(buf[0:n:blockSize], checkSum[:]...))
		}
	}
	fmt.Printf("final string is %s\n", utils.ByteToHexString(checkSum[:]))
}
