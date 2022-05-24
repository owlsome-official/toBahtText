package toBahtText

import (
	"strconv"
	"strings"
)

const MILLION = 6
const ED = "เอ็ด"
const YEE = "ยี่"

var units = map[int]string{
	0: "",
	1: "สิบ",
	2: "ร้อย",
	3: "พัน",
	4: "หมื่น",
	5: "แสน",
	6: "ล้าน",
}
var nums = map[string]string{
	// "0": "ศูนย์",
	"1": "หนึ่ง",
	"2": "สอง",
	"3": "สาม",
	"4": "สี่",
	"5": "ห้า",
	"6": "หก",
	"7": "เจ็ด",
	"8": "แปด",
	"9": "เก้า",
}

func From(input string) string {
	inputArr := validateValue(input)
	if len(inputArr) == 0 {
		return ""
	}
	revInputArr := reverse(inputArr)

	resultArr := []string{}

	for idx, v := range revInputArr {
		if v == "0" {
			resultArr = append(resultArr, "")
			continue
		}
		millionLvl := idx / MILLION
		unitIdx := idx - (millionLvl * MILLION)

		numberName := nums[v]
		unitName := units[unitIdx]
		numberWord := numberName + unitName

		if unitIdx == 1 { // หลักสิบ

			if len(resultArr) > 0 && resultArr[len(resultArr)-1] == "หนึ่ง" { // case หนึ่ง => เอ็ด
				resultArr[len(resultArr)-1] = ED
			}

			if v == "1" { // case หนึ่ง+สิบ => สิบ
				numberWord = unitName
			} else if v == "2" { // case สอง+สิบ -> ยี่สิบ
				numberWord = YEE + unitName
			}

		}

		resultArr = append(resultArr, numberWord)
	}

	for idx := 0; idx < len(resultArr); idx++ {
		if idx != 0 && idx%MILLION == 0 {
			resultArr[idx] = resultArr[idx] + "ล้าน"
		}
	}

	return strings.Join(reverse(resultArr), "")
}

func reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}

func validateValue(input string) []string {
	arr := strings.Split(input, "")
	for _, v := range arr {
		if _, err := strconv.Atoi(v); err != nil {
			return []string{}
		}
	}
	return arr
}
