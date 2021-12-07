package day3

import (
	"log"
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

func filterOut(codes []string, pos int, num int) {
	for i, code := range codes {
		if code == "" {
			continue
		}
		csplit := strings.Split(code, "")
		if csplit[pos] != strconv.Itoa(num) {
			codes = append(codes[:i], codes[i+1:]...)
		}
	}
}

func BinaryFilter(codes []string) (int64, error) {
	lineLen := len(codes[0])

	var (
		commonCodes = make([]string, len(codes))
		leastCodes  = make([]string, len(codes))
	)

	copy(commonCodes, codes)
	copy(leastCodes, codes)

	for i := 0; i <= lineLen-1; i++ {
		commonNums := make(map[int]int, 0)
		for _, code := range codes {
			if err := commonBin(i, code, commonNums); err != nil {
				return 0, err
			}
		}

		if commonNums[0] > commonNums[1] {
			filterOut(commonCodes, i, 0)
			filterOut(leastCodes, i, 1)
		} else {
			filterOut(commonCodes, i, 1)
			filterOut(leastCodes, i, 0)
		}
		log.Printf("common: %v least: %v", commonCodes, leastCodes)
	}

	commonN, err := strconv.ParseInt(commonCodes[0], 2, 64)
	if err != nil {
		return 0, err
	}

	leastN, err := strconv.ParseInt(leastCodes[0], 2, 64)
	if err != nil {
		return 0, err
	}

	return leastN * commonN, nil
}

func commonBin(i int, code string, m map[int]int) error {
	if code == "" {
		return nil
	}

	csplit := strings.Split(code, "")

	num, err := strconv.Atoi(string(csplit[i]))
	if err != nil {
		return err
	}

	_, ok := m[num]
	if !ok {
		m[num] = 1
	} else {
		m[num]++
	}

	return nil
}

func BinaryDiagnostic(codes []string) (int64, error) {
	lineLen := len(codes[0])

	var commonNum = ""

	for i := 0; i <= lineLen-1; i++ {
		commonNums := make(map[int]int, 0)
		for _, code := range codes {
			if err := commonBin(i, code, commonNums); err != nil {
				return 0, err
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
