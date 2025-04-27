package propertybasedtests

import "strings"

var (
	romanToArabic = []struct {
		roman  string
		arabic uint16
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	romanToArabicMap = map[string]uint16{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
)

func ToRoman(arabic uint16) string {
	var roman strings.Builder
	for _, v := range romanToArabic {
		for arabic >= v.arabic {
			roman.WriteString(v.roman)
			arabic -= v.arabic
		}
	}
	return roman.String()
}

func ToArabic(roman string) (arabic uint16) {
	for i := 0; i < len(roman); i++ {
		if i < len(roman)-1 {
			symbol := roman[i : i+2]
			converted := romanToArabicMap[symbol]
			if converted != 0 {
				arabic += converted
				i++
			} else {
				arabic += romanToArabicMap[string(roman[i])]
			}
		} else {
			arabic += romanToArabicMap[string(roman[i])]
		}
	}
	return arabic
}
