package utils

func HexStringToByte(hex string) []byte {
	var temp, charInt int
	result := []byte{}
	mapHex := make(map[string]int)
	mapHex["0"] = 0
	mapHex["1"] = 1
	mapHex["2"] = 2
	mapHex["3"] = 3
	mapHex["4"] = 4
	mapHex["5"] = 5
	mapHex["6"] = 6
	mapHex["7"] = 7
	mapHex["8"] = 8
	mapHex["9"] = 9
	mapHex["a"] = 10
	mapHex["b"] = 11
	mapHex["c"] = 12
	mapHex["d"] = 13
	mapHex["e"] = 14
	mapHex["f"] = 15

	for i := 0; i < len(hex); i++ {
		s := hex[i : i+1]
		temp = mapHex[s]
		if i%2 == 0 {
			charInt = temp
		} else {
			charInt = temp + charInt*16
			result = append(result, byte(charInt))
		}
	}
	return result
}
