package day3

import (
	"strconv"
	"strings"
)

func makeMask(binaryStr string) string {
	var mask string
	for _, m := range strings.Split(binaryStr, "") {
		if m == "0" {
			mask += "1"
		} else {
			mask += "0"
		}
	}
	return mask
}

func BinaryDiagnostic(codes []string) (int64, error) {
	lineLen := len(codes[0])

	var commonNum = ""

	for i := 0; i <= lineLen-1; i++ {
		commonNums := make(map[int]int, 0)
		for _, code := range codes {
			if code == "" {
				continue
			}

			csplit := strings.Split(code, "")

			num, err := strconv.Atoi(string(csplit[i]))
			if err != nil {
				return 0, err
			}

			_, ok := commonNums[num]
			if !ok {
				commonNums[num] = 1
			} else {
				commonNums[num]++
			}
		}

		if commonNums[0] > commonNums[1] {
			commonNum += "0"
		} else {
			commonNum += "1"
		}
	}

	commonN, err := strconv.ParseInt(commonNum, 2, 64)
	if err != nil {
		return 0, err
	}

	mask := makeMask(commonNum)
	maskN, err := strconv.ParseInt(mask, 2, 64)
	if err != nil {
		return 0, err
	}

	return maskN * commonN, nil
}
