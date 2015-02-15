package main

import (
	"fmt"
	"xjtuwangke/cryptographyCourse"
	"xjtuwangke/utils"
)

func main() {
	//cryptographyCourse.RunCourseProjWeek2()
	//cryptographyCourse.RunCourseProjWeek3()
	//testUtils()
	cryptographyCourse.RunCourseProjWeek4()
}

func testUtils() {
	stream := []byte{0x00, 0x1a, 0x2b, 0x3c, 0x5f}
	fmt.Printf("%s", utils.ByteToHexString(stream))
}

func testSliceXor() {

}
